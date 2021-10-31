package logic

import (
	"context"
	"strconv"
	"time"

	"arod-im/internal/logic/conf"
	"arod-im/internal/logic/dao"
	"arod-im/internal/logic/model"

	"github.com/bilibili/discovery/naming"
	"github.com/golang/glog"
)

const (
	_onlineTick     = time.Second * 10
	_onlineDeadline = time.Minute * 5
)

type Logic struct {
	c   *conf.Config
	dis *naming.Discovery
	dao *dao.Dao
	// online
	totalIPs   int64
	totalConns int64
	roomCount  map[string]int32
	// load balancer
	nodes        []*naming.Instance
	loadBalancer *LoadBalancer
	regions      map[string]string
}

func New(c *conf.Config) (l *Logic) {
	l = &Logic{
		c:            c,
		dao:          dao.New(c),
		dis:          naming.New(c.Discovery),
		loadBalancer: NewLoadBalancer(),
		regions:      make(map[string]string),
	}
	l.initRegions()
	l.initNodes()
	_ = l.loadOnline()
	go l.onlineproc()
	return l
}

// Ping ping resources is ok.
func (l *Logic) Ping(c context.Context) error {
	return l.dao.Ping(c)
}

// Close close resources.
func (l *Logic) Close() {
	l.dao.Close()
}

func (l *Logic) initRegions() {
	for region, ps := range l.c.Regions {
		for _, province := range ps {
			l.regions[province] = region
		}
	}
}

func (l *Logic) initNodes() {
	res := l.dis.Build("gim.comet")
	event := res.Watch()
	select {
	case _, ok := <-event:
		if ok {
			l.newNodes(res)
		} else {
			panic("discovery watch failed")
		}
	case <-time.After(10 * time.Second):
		glog.Error("discovery start timeout")
	}
	go func() {
		for {
			if _, ok := <-event; !ok {
				return
			}
			l.newNodes(res)
		}
	}()
}

func (l *Logic) newNodes(res naming.Resolver) {
	if zoneIns, ok := res.Fetch(); ok {
		var (
			totalConns int64
			totalIps   int64
			allIns     []*naming.Instance
		)
		for _, zins := range zoneIns.Instances {
			for _, ins := range zins {
				if ins.Metadata == nil {
					glog.Errorf("node instance metadata is empty(%+v)", ins)
					continue
				}
				offline, err := strconv.ParseBool(ins.Metadata[model.MetaOffline])
				if err != nil || offline {
					glog.Warningf("strconv.ParseBool(offline:%t) error(%v)", offline, err)
					continue
				}
				conns, err := strconv.ParseInt(ins.Metadata[model.MetaConnCount], 10, 32)
				if err != nil {
					glog.Errorf("strconv.ParseInt(conns:%d) error(%v)", conns, err)
					continue
				}
				ips, err := strconv.ParseInt(ins.Metadata[model.MetaIPCount], 10, 32)
				if err != nil {
					glog.Errorf("strconv.ParseInt(ips:%d) error(%v)", ips, err)
					continue
				}
				totalConns += conns
				totalIps += ips
				allIns = append(allIns, ins)
			}
		}
		l.totalConns = totalConns
		l.totalIPs = totalIps
		l.nodes = allIns
		l.loadBalancer.Update(allIns)
	}
}

func (l *Logic) onlineproc() {
	for {
		time.Sleep(_onlineTick)
		if err := l.loadOnline(); err != nil {
			glog.Errorf("onlineproc error(%v)", err)
		}
	}
}

func (l *Logic) loadOnline() (err error) {
	roomCount := make(map[string]int32)
	for _, server := range l.nodes {
		var online *model.Online
		online, err = l.dao.ServerOnline(context.Background(), server.Hostname)
		if err != nil {
			return
		}
		if time.Since(time.Unix(online.Updated, 0)) > _onlineDeadline {
			l.dao.DelServerOnline(context.Background(), server.Hostname)
			continue
		}
		for roomID, count := range online.RoomCount {
			roomCount[roomID] += count
		}
	}
	l.roomCount = roomCount
	return
}

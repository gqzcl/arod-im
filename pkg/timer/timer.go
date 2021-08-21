package timer

import (
	"sync"
	itime "time"

	log "github.com/golang/glog"
)

const (
	timerFormat = "2021-08-13 23:04:05"
	// 最大持续时间
	infiniteDuration = itime.Duration(1<<63 - 1)
)

// TimerData 定时数据
type TimerData struct {
	Key    string
	expire itime.Time
	fn     func()
	index  int
	next   *TimerData
}

// Delay 返回持续时间
func (td *TimerData) Delay() itime.Duration {
	return itime.Until(td.expire)
}

// ExpireString 返回字符串过期时间
func (td *TimerData) ExpireString() string {
	return td.expire.Format(timerFormat)
}

// Timer 定时器
type Timer struct {
	lock   sync.Mutex
	free   *TimerData
	timers []*TimerData
	// 定时器
	signal *itime.Timer
	num    int
}

// NewTimer new a timer.
// A heap must be initialized before any of the heap operations
// can be used. Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// Its complexity is O(n) where n = h.Len().
func NewTimer(num int) (t *Timer) {
	t = new(Timer)
	t.init(num)
	return t
}

func (t *Timer) Init(num int) {
	t.init(num)
}

// Del removes the element at index i from the heap. The complexity is O(log(n)) where n = h.Len().
func (t *Timer) Del(td *TimerData) {
	t.lock.Lock()
	t.del(td)
	t.put(td)
	t.lock.Unlock()
}

func (t *Timer) Add(expire itime.Duration, fn func()) (td *TimerData) {
	t.lock.Lock()
	td = t.get()
	td.expire = itime.Now().Add(expire)
	td.fn = fn
	t.add(td)
	t.lock.Unlock()
	return
}

func (t *Timer) Set(td *TimerData, expire itime.Duration) {
	t.lock.Lock()
	// 不进行修改操作，而是删除并添加
	t.del(td)
	td.expire = itime.Now().Add(expire)
	t.add(td)
	t.lock.Unlock()
}

func (t *Timer) init(num int) {
	t.timers = make([]*TimerData, 0, num)
	t.signal = itime.NewTimer(infiniteDuration)
	t.num = num
	t.grow()
	go t.start()
}

// 赋值给 t.free 一个长度为t.num的TimerData链表
func (t *Timer) grow() {
	var (
		i   int
		td  *TimerData
		tds = make([]TimerData, t.num)
	)
	t.free = &(tds[0])
	td = t.free
	for i = 1; i < t.num; i++ {
		td.next = &tds[i]
		td = td.next
	}
	td.next = nil
}

// 开始计时
func (t *Timer) start() {
	for {
		t.expire()
		<-t.signal.C
	}
}

func (t *Timer) expire() {
	var fn func()
	var td *TimerData
	var d itime.Duration

	// 加锁是因为在start函数中被调用了，而在init函数中开启了start函数携程
	t.lock.Lock()
	for {
		if len(t.timers) == 0 {
			d = infiniteDuration
			if Debug {
				log.Info("timer: no other instance")
			}
			break
		}
		td = t.timers[0]
		if d = td.Delay(); d > 0 {
			break
		}
		fn = td.fn
		t.del(td)
		t.lock.Unlock()
		if fn == nil {
			log.Warning("expire timer no fn")
		} else {
			if Debug {
				log.Infof("timer key: %s,expire: %s,index: %d expired,call fn", td.Key, td.ExpireString(), td.index)
			}
			fn()
		}
		t.lock.Lock()
	}
	t.signal.Reset(d)
	if Debug {
		log.Infof("timer: expire reset delay %d ms", int64(d)/int64(itime.Millisecond))
	}
	t.lock.Unlock()
}

// 删除此TimerData并对列表进行一定调整
func (t *Timer) del(td *TimerData) {
	var (
		i    = td.index
		last = len(t.timers) - 1
	)
	if i < 0 || i > last || t.timers[i] != td {
		// already remove, usually by expire
		if Debug {
			log.Infof("timer del i: %d, last: %d, %p", i, last, td)
		}
		return
	}
	if i != last {
		// 放到队尾
		t.swap(i, last)
		t.down(i, last)
		t.up(i)
	}
	// remove item is the last node
	t.timers[last].index = -1 // for safely
	t.timers = t.timers[:last]
	if Debug {
		log.Infof("timer:remove item key: %s.expire: %s,index: %d", td.Key, td.ExpireString(), td.index)
	}
}

func (t *Timer) swap(i, j int) {
	t.timers[i], t.timers[j] = t.timers[j], t.timers[i]
	t.timers[i].index = i
	t.timers[j].index = j
}

// 检查子节点，交换大小顺序，将小的排到前面
func (t *Timer) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			// j1<0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && t.less(j2, j1) {
			// j2 < j1，如果j2比j1小，j=j2
			j = j2 // right child
		}
		// 如果i<j,退出循环
		if t.less(i, j) {
			break
		}
		// 如果j<i,交换
		t.swap(i, j)
		i = j
	}
}

// 检查父节点，如果j小于其父亲，则交换位置
func (t *Timer) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i >= j || t.less(i, j) {
			break
		}
		t.swap(i, j)
		j = i
	}
}
func (t *Timer) less(i, j int) bool {
	return t.timers[i].expire.Before(t.timers[j].expire)
}

// 获取下一个空闲的TimerData
func (t *Timer) get() (td *TimerData) {
	if td = t.free; td == nil {
		t.grow()
		td = t.free
	}
	t.free = td.next
	return
}

// Push pushes the element x onto the heap. The complexity is O(log(n)) where n = h.Len().
func (t *Timer) add(td *TimerData) {
	var d itime.Duration
	td.index = len(t.timers)
	t.timers = append(t.timers, td)
	t.up(td.index)
	if td.index == 0 {
		d = td.Delay()
		t.signal.Reset(d)
		if Debug {
			log.Infof("timer: add reset delay %d ms", int64(d)/int64(itime.Millisecond))
		}
	}
	if Debug {
		log.Infof("timer: push item key:%s, expire: %s,index: %d", td.Key, td.ExpireString(), td.index)
	}
}

// 回收TimerData，将其交给t.free
func (t *Timer) put(td *TimerData) {
	td.fn = nil
	td.next = t.free
	t.free = td
}

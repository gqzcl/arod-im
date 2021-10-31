package conf

import (
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/bilibili/discovery/naming"
)

var (
	confPath  string
	region    string
	zone      string
	deployEnv string
	host      string
	weight    int64
	Conf      *Config
)

func Init() error {
	Conf = Default()
	_, err := toml.DecodeFile(confPath, &Conf)
	return err
}

func init() {
	var (
		defHost, _   = os.Hostname()
		defWeight, _ = strconv.ParseInt(os.Getenv("WEIGHT"), 10, 32)
	)
	flag.StringVar(&confPath, "conf", "logic.toml", "default config path")
	flag.StringVar(&region, "region", os.Getenv("REGION"), "avaliable region. or use REGION env variable,value:sh etc.")
	flag.StringVar(&zone, "zone", os.Getenv("ZONE"), "avaliable zone. or use ZONE env variable, value: sh001/sh002 etc.")
	flag.StringVar(&deployEnv, "deploy.env", os.Getenv("DEPLOY_ENV"), "deploy env. or use DEPLOY_ENV env variable, value: dev/fat1/uat/pre/prod etc.")
	flag.StringVar(&host, "host", defHost, "machine hostname. or use default machine hostname.")
	flag.Int64Var(&weight, "weight", defWeight, "load balancing weight, or use WEIGHT env variable, value: 10 etc.")
}

type Config struct {
	Env *Env
	// ???
	Discovery  *naming.Config
	RPCClient  *RPCClient
	RPCServer  *RPCServer
	HTTPServer *HTTPServer
	Kafka      *Kafka
	Redis      *Redis
	Node       *Node
	Backoff    *Backoff
	Regions    map[string][]string
}

func Default() *Config {
	return &Config{
		Env: &Env{
			Region:    region,
			Zone:      zone,
			DeployEnv: deployEnv,
			Host:      host,
			Weight:    weight,
		},
		Discovery: &naming.Config{
			Region: region,
			Zone:   zone,
			Env:    deployEnv,
			Host:   host,
		},
		HTTPServer: &HTTPServer{
			Network:      "tcp",
			Addr:         "3111",
			ReadTimeout:  time.Second,
			WriteTimeout: time.Second,
		},
		RPCClient: &RPCClient{
			Dial:    time.Second,
			Timeout: time.Second,
		},
		RPCServer: &RPCServer{
			Network:           "tcp",
			Addr:              "3119",
			Timeout:           time.Second,
			IdleTimeout:       time.Second * 60,
			MaxLifeTime:       time.Hour * 2,
			ForceCloseWait:    time.Second * 20,
			KeepAliveInterval: time.Second * 60,
			KeepAliveTimeout:  time.Second * 20,
		},
		Backoff: &Backoff{MaxDelay: 300, BaseDelay: 3, Factor: 1.8, Jitter: 1.3},
	}
}

type Env struct {
	Region    string
	Zone      string
	DeployEnv string
	Host      string
	Weight    int64
}

type Node struct {
	DefaultDomain string
	HostDomain    string
	TCPPort       int
	WSPort        int
	WSSPort       int
	HeartbeatMax  int
	Heartbeat     time.Duration
	RegionWeight  float64
}

type Backoff struct {
	MaxDelay  int32
	BaseDelay int32
	Factor    float32
	Jitter    float32
}

type Redis struct {
	Network      string
	Addr         string
	Auth         string
	Active       int
	Idle         int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Expire       time.Duration
}

type Kafka struct {
	Topic   string
	Brokers []string
}

type RPCClient struct {
	Dial    time.Duration
	Timeout time.Duration
}

type RPCServer struct {
	Network           string
	Addr              string
	Timeout           time.Duration
	IdleTimeout       time.Duration
	MaxLifeTime       time.Duration
	ForceCloseWait    time.Duration
	KeepAliveInterval time.Duration
	KeepAliveTimeout  time.Duration
}

type HTTPServer struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

package conf

import "time"

type Config struct {
	Env *Env
	// ???
	//Discovery  *naming.Config
	RPCClient  *RPCClient
	RPCServer  *RPCServer
	HTTPServer *HTTPServer
	Kafka      *Kafka
	Redis      *Redis
	Node       *Node
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

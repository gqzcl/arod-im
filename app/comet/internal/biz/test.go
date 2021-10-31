type Bucket struct {
	c     *conf.Bucket
	cLock sync.RWMutex        // protect the channels for chs
	chs   map[string]*Channel // map sub key to a channel
	// room
	rooms       map[string]*Room // bucket room channels
	routines    []chan *v1.BroadcastRoomReq
	routinesNum uint64

	ipCnts map[string]int32
}
type BucketRepo interface {
}
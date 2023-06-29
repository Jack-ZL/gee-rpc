package geeCache

import pb "gee-rpc/geeCache/geeCachepb"

//根据传入的 key 选择相应节点
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

//从对应 group 查找缓存值
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}

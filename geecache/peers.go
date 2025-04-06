package geecache

import pb "gee-cache/geecachepb"

type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}

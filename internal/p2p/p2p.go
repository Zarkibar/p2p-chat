package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	peerstore "github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func CreateNode() host.Host {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
	)
	if err != nil {
		panic(err)
	}

	return node

}

func HandleChatStream(s network.Stream, callback func(string)) {
	defer s.Close()

	buff := make([]byte, 1024)
	n, _ := s.Read(buff)
	msg := string(buff[:n])
	callback(msg)
	s.Write([]byte("ack"))
}

func ConnectToPeer(node host.Host, ctx context.Context, nodeMultiaddr string) *peerstore.AddrInfo {
	peerAddr, _ := multiaddr.NewMultiaddr(nodeMultiaddr)
	peerinfo, _ := peerstore.AddrInfoFromP2pAddr(peerAddr)

	if err := node.Connect(ctx, *peerinfo); err != nil {
		panic(err)
	}

	return peerinfo
}

func SendMsg(node host.Host, ctx context.Context, peerinfo *peerstore.AddrInfo, msg string) string {
	s, err := node.NewStream(ctx, peerinfo.ID, "/chat/1.0.0")
	if err != nil {
		panic(err)
	}

	s.Write([]byte(msg))

	return RecvMsg(s)
}

func RecvMsg(s network.Stream) string {
	buff := make([]byte, 1024)
	n, _ := s.Read(buff)
	return string(buff[:n])
}

package datanode

import (
	"errors"
	"sync"

	"github.com/uber/aresdb/cluster/topology"
	"github.com/uber/aresdb/datanode/client"
	"github.com/uber/aresdb/datanode/generated/proto/rpc"
	"google.golang.org/grpc"
)

var (
	errPeerClosed   = errors.New("peer closed")
	errPeerNotExist = errors.New("peer does not exist")
)

var grpcDialer = func(target string, opts ...grpc.DialOption) (client rpc.PeerDataNodeClient, closeFn func() error, err error) {
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, nil, err
	}
	return rpc.NewPeerDataNodeClient(conn), func() error { return conn.Close() }, nil
}

type peer struct {
	sync.RWMutex
	sync.WaitGroup

	host topology.Host

	dialer  client.PeerConnDialer
	conn    rpc.PeerDataNodeClient
	closeFn func() error

	closed bool
}

func (p *peer) Host() topology.Host {
	_ = "STUB: not implemented"

	// BorrowConnection from peer
	return *new(topology.Host)
}

func (p *peer) BorrowConnection(fn client.WithConnectionFn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *peer) checkHealth(peerNodeClient rpc.PeerDataNodeClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Close close the peer
func (p *peer) Close() { _ = "STUB: not implemented"; return }

// waiting for on going operation with client connection

// newPeer create a new peer object
func newPeer(host topology.Host, dialer client.PeerConnDialer) *peer {
	_ = "STUB: not implemented"
	return nil
}

type peerSource struct {
	sync.RWMutex

	watch  topology.MapWatch
	peers  map[string]*peer
	dialer client.PeerConnDialer

	done chan struct{}
}

func (ps *peerSource) borrowConnection(hostID string, fn client.WithConnectionFn) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *peerSource) BorrowConnection(hostIDs []string, fn client.WithConnectionFn) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *peerSource) watchTopoChange() { _ = "STUB: not implemented"; return }

func (ps *peerSource) Close() { _ = "STUB: not implemented"; return }

func (ps *peerSource) updateTopoMap(topoMap topology.Map) { _ = "STUB: not implemented"; return }

// unknown host

// NewPeerSource creates PeerSource
func NewPeerSource(topo topology.Topology, dialerOverride client.PeerConnDialer) (client.PeerSource, error) {
	_ = "STUB: not implemented"
	return *new(client.PeerSource), nil
}

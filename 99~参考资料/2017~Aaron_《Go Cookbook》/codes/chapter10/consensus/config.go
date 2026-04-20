package consensus

import "github.com/hashicorp/raft"

var rafts map[string]*raft.Raft

func init() {
	rafts = make(map[string]*raft.Raft)
}

// Config creates num in-memory raft
// nodes and connects them
func Config(num int) {
	conf := raft.DefaultConfig()
	snapshotStore := raft.NewDiscardSnapshotStore()

	addrs := []string{}
	transports := []*raft.InmemTransport{}
	for i := 0; i < num; i++ {
		addr, transport := raft.NewInmemTransport("")
		addrs = append(addrs, addr)
		transports = append(transports, transport)
	}
	peerStore := &raft.StaticPeers{StaticPeers: addrs}
	memstore := raft.NewInmemStore()

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i != j {
				transports[i].Connect(addrs[j], transports[j])
			}
		}

		r, err := raft.NewRaft(conf, NewFSM(), memstore, memstore, snapshotStore, peerStore, transports[i])
		if err != nil {
			panic(err)
		}
		r.SetPeers(addrs)
		rafts[addrs[i]] = r
	}
}

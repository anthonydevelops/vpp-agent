package api

// SystemStats represents global system statistics.
type SystemStats struct {
	VectorRate     float64
	InputRate      float64
	LastUpdate     float64
	LastStatsClear float64
	Heartbeat      float64
}

// NodeStats represents per node statistics.
type NodeStats struct {
	Nodes []NodeCounters
}

// NodeCounters represents node counters.
type NodeCounters struct {
	NodeIndex uint32
	// TODO: node name is not currently retrievable via stats API (will be most likely added in 19.04)
	//NodeName string

	Clocks   uint64
	Vectors  uint64
	Calls    uint64
	Suspends uint64
}

// InterfaceStats represents per interface statistics.
type InterfaceStats struct {
	Interfaces []InterfaceCounters
}

// InterfaceCounters represents interface counters.
type InterfaceCounters struct {
	InterfaceIndex uint32
	// TODO: interface name is not currently retrievable via stats API (will be most likely added in 19.04)
	//InterfaceName string

	RxPackets uint64
	RxBytes   uint64
	RxErrors  uint64
	TxPackets uint64
	TxBytes   uint64
	TxErrors  uint64

	RxUnicast     [2]uint64 // packets[0], bytes[1]
	RxMulticast   [2]uint64 // packets[0], bytes[1]
	RxBroadcast   [2]uint64 // packets[0], bytes[1]
	TxUnicastMiss [2]uint64 // packets[0], bytes[1]
	TxMulticast   [2]uint64 // packets[0], bytes[1]
	TxBroadcast   [2]uint64 // packets[0], bytes[1]

	Drops   uint64
	Punts   uint64
	IP4     uint64
	IP6     uint64
	RxNoBuf uint64
	RxMiss  uint64
}

// ErrorStats represents statistics per error counter.
type ErrorStats struct {
	Errors []ErrorCounter
}

// ErrorCounter represents error counter.
type ErrorCounter struct {
	CounterName string
	Value       uint64
}

// StatsProvider provides the methods for getting statistics.
type StatsProvider interface {
	GetSystemStats() (*SystemStats, error)
	GetNodeStats() (*NodeStats, error)
	GetInterfaceStats() (*InterfaceStats, error)
	GetErrorStats(names ...string) (*ErrorStats, error)
}

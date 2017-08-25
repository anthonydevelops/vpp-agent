package vpp

import (
	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/resync"
	"github.com/ligato/cn-infra/flavors/etcdkafka"
	"github.com/ligato/vpp-agent/plugins/defaultplugins"
	"github.com/ligato/vpp-agent/plugins/govppmux"
	"github.com/ligato/vpp-agent/plugins/linuxplugin"
	"github.com/ligato/cn-infra/flavors/redis"
	"github.com/ligato/cn-infra/flavors/rpc"
)

// Flavor glues together multiple plugins to translate ETCD configuration into VPP.
type Flavor struct {
	EtcdKafka etcdkafka.FlavorEtcdKafka
	Redis     redis.FlavorRedis
	RPC       rpc.FlavorRPC
	Resync    resync.Plugin
	GoVPP     govppmux.GOVPPPlugin
	Linux     linuxplugin.Plugin
	VPP       defaultplugins.Plugin

	injected bool
}

// Inject sets object references
func (f *Flavor) Inject() error {
	if f.injected {
		return nil
	}
	f.injected = true

	f.EtcdKafka.Inject()
	f.Redis.Inject()
	f.RPC.Inject()

	// Aggregated transport
	adapters := []datasync.KeyProtoValWriter{&f.EtcdKafka.ETCDDataSync, &f.Redis.RedisDataSync}
	compositePublisher := datasync.CompositeKVProtoWriter{
		Adapters: adapters,
	}

	f.GoVPP.Deps.PluginInfraDeps = *f.EtcdKafka.FlavorLocal.InfraDeps("govpp")
	f.VPP.Deps.PluginInfraDeps = *f.EtcdKafka.FlavorLocal.InfraDeps("default-plugins")
	f.VPP.Deps.Publish = &f.EtcdKafka.ETCDDataSync
	f.VPP.Deps.PublishStatistics = compositePublisher
	f.VPP.Deps.Watch = &f.EtcdKafka.ETCDDataSync
	f.VPP.Deps.Kafka = &f.EtcdKafka.Kafka
	f.VPP.Deps.GoVppmux = &f.GoVPP
	f.VPP.Deps.Linux = &f.Linux
	f.VPP.Linux.Deps.Watcher = &f.EtcdKafka.ETCDDataSync

	return nil
}

// Plugins combines Generic Plugins and Standard VPP Plugins + (their ETCD Connector/Adapter with RESYNC)
func (f *Flavor) Plugins() []*core.NamedPlugin {
	f.Inject()
	return core.ListPluginsInFlavor(f)
}

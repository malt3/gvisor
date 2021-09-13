// automatically generated by stateify.

package network

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (e *Endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/internal/network.Endpoint"
}

func (e *Endpoint) StateFields() []string {
	return []string{
		"ops",
		"netProto",
		"transProto",
		"state",
		"info",
		"owner",
		"writeShutdown",
		"effectiveNetProto",
		"multicastMemberships",
		"ttl",
		"multicastTTL",
		"multicastAddr",
		"multicastNICID",
		"ipv4TOS",
		"ipv6TClass",
	}
}

func (e *Endpoint) beforeSave() {}

// +checklocksignore
func (e *Endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.ops)
	stateSinkObject.Save(1, &e.netProto)
	stateSinkObject.Save(2, &e.transProto)
	stateSinkObject.Save(3, &e.state)
	stateSinkObject.Save(4, &e.info)
	stateSinkObject.Save(5, &e.owner)
	stateSinkObject.Save(6, &e.writeShutdown)
	stateSinkObject.Save(7, &e.effectiveNetProto)
	stateSinkObject.Save(8, &e.multicastMemberships)
	stateSinkObject.Save(9, &e.ttl)
	stateSinkObject.Save(10, &e.multicastTTL)
	stateSinkObject.Save(11, &e.multicastAddr)
	stateSinkObject.Save(12, &e.multicastNICID)
	stateSinkObject.Save(13, &e.ipv4TOS)
	stateSinkObject.Save(14, &e.ipv6TClass)
}

func (e *Endpoint) afterLoad() {}

// +checklocksignore
func (e *Endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.ops)
	stateSourceObject.Load(1, &e.netProto)
	stateSourceObject.Load(2, &e.transProto)
	stateSourceObject.Load(3, &e.state)
	stateSourceObject.Load(4, &e.info)
	stateSourceObject.Load(5, &e.owner)
	stateSourceObject.Load(6, &e.writeShutdown)
	stateSourceObject.Load(7, &e.effectiveNetProto)
	stateSourceObject.Load(8, &e.multicastMemberships)
	stateSourceObject.Load(9, &e.ttl)
	stateSourceObject.Load(10, &e.multicastTTL)
	stateSourceObject.Load(11, &e.multicastAddr)
	stateSourceObject.Load(12, &e.multicastNICID)
	stateSourceObject.Load(13, &e.ipv4TOS)
	stateSourceObject.Load(14, &e.ipv6TClass)
}

func (m *multicastMembership) StateTypeName() string {
	return "pkg/tcpip/transport/internal/network.multicastMembership"
}

func (m *multicastMembership) StateFields() []string {
	return []string{
		"nicID",
		"multicastAddr",
	}
}

func (m *multicastMembership) beforeSave() {}

// +checklocksignore
func (m *multicastMembership) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.nicID)
	stateSinkObject.Save(1, &m.multicastAddr)
}

func (m *multicastMembership) afterLoad() {}

// +checklocksignore
func (m *multicastMembership) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.nicID)
	stateSourceObject.Load(1, &m.multicastAddr)
}

func init() {
	state.Register((*Endpoint)(nil))
	state.Register((*multicastMembership)(nil))
}

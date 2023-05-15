// automatically generated by stateify.

//go:build amd64
// +build amd64

package usertrap

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *State) StateTypeName() string {
	return "pkg/sentry/platform/systrap/usertrap.State"
}

func (s *State) StateFields() []string {
	return []string{
		"nextTrap",
		"tableAddr",
	}
}

func (s *State) beforeSave() {}

// +checklocksignore
func (s *State) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.nextTrap)
	stateSinkObject.Save(1, &s.tableAddr)
}

func (s *State) afterLoad() {}

// +checklocksignore
func (s *State) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.nextTrap)
	stateSourceObject.Load(1, &s.tableAddr)
}

func init() {
	state.Register((*State)(nil))
}
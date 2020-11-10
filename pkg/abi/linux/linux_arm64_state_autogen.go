// automatically generated by stateify.

// +build arm64
// +build arm64
// +build arm64
// +build arm64

package linux

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (p *PtraceRegs) StateTypeName() string {
	return "pkg/abi/linux.PtraceRegs"
}

func (p *PtraceRegs) StateFields() []string {
	return []string{
		"Regs",
		"Sp",
		"Pc",
		"Pstate",
	}
}

func (p *PtraceRegs) beforeSave() {}

func (p *PtraceRegs) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Regs)
	stateSinkObject.Save(1, &p.Sp)
	stateSinkObject.Save(2, &p.Pc)
	stateSinkObject.Save(3, &p.Pstate)
}

func (p *PtraceRegs) afterLoad() {}

func (p *PtraceRegs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Regs)
	stateSourceObject.Load(1, &p.Sp)
	stateSourceObject.Load(2, &p.Pc)
	stateSourceObject.Load(3, &p.Pstate)
}

func init() {
	state.Register((*PtraceRegs)(nil))
}

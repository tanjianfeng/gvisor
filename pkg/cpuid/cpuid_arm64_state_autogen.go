// automatically generated by stateify.

// +build arm64

package cpuid

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (fs *FeatureSet) StateTypeName() string {
	return "pkg/cpuid.FeatureSet"
}

func (fs *FeatureSet) StateFields() []string {
	return []string{
		"Set",
		"CPUImplementer",
		"CPUArchitecture",
		"CPUVariant",
		"CPUPartnum",
		"CPURevision",
	}
}

func (fs *FeatureSet) beforeSave() {}

func (fs *FeatureSet) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.Set)
	stateSinkObject.Save(1, &fs.CPUImplementer)
	stateSinkObject.Save(2, &fs.CPUArchitecture)
	stateSinkObject.Save(3, &fs.CPUVariant)
	stateSinkObject.Save(4, &fs.CPUPartnum)
	stateSinkObject.Save(5, &fs.CPURevision)
}

func (fs *FeatureSet) afterLoad() {}

func (fs *FeatureSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.Set)
	stateSourceObject.Load(1, &fs.CPUImplementer)
	stateSourceObject.Load(2, &fs.CPUArchitecture)
	stateSourceObject.Load(3, &fs.CPUVariant)
	stateSourceObject.Load(4, &fs.CPUPartnum)
	stateSourceObject.Load(5, &fs.CPURevision)
}

func init() {
	state.Register((*FeatureSet)(nil))
}

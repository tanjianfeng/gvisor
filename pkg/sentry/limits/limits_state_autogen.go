// automatically generated by stateify.

package limits

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (l *Limit) StateTypeName() string {
	return "pkg/sentry/limits.Limit"
}

func (l *Limit) StateFields() []string {
	return []string{
		"Cur",
		"Max",
	}
}

func (l *Limit) beforeSave() {}

func (l *Limit) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.Cur)
	stateSinkObject.Save(1, &l.Max)
}

func (l *Limit) afterLoad() {}

func (l *Limit) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.Cur)
	stateSourceObject.Load(1, &l.Max)
}

func (l *LimitSet) StateTypeName() string {
	return "pkg/sentry/limits.LimitSet"
}

func (l *LimitSet) StateFields() []string {
	return []string{
		"data",
	}
}

func (l *LimitSet) beforeSave() {}

func (l *LimitSet) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.data)
}

func (l *LimitSet) afterLoad() {}

func (l *LimitSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.data)
}

func init() {
	state.Register((*Limit)(nil))
	state.Register((*LimitSet)(nil))
}

package conflict

import (
	"goldenmaster/tb_conflict/api"
)

type conflict2Impl struct {
	api.INotifier
	sameName int64
}

var _ api.Conflict2 = (*conflict2Impl)(nil)
var _ api.INotifier = (*conflict2Impl)(nil)

func NewConflict2(notifier api.INotifier) api.Conflict2 {
	obj := &conflict2Impl{
		INotifier: notifier,
		sameName:  int64(0),
	}
	return obj
}

// property get sameName
func (c *conflict2Impl) GetSameName() int64 {
	return int64(0)
}

// property set sameName
func (c *conflict2Impl) SetSameName(sameName int64) {
}

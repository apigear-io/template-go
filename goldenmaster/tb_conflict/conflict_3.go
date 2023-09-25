package conflict

import (
	"goldenmaster/tb_conflict/api"
)

type conflict3Impl struct {
	api.INotifier
}

var _ api.Conflict3 = (*conflict3Impl)(nil)
var _ api.INotifier = (*conflict3Impl)(nil)

func NewConflict3(notifier api.INotifier) api.Conflict3 {
	obj := &conflict3Impl{
		INotifier: notifier,
	}
	return obj
}

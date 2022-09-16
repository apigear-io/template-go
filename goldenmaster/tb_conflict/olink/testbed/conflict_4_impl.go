package testbed

import (
	"goldenmaster/tb_conflict/api"
)

type conflict4Impl struct {
    api.INotifier
}

var _ api.Conflict4 = (*conflict4Impl)(nil)
var _ api.INotifier = (*conflict4Impl)(nil)

func NewConflict4(notifier api.INotifier) api.Conflict4 {
	obj := &conflict4Impl{
        INotifier: notifier,
    }
  	return obj
}

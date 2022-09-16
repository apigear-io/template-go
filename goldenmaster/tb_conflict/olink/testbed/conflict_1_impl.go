package testbed

import (
	"goldenmaster/tb_conflict/api"
)

type conflict1Impl struct {
    api.INotifier
    sameName int64
}

var _ api.Conflict1 = (*conflict1Impl)(nil)
var _ api.INotifier = (*conflict1Impl)(nil)

func NewConflict1(notifier api.INotifier) api.Conflict1 {
	obj := &conflict1Impl{
        INotifier: notifier,
        sameName: int64(0),
    }
  	return obj
}
// property get sameName
func (c *conflict1Impl) GetSameName() int64 {
	return c.sameName
}

// property set sameName
func (c *conflict1Impl) SetSameName(sameName int64) {
    c.sameName = sameName
    c.NotifyPropertyChanged("sameName", sameName)
}


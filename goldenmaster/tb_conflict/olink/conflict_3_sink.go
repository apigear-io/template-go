package olink

import (
	"fmt"
	"goldenmaster/tb_conflict/api"
	"olink/pkg/client"
	"olink/pkg/core"
	"sync"
)

type Conflict3Sink struct {
	node       *client.Node
	OnSameName func()
}

var _ client.IObjectSink = (*Conflict3Sink)(nil)

func NewConflict3Sink(node *client.Node) *Conflict3Sink {
	return &Conflict3Sink{
		node: node,
	}
}

func (s *Conflict3Sink) ObjectId() string {
	return "tb.conflict.Conflict3"
}

func (s *Conflict3Sink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.ToMember(signalId)
	switch name {
	case "sameName":
		if s.OnSameName != nil {
			s.OnSameName()
		}
	}
}

func (s *Conflict3Sink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
}

func (s *Conflict3Sink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *Conflict3Sink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

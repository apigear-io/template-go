package olink

import (
    "fmt"
    "olink/pkg/client"
	"olink/pkg/core"
    "goldenmaster/tb_conflict/api"
    "sync"
)


type Conflict4Sink struct {
    node *client.Node    
}

var _ client.IObjectSink = (*Conflict4Sink)(nil)


func NewConflict4Sink(node *client.Node) *Conflict4Sink {
	return &Conflict4Sink{
		node: node,
	}
}

func (s *Conflict4Sink) ObjectId() string {
	return "tb.conflict.Conflict4"
}




func (s *Conflict4Sink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
    name := core.ToMember(signalId)
    switch name {
    }
}


func (s *Conflict4Sink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
        return
    }
    s.node = node
}

func (s *Conflict4Sink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}


func (s *Conflict4Sink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}
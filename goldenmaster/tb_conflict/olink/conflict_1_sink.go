package olink

import (
	"fmt"
	"goldenmaster/tb_conflict/api"
	"olink/pkg/client"
	"olink/pkg/core"
	"sync"
)

type Conflict1Sink struct {
	node     *client.Node
	SameName int64 `json:"sameName"`
}

var _ client.IObjectSink = (*Conflict1Sink)(nil)

func NewConflict1Sink(node *client.Node) *Conflict1Sink {
	return &Conflict1Sink{
		node: node,
	}
}

func (s *Conflict1Sink) ObjectId() string {
	return "tb.conflict.Conflict1"
}

func (s *Conflict1Sink) SetSameName(sameName int64) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "sameName")
	s.node.SetRemoteProperty(propertyId, sameName)
}

func (s *Conflict1Sink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.ToMember(signalId)
	switch name {
	}
}

func (s *Conflict1Sink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["sameName"]; ok {
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetSameName(v)
		}
	}
}

func (s *Conflict1Sink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	case "sameName":
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetSameName(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *Conflict1Sink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

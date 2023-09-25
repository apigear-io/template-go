package olink

import (
	"fmt"
	"goldenmaster/tb_again/api"
	"olink/pkg/client"
	"olink/pkg/core"
	"sync"
)

type SameEnum1InterfaceSink struct {
	node   *client.Node
	Prop1  api.Enum1 `json:"prop1"`
	OnSig1 func()
}

var _ client.IObjectSink = (*SameEnum1InterfaceSink)(nil)

func NewSameEnum1InterfaceSink(node *client.Node) *SameEnum1InterfaceSink {
	return &SameEnum1InterfaceSink{
		node: node,
	}
}

func (s *SameEnum1InterfaceSink) ObjectId() string {
	return "tb.again.SameEnum1Interface"
}

func (s *SameEnum1InterfaceSink) SetProp1(prop1 api.Enum1) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *SameEnum1InterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.ToMember(signalId)
	switch name {
	case "sig1":
		if s.OnSig1 != nil {
			s.OnSig1()
		}
	}
}

func (s *SameEnum1InterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["prop1"]; ok {
		v, err := api.AsEnum1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	}
}

func (s *SameEnum1InterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	case "prop1":
		v, err := api.AsEnum1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *SameEnum1InterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

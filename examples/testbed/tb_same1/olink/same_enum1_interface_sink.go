package olink

import (
	"fmt"
	"testbed/tb_same1/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type SameEnum1InterfaceSink struct {
	node   *client.Node
	Prop1  api.Enum1 `json:"prop1"`
	OnSig1 func(param1 api.Enum1)
}

var _ client.IObjectSink = (*SameEnum1InterfaceSink)(nil)

func NewSameEnum1InterfaceSink(node *client.Node) *SameEnum1InterfaceSink {
	return &SameEnum1InterfaceSink{
		node: node,
	}
}

func (s *SameEnum1InterfaceSink) ObjectId() string {
	return "tb.same1.SameEnum1Interface"
}

func (s *SameEnum1InterfaceSink) SetProp1(prop1 api.Enum1) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *SameEnum1InterfaceSink) Func1(param1 api.Enum1) api.Enum1 {
	var reply api.Enum1
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func1")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *SameEnum1InterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sig1":
		if s.OnSig1 != nil {
			param1 := args[0].(api.Enum1)
			s.OnSig1(param1)
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
	name := core.SymbolIdToMember(propertyId)
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

package olink

import (
	"fmt"
	"testbed/testbed2/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type NestedStruct1InterfaceSink struct {
	node   *client.Node
	Prop1  api.NestedStruct1 `json:"prop1"`
	OnSig1 func(param1 api.NestedStruct1)
}

var _ client.IObjectSink = (*NestedStruct1InterfaceSink)(nil)

func NewNestedStruct1InterfaceSink(node *client.Node) *NestedStruct1InterfaceSink {
	return &NestedStruct1InterfaceSink{
		node: node,
	}
}

func (s *NestedStruct1InterfaceSink) ObjectId() string {
	return "testbed2.NestedStruct1Interface"
}

func (s *NestedStruct1InterfaceSink) SetProp1(prop1 api.NestedStruct1) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *NestedStruct1InterfaceSink) Func1(param1 api.NestedStruct1) api.NestedStruct1 {
	var reply api.NestedStruct1
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func1")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *NestedStruct1InterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sig1":
		if s.OnSig1 != nil {
			param1 := args[0].(api.NestedStruct1)
			s.OnSig1(param1)
		}
	}
}

func (s *NestedStruct1InterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["prop1"]; ok {
		v, err := api.AsNestedStruct1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	}
}

func (s *NestedStruct1InterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop1":
		v, err := api.AsNestedStruct1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *NestedStruct1InterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

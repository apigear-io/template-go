package olink

import (
	"fmt"
	"testbed/tb_same1/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type SameStruct2InterfaceSink struct {
	node   *client.Node
	Prop1  api.Struct1 `json:"prop1"`
	Prop2  api.Struct2 `json:"prop2"`
	OnSig1 func(param1 api.Struct1)
	OnSig2 func(param1 api.Struct1, param2 api.Struct2)
}

var _ client.IObjectSink = (*SameStruct2InterfaceSink)(nil)

func NewSameStruct2InterfaceSink(node *client.Node) *SameStruct2InterfaceSink {
	return &SameStruct2InterfaceSink{
		node: node,
	}
}

func (s *SameStruct2InterfaceSink) ObjectId() string {
	return "tb.same1.SameStruct2Interface"
}

func (s *SameStruct2InterfaceSink) SetProp1(prop1 api.Struct1) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *SameStruct2InterfaceSink) SetProp2(prop2 api.Struct2) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop2")
	s.node.SetRemoteProperty(propertyId, prop2)
}

func (s *SameStruct2InterfaceSink) Func1(param1 api.Struct1) api.Struct1 {
	var reply api.Struct1
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func1")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *SameStruct2InterfaceSink) Func2(param1 api.Struct1, param2 api.Struct2) api.Struct1 {
	var reply api.Struct1
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func2")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *SameStruct2InterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sig1":
		if s.OnSig1 != nil {
			param1 := args[0].(api.Struct1)
			s.OnSig1(param1)
		}
	case "sig2":
		if s.OnSig2 != nil {
			param1 := args[0].(api.Struct1)
			param2 := args[1].(api.Struct2)
			s.OnSig2(param1, param2)
		}
	}
}

func (s *SameStruct2InterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["prop1"]; ok {
		v, err := api.AsStruct1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	}
	if value, ok := props["prop2"]; ok {
		v, err := api.AsStruct2(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	}
}

func (s *SameStruct2InterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop1":
		v, err := api.AsStruct1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	case "prop2":
		v, err := api.AsStruct2(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *SameStruct2InterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

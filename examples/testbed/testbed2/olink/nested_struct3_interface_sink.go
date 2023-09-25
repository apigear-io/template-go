package olink

import (
	"fmt"
	"testbed/testbed2/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type NestedStruct3InterfaceSink struct {
	node   *client.Node
	Prop1  api.NestedStruct1 `json:"prop1"`
	Prop2  api.NestedStruct2 `json:"prop2"`
	Prop3  api.NestedStruct3 `json:"prop3"`
	OnSig1 func(param1 api.NestedStruct1)
	OnSig2 func(param1 api.NestedStruct1, param2 api.NestedStruct2)
	OnSig3 func(param1 api.NestedStruct1, param2 api.NestedStruct2, param3 api.NestedStruct3)
}

var _ client.IObjectSink = (*NestedStruct3InterfaceSink)(nil)

func NewNestedStruct3InterfaceSink(node *client.Node) *NestedStruct3InterfaceSink {
	return &NestedStruct3InterfaceSink{
		node: node,
	}
}

func (s *NestedStruct3InterfaceSink) ObjectId() string {
	return "testbed2.NestedStruct3Interface"
}

func (s *NestedStruct3InterfaceSink) SetProp1(prop1 api.NestedStruct1) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *NestedStruct3InterfaceSink) SetProp2(prop2 api.NestedStruct2) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop2")
	s.node.SetRemoteProperty(propertyId, prop2)
}

func (s *NestedStruct3InterfaceSink) SetProp3(prop3 api.NestedStruct3) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop3")
	s.node.SetRemoteProperty(propertyId, prop3)
}

func (s *NestedStruct3InterfaceSink) Func1(param1 api.NestedStruct1) api.NestedStruct1 {
	var reply api.NestedStruct1
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func1")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *NestedStruct3InterfaceSink) Func2(param1 api.NestedStruct1, param2 api.NestedStruct2) api.NestedStruct1 {
	var reply api.NestedStruct1
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func2")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *NestedStruct3InterfaceSink) Func3(param1 api.NestedStruct1, param2 api.NestedStruct2, param3 api.NestedStruct3) api.NestedStruct1 {
	var reply api.NestedStruct1
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func3")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *NestedStruct3InterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sig1":
		if s.OnSig1 != nil {
			param1 := args[0].(api.NestedStruct1)
			s.OnSig1(param1)
		}
	case "sig2":
		if s.OnSig2 != nil {
			param1 := args[0].(api.NestedStruct1)
			param2 := args[1].(api.NestedStruct2)
			s.OnSig2(param1, param2)
		}
	case "sig3":
		if s.OnSig3 != nil {
			param1 := args[0].(api.NestedStruct1)
			param2 := args[1].(api.NestedStruct2)
			param3 := args[2].(api.NestedStruct3)
			s.OnSig3(param1, param2, param3)
		}
	}
}

func (s *NestedStruct3InterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
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
	if value, ok := props["prop2"]; ok {
		v, err := api.AsNestedStruct2(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	}
	if value, ok := props["prop3"]; ok {
		v, err := api.AsNestedStruct3(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp3(v)
		}
	}
}

func (s *NestedStruct3InterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
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
	case "prop2":
		v, err := api.AsNestedStruct2(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	case "prop3":
		v, err := api.AsNestedStruct3(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp3(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *NestedStruct3InterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

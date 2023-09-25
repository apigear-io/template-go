package olink

import (
	"fmt"
	"testbed/testbed2/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type ManyParamInterfaceSink struct {
	node   *client.Node
	Prop1  int32 `json:"prop1"`
	Prop2  int32 `json:"prop2"`
	Prop3  int32 `json:"prop3"`
	Prop4  int32 `json:"prop4"`
	OnSig1 func(param1 int32)
	OnSig2 func(param1 int32, param2 int32)
	OnSig3 func(param1 int32, param2 int32, param3 int32)
	OnSig4 func(param1 int32, param2 int32, param3 int32, param4 int32)
}

var _ client.IObjectSink = (*ManyParamInterfaceSink)(nil)

func NewManyParamInterfaceSink(node *client.Node) *ManyParamInterfaceSink {
	return &ManyParamInterfaceSink{
		node: node,
	}
}

func (s *ManyParamInterfaceSink) ObjectId() string {
	return "testbed2.ManyParamInterface"
}

func (s *ManyParamInterfaceSink) SetProp1(prop1 int32) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *ManyParamInterfaceSink) SetProp2(prop2 int32) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop2")
	s.node.SetRemoteProperty(propertyId, prop2)
}

func (s *ManyParamInterfaceSink) SetProp3(prop3 int32) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop3")
	s.node.SetRemoteProperty(propertyId, prop3)
}

func (s *ManyParamInterfaceSink) SetProp4(prop4 int32) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "prop4")
	s.node.SetRemoteProperty(propertyId, prop4)
}

func (s *ManyParamInterfaceSink) Func1(param1 int32) int32 {
	var reply int32
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func1")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *ManyParamInterfaceSink) Func2(param1 int32, param2 int32) int32 {
	var reply int32
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func2")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *ManyParamInterfaceSink) Func3(param1 int32, param2 int32, param3 int32) int32 {
	var reply int32
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func3")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *ManyParamInterfaceSink) Func4(param1 int32, param2 int32, param3 int32, param4 int32) int32 {
	var reply int32
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "func4")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *ManyParamInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sig1":
		if s.OnSig1 != nil {
			param1 := args[0].(int32)
			s.OnSig1(param1)
		}
	case "sig2":
		if s.OnSig2 != nil {
			param1 := args[0].(int32)
			param2 := args[1].(int32)
			s.OnSig2(param1, param2)
		}
	case "sig3":
		if s.OnSig3 != nil {
			param1 := args[0].(int32)
			param2 := args[1].(int32)
			param3 := args[2].(int32)
			s.OnSig3(param1, param2, param3)
		}
	case "sig4":
		if s.OnSig4 != nil {
			param1 := args[0].(int32)
			param2 := args[1].(int32)
			param3 := args[2].(int32)
			param4 := args[3].(int32)
			s.OnSig4(param1, param2, param3, param4)
		}
	}
}

func (s *ManyParamInterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["prop1"]; ok {
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	}
	if value, ok := props["prop2"]; ok {
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	}
	if value, ok := props["prop3"]; ok {
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp3(v)
		}
	}
	if value, ok := props["prop4"]; ok {
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp4(v)
		}
	}
}

func (s *ManyParamInterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop1":
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	case "prop2":
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	case "prop3":
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp3(v)
		}
	case "prop4":
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp4(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *ManyParamInterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

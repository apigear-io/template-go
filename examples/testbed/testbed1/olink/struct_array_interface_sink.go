package olink

import (
	"fmt"
	"testbed/testbed1/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type StructArrayInterfaceSink struct {
	node        *client.Node
	PropBool    []api.StructBool   `json:"propBool"`
	PropInt     []api.StructInt    `json:"propInt"`
	PropFloat   []api.StructFloat  `json:"propFloat"`
	PropString  []api.StructString `json:"propString"`
	OnSigBool   func(paramBool []api.StructBool)
	OnSigInt    func(paramInt []api.StructInt)
	OnSigFloat  func(paramFloat []api.StructFloat)
	OnSigString func(paramString []api.StructString)
}

var _ client.IObjectSink = (*StructArrayInterfaceSink)(nil)

func NewStructArrayInterfaceSink(node *client.Node) *StructArrayInterfaceSink {
	return &StructArrayInterfaceSink{
		node: node,
	}
}

func (s *StructArrayInterfaceSink) ObjectId() string {
	return "testbed1.StructArrayInterface"
}

func (s *StructArrayInterfaceSink) SetPropBool(propBool []api.StructBool) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propBool")
	s.node.SetRemoteProperty(propertyId, propBool)
}

func (s *StructArrayInterfaceSink) SetPropInt(propInt []api.StructInt) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propInt")
	s.node.SetRemoteProperty(propertyId, propInt)
}

func (s *StructArrayInterfaceSink) SetPropFloat(propFloat []api.StructFloat) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propFloat")
	s.node.SetRemoteProperty(propertyId, propFloat)
}

func (s *StructArrayInterfaceSink) SetPropString(propString []api.StructString) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propString")
	s.node.SetRemoteProperty(propertyId, propString)
}

func (s *StructArrayInterfaceSink) FuncBool(paramBool []api.StructBool) []api.StructBool {
	var reply []api.StructBool
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcBool")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructArrayInterfaceSink) FuncInt(paramInt []api.StructInt) []api.StructInt {
	var reply []api.StructInt
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcInt")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructArrayInterfaceSink) FuncFloat(paramFloat []api.StructFloat) []api.StructFloat {
	var reply []api.StructFloat
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcFloat")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructArrayInterfaceSink) FuncString(paramString []api.StructString) []api.StructString {
	var reply []api.StructString
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcString")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructArrayInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sigBool":
		if s.OnSigBool != nil {
			paramBool := args[0].([]api.StructBool)
			s.OnSigBool(paramBool)
		}
	case "sigInt":
		if s.OnSigInt != nil {
			paramInt := args[0].([]api.StructInt)
			s.OnSigInt(paramInt)
		}
	case "sigFloat":
		if s.OnSigFloat != nil {
			paramFloat := args[0].([]api.StructFloat)
			s.OnSigFloat(paramFloat)
		}
	case "sigString":
		if s.OnSigString != nil {
			paramString := args[0].([]api.StructString)
			s.OnSigString(paramString)
		}
	}
}

func (s *StructArrayInterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["propBool"]; ok {
		v, err := api.AsStructBoolArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropBool(v)
		}
	}
	if value, ok := props["propInt"]; ok {
		v, err := api.AsStructIntArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropInt(v)
		}
	}
	if value, ok := props["propFloat"]; ok {
		v, err := api.AsStructFloatArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropFloat(v)
		}
	}
	if value, ok := props["propString"]; ok {
		v, err := api.AsStructStringArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropString(v)
		}
	}
}

func (s *StructArrayInterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "propBool":
		v, err := api.AsStructBoolArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropBool(v)
		}
	case "propInt":
		v, err := api.AsStructIntArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropInt(v)
		}
	case "propFloat":
		v, err := api.AsStructFloatArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropFloat(v)
		}
	case "propString":
		v, err := api.AsStructStringArray(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropString(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *StructArrayInterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

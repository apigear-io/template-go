package olink

import (
	"fmt"
	"testbed/testbed1/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type StructInterfaceSink struct {
	node        *client.Node
	PropBool    api.StructBool   `json:"propBool"`
	PropInt     api.StructInt    `json:"propInt"`
	PropFloat   api.StructFloat  `json:"propFloat"`
	PropString  api.StructString `json:"propString"`
	OnSigBool   func(paramBool api.StructBool)
	OnSigInt    func(paramInt api.StructInt)
	OnSigFloat  func(paramFloat api.StructFloat)
	OnSigString func(paramString api.StructString)
}

var _ client.IObjectSink = (*StructInterfaceSink)(nil)

func NewStructInterfaceSink(node *client.Node) *StructInterfaceSink {
	return &StructInterfaceSink{
		node: node,
	}
}

func (s *StructInterfaceSink) ObjectId() string {
	return "testbed1.StructInterface"
}

func (s *StructInterfaceSink) SetPropBool(propBool api.StructBool) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propBool")
	s.node.SetRemoteProperty(propertyId, propBool)
}

func (s *StructInterfaceSink) SetPropInt(propInt api.StructInt) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propInt")
	s.node.SetRemoteProperty(propertyId, propInt)
}

func (s *StructInterfaceSink) SetPropFloat(propFloat api.StructFloat) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propFloat")
	s.node.SetRemoteProperty(propertyId, propFloat)
}

func (s *StructInterfaceSink) SetPropString(propString api.StructString) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propString")
	s.node.SetRemoteProperty(propertyId, propString)
}

func (s *StructInterfaceSink) FuncBool(paramBool api.StructBool) api.StructBool {
	var reply api.StructBool
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcBool")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructInterfaceSink) FuncInt(paramInt api.StructInt) api.StructInt {
	var reply api.StructInt
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcInt")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructInterfaceSink) FuncFloat(paramFloat api.StructFloat) api.StructFloat {
	var reply api.StructFloat
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcFloat")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructInterfaceSink) FuncString(paramString api.StructString) api.StructString {
	var reply api.StructString
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcString")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *StructInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sigBool":
		if s.OnSigBool != nil {
			paramBool := args[0].(api.StructBool)
			s.OnSigBool(paramBool)
		}
	case "sigInt":
		if s.OnSigInt != nil {
			paramInt := args[0].(api.StructInt)
			s.OnSigInt(paramInt)
		}
	case "sigFloat":
		if s.OnSigFloat != nil {
			paramFloat := args[0].(api.StructFloat)
			s.OnSigFloat(paramFloat)
		}
	case "sigString":
		if s.OnSigString != nil {
			paramString := args[0].(api.StructString)
			s.OnSigString(paramString)
		}
	}
}

func (s *StructInterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["propBool"]; ok {
		v, err := api.AsStructBool(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropBool(v)
		}
	}
	if value, ok := props["propInt"]; ok {
		v, err := api.AsStructInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropInt(v)
		}
	}
	if value, ok := props["propFloat"]; ok {
		v, err := api.AsStructFloat(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropFloat(v)
		}
	}
	if value, ok := props["propString"]; ok {
		v, err := api.AsStructString(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropString(v)
		}
	}
}

func (s *StructInterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "propBool":
		v, err := api.AsStructBool(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropBool(v)
		}
	case "propInt":
		v, err := api.AsStructInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropInt(v)
		}
	case "propFloat":
		v, err := api.AsStructFloat(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropFloat(v)
		}
	case "propString":
		v, err := api.AsStructString(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropString(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *StructInterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

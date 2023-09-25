package olink

import (
	"fmt"
	"testbed/tb_simple/api"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type SimpleInterfaceSink struct {
	node        *client.Node
	PropBool    bool    `json:"propBool"`
	PropInt     int32   `json:"propInt"`
	PropFloat   float32 `json:"propFloat"`
	PropString  string  `json:"propString"`
	OnSigBool   func(paramBool bool)
	OnSigInt    func(paramInt int32)
	OnSigFloat  func(paramFloat float32)
	OnSigString func(paramString string)
}

var _ client.IObjectSink = (*SimpleInterfaceSink)(nil)

func NewSimpleInterfaceSink(node *client.Node) *SimpleInterfaceSink {
	return &SimpleInterfaceSink{
		node: node,
	}
}

func (s *SimpleInterfaceSink) ObjectId() string {
	return "tb.simple.SimpleInterface"
}

func (s *SimpleInterfaceSink) SetPropBool(propBool bool) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propBool")
	s.node.SetRemoteProperty(propertyId, propBool)
}

func (s *SimpleInterfaceSink) SetPropInt(propInt int32) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propInt")
	s.node.SetRemoteProperty(propertyId, propInt)
}

func (s *SimpleInterfaceSink) SetPropFloat(propFloat float32) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propFloat")
	s.node.SetRemoteProperty(propertyId, propFloat)
}

func (s *SimpleInterfaceSink) SetPropString(propString string) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeSymbolId(s.ObjectId(), "propString")
	s.node.SetRemoteProperty(propertyId, propString)
}

func (s *SimpleInterfaceSink) FuncBool(paramBool bool) bool {
	var reply bool
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcBool")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *SimpleInterfaceSink) FuncInt(paramInt int32) int32 {
	var reply int32
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcInt")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *SimpleInterfaceSink) FuncFloat(paramFloat float32) float32 {
	var reply float32
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcFloat")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *SimpleInterfaceSink) FuncString(paramString string) string {
	var reply string
	if s.node != nil {
		methodId := core.MakeSymbolId(s.ObjectId(), "funcString")
		s.node.InvokeRemote(methodId, core.Args{}, nil)
	}
	return reply
}

func (s *SimpleInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.SymbolIdToMember(signalId)
	switch name {
	case "sigBool":
		if s.OnSigBool != nil {
			paramBool := args[0].(bool)
			s.OnSigBool(paramBool)
		}
	case "sigInt":
		if s.OnSigInt != nil {
			paramInt := args[0].(int32)
			s.OnSigInt(paramInt)
		}
	case "sigFloat":
		if s.OnSigFloat != nil {
			paramFloat := args[0].(float32)
			s.OnSigFloat(paramFloat)
		}
	case "sigString":
		if s.OnSigString != nil {
			paramString := args[0].(string)
			s.OnSigString(paramString)
		}
	}
}

func (s *SimpleInterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["propBool"]; ok {
		v, err := api.AsBool(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropBool(v)
		}
	}
	if value, ok := props["propInt"]; ok {
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropInt(v)
		}
	}
	if value, ok := props["propFloat"]; ok {
		v, err := api.AsFloat(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropFloat(v)
		}
	}
	if value, ok := props["propString"]; ok {
		v, err := api.AsString(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropString(v)
		}
	}
}

func (s *SimpleInterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "propBool":
		v, err := api.AsBool(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropBool(v)
		}
	case "propInt":
		v, err := api.AsInt(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropInt(v)
		}
	case "propFloat":
		v, err := api.AsFloat(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropFloat(v)
		}
	case "propString":
		v, err := api.AsString(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetPropString(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *SimpleInterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

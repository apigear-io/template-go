package olink

import (
	"fmt"
	"goldenmaster/tb_simple/api"
	"olink/pkg/client"
	"olink/pkg/core"
	"sync"
)

type SimpleInterfaceSink struct {
	node        *client.Node
	PropBool    bool    `json:"propBool"`
	PropInt     int64   `json:"propInt"`
	PropFloat   float64 `json:"propFloat"`
	PropString  string  `json:"propString"`
	OnSigBool   func(paramBool bool)
	OnSigInt    func(paramInt int64)
	OnSigFloat  func(paramFloat float64)
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
	propertyId := core.MakeIdentifier(s.ObjectId(), "propBool")
	s.node.SetRemoteProperty(propertyId, propBool)
}

func (s *SimpleInterfaceSink) SetPropInt(propInt int64) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "propInt")
	s.node.SetRemoteProperty(propertyId, propInt)
}

func (s *SimpleInterfaceSink) SetPropFloat(propFloat float64) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "propFloat")
	s.node.SetRemoteProperty(propertyId, propFloat)
}

func (s *SimpleInterfaceSink) SetPropString(propString string) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "propString")
	s.node.SetRemoteProperty(propertyId, propString)
}

func (s *SimpleInterfaceSink) FuncBool(paramBool bool) bool {
	var reply bool
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "funcBool")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{paramBool}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(bool)
		})
		wg.Wait()
	}
	return reply
}

func (s *SimpleInterfaceSink) FuncInt(paramInt int64) int64 {
	var reply int64
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "funcInt")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{paramInt}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(int64)
		})
		wg.Wait()
	}
	return reply
}

func (s *SimpleInterfaceSink) FuncFloat(paramFloat float64) float64 {
	var reply float64
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "funcFloat")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{paramFloat}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(float64)
		})
		wg.Wait()
	}
	return reply
}

func (s *SimpleInterfaceSink) FuncString(paramString string) string {
	var reply string
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "funcString")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{paramString}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(string)
		})
		wg.Wait()
	}
	return reply
}

func (s *SimpleInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.ToMember(signalId)
	switch name {
	case "sigBool":
		if s.OnSigBool != nil {
			paramBool := args[0].(bool)
			s.OnSigBool(paramBool)
		}
	case "sigInt":
		if s.OnSigInt != nil {
			paramInt := args[0].(int64)
			s.OnSigInt(paramInt)
		}
	case "sigFloat":
		if s.OnSigFloat != nil {
			paramFloat := args[0].(float64)
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
	name := core.ToMember(propertyId)
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

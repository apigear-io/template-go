package olink

import (
	"fmt"
	"goldenmaster/tb_enum/api"
	"olink/pkg/client"
	"olink/pkg/core"
	"sync"
)

type EnumInterfaceSink struct {
	node   *client.Node
	Prop0  api.Enum0 `json:"prop0"`
	Prop1  api.Enum1 `json:"prop1"`
	Prop2  api.Enum2 `json:"prop2"`
	Prop3  api.Enum3 `json:"prop3"`
	OnSig0 func(param0 api.Enum0)
	OnSig1 func(param1 api.Enum1)
	OnSig2 func(param2 api.Enum2)
	OnSig3 func(param3 api.Enum3)
}

var _ client.IObjectSink = (*EnumInterfaceSink)(nil)

func NewEnumInterfaceSink(node *client.Node) *EnumInterfaceSink {
	return &EnumInterfaceSink{
		node: node,
	}
}

func (s *EnumInterfaceSink) ObjectId() string {
	return "tb.enum.EnumInterface"
}

func (s *EnumInterfaceSink) SetProp0(prop0 api.Enum0) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop0")
	s.node.SetRemoteProperty(propertyId, prop0)
}

func (s *EnumInterfaceSink) SetProp1(prop1 api.Enum1) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *EnumInterfaceSink) SetProp2(prop2 api.Enum2) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop2")
	s.node.SetRemoteProperty(propertyId, prop2)
}

func (s *EnumInterfaceSink) SetProp3(prop3 api.Enum3) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop3")
	s.node.SetRemoteProperty(propertyId, prop3)
}

func (s *EnumInterfaceSink) Func0(param0 api.Enum0) api.Enum0 {
	var reply api.Enum0
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "func0")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{param0}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(api.Enum0)
		})
		wg.Wait()
	}
	return reply
}

func (s *EnumInterfaceSink) Func1(param1 api.Enum1) api.Enum1 {
	var reply api.Enum1
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "func1")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{param1}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(api.Enum1)
		})
		wg.Wait()
	}
	return reply
}

func (s *EnumInterfaceSink) Func2(param2 api.Enum2) api.Enum2 {
	var reply api.Enum2
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "func2")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{param2}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(api.Enum2)
		})
		wg.Wait()
	}
	return reply
}

func (s *EnumInterfaceSink) Func3(param3 api.Enum3) api.Enum3 {
	var reply api.Enum3
	if s.node != nil {
		methodId := core.MakeIdentifier(s.ObjectId(), "func3")
		wg := sync.WaitGroup{}
		wg.Add(1)
		args := core.Args{param3}
		s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			wg.Done()
			reply = arg.Value.(api.Enum3)
		})
		wg.Wait()
	}
	return reply
}

func (s *EnumInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.ToMember(signalId)
	switch name {
	case "sig0":
		if s.OnSig0 != nil {
			param0 := args[0].(api.Enum0)
			s.OnSig0(param0)
		}
	case "sig1":
		if s.OnSig1 != nil {
			param1 := args[0].(api.Enum1)
			s.OnSig1(param1)
		}
	case "sig2":
		if s.OnSig2 != nil {
			param2 := args[0].(api.Enum2)
			s.OnSig2(param2)
		}
	case "sig3":
		if s.OnSig3 != nil {
			param3 := args[0].(api.Enum3)
			s.OnSig3(param3)
		}
	}
}

func (s *EnumInterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
		return
	}
	s.node = node
	if value, ok := props["prop0"]; ok {
		v, err := api.AsEnum0(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp0(v)
		}
	}
	if value, ok := props["prop1"]; ok {
		v, err := api.AsEnum1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	}
	if value, ok := props["prop2"]; ok {
		v, err := api.AsEnum2(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	}
	if value, ok := props["prop3"]; ok {
		v, err := api.AsEnum3(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp3(v)
		}
	}
}

func (s *EnumInterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	case "prop0":
		v, err := api.AsEnum0(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp0(v)
		}
	case "prop1":
		v, err := api.AsEnum1(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp1(v)
		}
	case "prop2":
		v, err := api.AsEnum2(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp2(v)
		}
	case "prop3":
		v, err := api.AsEnum3(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			s.SetProp3(v)
		}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}

func (s *EnumInterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}

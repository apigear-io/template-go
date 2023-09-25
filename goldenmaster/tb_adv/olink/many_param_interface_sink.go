package olink

import (
	"fmt"
	"goldenmaster/tb_adv/api"
	"olink/pkg/client"
	"olink/pkg/core"
	"sync"
)

type ManyParamInterfaceSink struct {
	node   *client.Node
	Prop1  int64 `json:"prop1"`
	Prop2  int64 `json:"prop2"`
	Prop3  int64 `json:"prop3"`
	Prop4  int64 `json:"prop4"`
	OnSig1 func()
	OnSig2 func()
	OnSig3 func()
	OnSig4 func()
}

var _ client.IObjectSink = (*ManyParamInterfaceSink)(nil)

func NewManyParamInterfaceSink(node *client.Node) *ManyParamInterfaceSink {
	return &ManyParamInterfaceSink{
		node: node,
	}
}

func (s *ManyParamInterfaceSink) ObjectId() string {
	return "tb.adv.ManyParamInterface"
}

func (s *ManyParamInterfaceSink) SetProp1(prop1 int64) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop1")
	s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *ManyParamInterfaceSink) SetProp2(prop2 int64) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop2")
	s.node.SetRemoteProperty(propertyId, prop2)
}

func (s *ManyParamInterfaceSink) SetProp3(prop3 int64) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop3")
	s.node.SetRemoteProperty(propertyId, prop3)
}

func (s *ManyParamInterfaceSink) SetProp4(prop4 int64) {
	if s.node == nil {
		return
	}
	propertyId := core.MakeIdentifier(s.ObjectId(), "prop4")
	s.node.SetRemoteProperty(propertyId, prop4)
}

func (s *ManyParamInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
	name := core.ToMember(signalId)
	switch name {
	case "sig1":
		if s.OnSig1 != nil {
			s.OnSig1()
		}
	case "sig2":
		if s.OnSig2 != nil {
			s.OnSig2()
		}
	case "sig3":
		if s.OnSig3 != nil {
			s.OnSig3()
		}
	case "sig4":
		if s.OnSig4 != nil {
			s.OnSig4()
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
	name := core.ToMember(propertyId)
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

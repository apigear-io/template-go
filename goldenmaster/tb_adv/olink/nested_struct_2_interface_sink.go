package olink

import (
    "fmt"
    "olink/pkg/client"
	"olink/pkg/core"
    "goldenmaster/tb_adv/api"
    "sync"
)


type NestedStruct2InterfaceSink struct {
    node *client.Node
    Prop1 api.NestedStruct1 `json:"prop1"`
    Prop2 api.NestedStruct2 `json:"prop2"`
    OnSig1 func()
    OnSig2 func()    
}

var _ client.IObjectSink = (*NestedStruct2InterfaceSink)(nil)


func NewNestedStruct2InterfaceSink(node *client.Node) *NestedStruct2InterfaceSink {
	return &NestedStruct2InterfaceSink{
		node: node,
	}
}

func (s *NestedStruct2InterfaceSink) ObjectId() string {
	return "tb.adv.NestedStruct2Interface"
}


func (s *NestedStruct2InterfaceSink) SetProp1(prop1 api.NestedStruct1) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "prop1")
    s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *NestedStruct2InterfaceSink) SetProp2(prop2 api.NestedStruct2) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "prop2")
    s.node.SetRemoteProperty(propertyId, prop2)
}



func (s *NestedStruct2InterfaceSink) OnSignal(signalId string, args core.Args) {
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
    }
}


func (s *NestedStruct2InterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
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
}

func (s *NestedStruct2InterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
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
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}


func (s *NestedStruct2InterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}
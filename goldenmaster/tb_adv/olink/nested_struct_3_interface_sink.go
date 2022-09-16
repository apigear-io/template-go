package olink

import (
    "fmt"
    "olink/pkg/client"
	"olink/pkg/core"
    "goldenmaster/tb_adv/api"
    "sync"
)


type NestedStruct3InterfaceSink struct {
    node *client.Node
    Prop1 api.NestedStruct1 `json:"prop1"`
    Prop2 api.NestedStruct2 `json:"prop2"`
    Prop3 api.NestedStruct3 `json:"prop3"`
    OnSig1 func()
    OnSig2 func()
    OnSig3 func()    
}

var _ client.IObjectSink = (*NestedStruct3InterfaceSink)(nil)


func NewNestedStruct3InterfaceSink(node *client.Node) *NestedStruct3InterfaceSink {
	return &NestedStruct3InterfaceSink{
		node: node,
	}
}

func (s *NestedStruct3InterfaceSink) ObjectId() string {
	return "tb.adv.NestedStruct3Interface"
}


func (s *NestedStruct3InterfaceSink) SetProp1(prop1 api.NestedStruct1) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "prop1")
    s.node.SetRemoteProperty(propertyId, prop1)
}

func (s *NestedStruct3InterfaceSink) SetProp2(prop2 api.NestedStruct2) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "prop2")
    s.node.SetRemoteProperty(propertyId, prop2)
}

func (s *NestedStruct3InterfaceSink) SetProp3(prop3 api.NestedStruct3) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "prop3")
    s.node.SetRemoteProperty(propertyId, prop3)
}



func (s *NestedStruct3InterfaceSink) OnSignal(signalId string, args core.Args) {
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
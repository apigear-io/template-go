package olink

import (
    "fmt"
    "olink/pkg/client"
	"olink/pkg/core"
    "goldenmaster/tb_again/api"
    "sync"
)


type SameStruct1InterfaceSink struct {
    node *client.Node
    Prop1 api.Struct1 `json:"prop1"`
    OnSig1 func()    
}

var _ client.IObjectSink = (*SameStruct1InterfaceSink)(nil)


func NewSameStruct1InterfaceSink(node *client.Node) *SameStruct1InterfaceSink {
	return &SameStruct1InterfaceSink{
		node: node,
	}
}

func (s *SameStruct1InterfaceSink) ObjectId() string {
	return "tb.again.SameStruct1Interface"
}


func (s *SameStruct1InterfaceSink) SetProp1(prop1 api.Struct1) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "prop1")
    s.node.SetRemoteProperty(propertyId, prop1)
}



func (s *SameStruct1InterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
    name := core.ToMember(signalId)
    switch name {
    case "sig1":
        if s.OnSig1 != nil {
            s.OnSig1()
        }
    }
}


func (s *SameStruct1InterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
        return
    }
    s.node = node
    if value, ok := props["prop1"]; ok {
        v, err := api.AsStruct1(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetProp1(v)
        }
    }
}

func (s *SameStruct1InterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	case "prop1":
        v, err := api.AsStruct1(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetProp1(v)
        }
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}


func (s *SameStruct1InterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}
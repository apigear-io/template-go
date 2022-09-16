package olink

import (
    "fmt"
    "olink/pkg/client"
	"olink/pkg/core"
    "goldenmaster/tb_conflict/api"
    "sync"
)


type Conflict2Sink struct {
    node *client.Node
    SameName int64 `json:"sameName"`
    OnSameName func()    
}

var _ client.IObjectSink = (*Conflict2Sink)(nil)


func NewConflict2Sink(node *client.Node) *Conflict2Sink {
	return &Conflict2Sink{
		node: node,
	}
}

func (s *Conflict2Sink) ObjectId() string {
	return "tb.conflict.Conflict2"
}


func (s *Conflict2Sink) SetSameName(sameName int64) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "sameName")
    s.node.SetRemoteProperty(propertyId, sameName)
}



func (s *Conflict2Sink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
    name := core.ToMember(signalId)
    switch name {
    case "sameName":
        if s.OnSameName != nil {
            s.OnSameName()
        }
    }
}


func (s *Conflict2Sink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
        return
    }
    s.node = node
    if value, ok := props["sameName"]; ok {
        v, err := api.AsInt(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetSameName(v)
        }
    }
}

func (s *Conflict2Sink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	case "sameName":
        v, err := api.AsInt(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetSameName(v)
        }
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}


func (s *Conflict2Sink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}
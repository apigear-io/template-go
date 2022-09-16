package olink

import (
    "fmt"
    "olink/pkg/client"
	"olink/pkg/core"
    "goldenmaster/tb_simple/api"
    "sync"
)


type SimpleArrayInterfaceSink struct {
    node *client.Node
    PropBool []bool `json:"propBool"`
    PropInt []int64 `json:"propInt"`
    PropFloat []float64 `json:"propFloat"`
    PropString []string `json:"propString"`
    OnSigBool func(paramBool []bool)
    OnSigInt func(paramInt []int64)
    OnSigFloat func(paramFloat []float64)
    OnSigString func(paramString []string)    
}

var _ client.IObjectSink = (*SimpleArrayInterfaceSink)(nil)


func NewSimpleArrayInterfaceSink(node *client.Node) *SimpleArrayInterfaceSink {
	return &SimpleArrayInterfaceSink{
		node: node,
	}
}

func (s *SimpleArrayInterfaceSink) ObjectId() string {
	return "tb.simple.SimpleArrayInterface"
}


func (s *SimpleArrayInterfaceSink) SetPropBool(propBool []bool) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propBool")
    s.node.SetRemoteProperty(propertyId, propBool)
}

func (s *SimpleArrayInterfaceSink) SetPropInt(propInt []int64) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propInt")
    s.node.SetRemoteProperty(propertyId, propInt)
}

func (s *SimpleArrayInterfaceSink) SetPropFloat(propFloat []float64) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propFloat")
    s.node.SetRemoteProperty(propertyId, propFloat)
}

func (s *SimpleArrayInterfaceSink) SetPropString(propString []string) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propString")
    s.node.SetRemoteProperty(propertyId, propString)
}

func (s *SimpleArrayInterfaceSink) FuncBool(paramBool []bool) []bool {
    var reply []bool    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcBool")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramBool }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.([]bool)
        })
        wg.Wait()
    }
    return reply
}

func (s *SimpleArrayInterfaceSink) FuncInt(paramInt []int64) []int64 {
    var reply []int64    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcInt")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramInt }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.([]int64)
        })
        wg.Wait()
    }
    return reply
}

func (s *SimpleArrayInterfaceSink) FuncFloat(paramFloat []float64) []float64 {
    var reply []float64    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcFloat")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramFloat }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.([]float64)
        })
        wg.Wait()
    }
    return reply
}

func (s *SimpleArrayInterfaceSink) FuncString(paramString []string) []string {
    var reply []string    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcString")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramString }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.([]string)
        })
        wg.Wait()
    }
    return reply
}



func (s *SimpleArrayInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
    name := core.ToMember(signalId)
    switch name {
    case "sigBool":
        if s.OnSigBool != nil {
            paramBool := args[0].([]bool)
            s.OnSigBool(paramBool)
        }
    case "sigInt":
        if s.OnSigInt != nil {
            paramInt := args[0].([]int64)
            s.OnSigInt(paramInt)
        }
    case "sigFloat":
        if s.OnSigFloat != nil {
            paramFloat := args[0].([]float64)
            s.OnSigFloat(paramFloat)
        }
    case "sigString":
        if s.OnSigString != nil {
            paramString := args[0].([]string)
            s.OnSigString(paramString)
        }
    }
}


func (s *SimpleArrayInterfaceSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
        return
    }
    s.node = node
    if value, ok := props["propBool"]; ok {
        v, err := api.AsBoolArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropBool(v)
        }
    }
    if value, ok := props["propInt"]; ok {
        v, err := api.AsIntArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropInt(v)
        }
    }
    if value, ok := props["propFloat"]; ok {
        v, err := api.AsFloatArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropFloat(v)
        }
    }
    if value, ok := props["propString"]; ok {
        v, err := api.AsStringArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropString(v)
        }
    }
}

func (s *SimpleArrayInterfaceSink) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.ToMember(propertyId)
	switch name {
	case "propBool":
        v, err := api.AsBoolArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropBool(v)
        }
	case "propInt":
        v, err := api.AsIntArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropInt(v)
        }
	case "propFloat":
        v, err := api.AsFloatArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropFloat(v)
        }
	case "propString":
        v, err := api.AsStringArray(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.SetPropString(v)
        }
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}


func (s *SimpleArrayInterfaceSink) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}
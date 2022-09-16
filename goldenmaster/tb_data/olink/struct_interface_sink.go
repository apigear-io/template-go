package olink

import (
    "fmt"
    "olink/pkg/client"
	"olink/pkg/core"
    "goldenmaster/tb_data/api"
    "sync"
)


type StructInterfaceSink struct {
    node *client.Node
    PropBool api.StructBool `json:"propBool"`
    PropInt api.StructInt `json:"propInt"`
    PropFloat api.StructFloat `json:"propFloat"`
    PropString api.StructString `json:"propString"`
    OnSigBool func(paramBool api.StructBool)
    OnSigInt func(paramInt api.StructInt)
    OnSigFloat func(paramFloat api.StructFloat)
    OnSigString func(paramString api.StructString)    
}

var _ client.IObjectSink = (*StructInterfaceSink)(nil)


func NewStructInterfaceSink(node *client.Node) *StructInterfaceSink {
	return &StructInterfaceSink{
		node: node,
	}
}

func (s *StructInterfaceSink) ObjectId() string {
	return "tb.data.StructInterface"
}


func (s *StructInterfaceSink) SetPropBool(propBool api.StructBool) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propBool")
    s.node.SetRemoteProperty(propertyId, propBool)
}

func (s *StructInterfaceSink) SetPropInt(propInt api.StructInt) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propInt")
    s.node.SetRemoteProperty(propertyId, propInt)
}

func (s *StructInterfaceSink) SetPropFloat(propFloat api.StructFloat) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propFloat")
    s.node.SetRemoteProperty(propertyId, propFloat)
}

func (s *StructInterfaceSink) SetPropString(propString api.StructString) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeIdentifier(s.ObjectId(), "propString")
    s.node.SetRemoteProperty(propertyId, propString)
}

func (s *StructInterfaceSink) FuncBool(paramBool api.StructBool) api.StructBool {
    var reply api.StructBool    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcBool")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramBool }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.(api.StructBool)
        })
        wg.Wait()
    }
    return reply
}

func (s *StructInterfaceSink) FuncInt(paramInt api.StructInt) api.StructInt {
    var reply api.StructInt    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcInt")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramInt }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.(api.StructInt)
        })
        wg.Wait()
    }
    return reply
}

func (s *StructInterfaceSink) FuncFloat(paramFloat api.StructFloat) api.StructFloat {
    var reply api.StructFloat    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcFloat")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramFloat }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.(api.StructFloat)
        })
        wg.Wait()
    }
    return reply
}

func (s *StructInterfaceSink) FuncString(paramString api.StructString) api.StructString {
    var reply api.StructString    
    if s.node != nil {
        methodId := core.MakeIdentifier(s.ObjectId(), "funcString")
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ paramString }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.(api.StructString)
        })
        wg.Wait()
    }
    return reply
}



func (s *StructInterfaceSink) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
    name := core.ToMember(signalId)
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
	name := core.ToMember(propertyId)
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
package olink

import (
    "fmt"
    "log"
	"olink/pkg/core"
    "olink/pkg/remote"
    "goldenmaster/tb_adv/api"
)

type NestedStruct2InterfaceSource struct {
    node *remote.Node
    impl api.NestedStruct2Interface
}

var _ remote.IObjectSource = (*NestedStruct2InterfaceSource)(nil)
var _ api.INotifier = (*NestedStruct2InterfaceSource)(nil)

func NewNestedStruct2InterfaceSource() *NestedStruct2InterfaceSource {
    return &NestedStruct2InterfaceSource{}
}

func (s *NestedStruct2InterfaceSource) SetImplementation(impl api.NestedStruct2Interface) {
    s.impl = impl
}

func (s *NestedStruct2InterfaceSource) ObjectId() string {
	return "tb.adv.NestedStruct2Interface"
}

func (s *NestedStruct2InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.ToMember(methodId)
	switch name {
    default:
        return nil, fmt.Errorf("unknown method: %s", name)
    }
}

func (s *NestedStruct2InterfaceSource) SetProperty(propertyId string, value core.Any) error {
    if s.impl == nil {
        return fmt.Errorf("no implementation")
    }
    name := core.ToMember(propertyId)
    switch name {
    case "prop1":
        prop, err := api.AsNestedStruct1(value)
        if err != nil {
            return err
        }
        s.impl.SetProp1(prop)
    case "prop2":
        prop, err := api.AsNestedStruct2(value)
        if err != nil {
            return err
        }
        s.impl.SetProp2(prop)
    default:
        return fmt.Errorf("NestedStruct2InterfaceSource.SetProperty: unknown property %s", propertyId)
    }
    return nil
}

func (s *NestedStruct2InterfaceSource) CollectProperties() (core.KWArgs, error) {
    if s.impl == nil {
        return nil, fmt.Errorf("no implementation")
    }
    return core.KWArgs{
        "prop1": s.impl.GetProp1(),
        "prop2": s.impl.GetProp2(),
    }, nil
}


func (s *NestedStruct2InterfaceSource) Linked(objectId string, node *remote.Node) error {
    if objectId != s.ObjectId() {
        return fmt.Errorf("not my object: %s", objectId)
    }
    s.node = node
    return nil
}

func (s *NestedStruct2InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *NestedStruct2InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}

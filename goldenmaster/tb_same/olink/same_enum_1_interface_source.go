package olink

import (
    "fmt"
    "log"
	"olink/pkg/core"
    "olink/pkg/remote"
    "goldenmaster/tb_same/api"
)

type SameEnum1InterfaceSource struct {
    node *remote.Node
    impl api.SameEnum1Interface
}

var _ remote.IObjectSource = (*SameEnum1InterfaceSource)(nil)
var _ api.INotifier = (*SameEnum1InterfaceSource)(nil)

func NewSameEnum1InterfaceSource() *SameEnum1InterfaceSource {
    return &SameEnum1InterfaceSource{}
}

func (s *SameEnum1InterfaceSource) SetImplementation(impl api.SameEnum1Interface) {
    s.impl = impl
}

func (s *SameEnum1InterfaceSource) ObjectId() string {
	return "tb.same.SameEnum1Interface"
}

func (s *SameEnum1InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
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

func (s *SameEnum1InterfaceSource) SetProperty(propertyId string, value core.Any) error {
    if s.impl == nil {
        return fmt.Errorf("no implementation")
    }
    name := core.ToMember(propertyId)
    switch name {
    case "prop1":
        prop, err := api.AsEnum1(value)
        if err != nil {
            return err
        }
        s.impl.SetProp1(prop)
    default:
        return fmt.Errorf("SameEnum1InterfaceSource.SetProperty: unknown property %s", propertyId)
    }
    return nil
}

func (s *SameEnum1InterfaceSource) CollectProperties() (core.KWArgs, error) {
    if s.impl == nil {
        return nil, fmt.Errorf("no implementation")
    }
    return core.KWArgs{
        "prop1": s.impl.GetProp1(),
    }, nil
}


func (s *SameEnum1InterfaceSource) Linked(objectId string, node *remote.Node) error {
    if objectId != s.ObjectId() {
        return fmt.Errorf("not my object: %s", objectId)
    }
    s.node = node
    return nil
}

func (s *SameEnum1InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *SameEnum1InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}

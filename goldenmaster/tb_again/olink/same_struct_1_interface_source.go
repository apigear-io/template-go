package olink

import (
    "fmt"
    "log"
	"olink/pkg/core"
    "olink/pkg/remote"
    "goldenmaster/tb_again/api"
)

type SameStruct1InterfaceSource struct {
    node *remote.Node
    impl api.SameStruct1Interface
}

var _ remote.IObjectSource = (*SameStruct1InterfaceSource)(nil)
var _ api.INotifier = (*SameStruct1InterfaceSource)(nil)

func NewSameStruct1InterfaceSource() *SameStruct1InterfaceSource {
    return &SameStruct1InterfaceSource{}
}

func (s *SameStruct1InterfaceSource) SetImplementation(impl api.SameStruct1Interface) {
    s.impl = impl
}

func (s *SameStruct1InterfaceSource) ObjectId() string {
	return "tb.again.SameStruct1Interface"
}

func (s *SameStruct1InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
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

func (s *SameStruct1InterfaceSource) SetProperty(propertyId string, value core.Any) error {
    if s.impl == nil {
        return fmt.Errorf("no implementation")
    }
    name := core.ToMember(propertyId)
    switch name {
    case "prop1":
        prop, err := api.AsStruct1(value)
        if err != nil {
            return err
        }
        s.impl.SetProp1(prop)
    default:
        return fmt.Errorf("SameStruct1InterfaceSource.SetProperty: unknown property %s", propertyId)
    }
    return nil
}

func (s *SameStruct1InterfaceSource) CollectProperties() (core.KWArgs, error) {
    if s.impl == nil {
        return nil, fmt.Errorf("no implementation")
    }
    return core.KWArgs{
        "prop1": s.impl.GetProp1(),
    }, nil
}


func (s *SameStruct1InterfaceSource) Linked(objectId string, node *remote.Node) error {
    if objectId != s.ObjectId() {
        return fmt.Errorf("not my object: %s", objectId)
    }
    s.node = node
    return nil
}

func (s *SameStruct1InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *SameStruct1InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}

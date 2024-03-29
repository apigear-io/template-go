package main
import (
	"flag"
	"log"
	"net/http"
    "olink/pkg/ws"
    "olink/pkg/remote"
    tb_adv_olink "goldenmaster/tb_adv/olink"
    tb_adv_impl "goldenmaster/tb_adv/olink/testbed"
    tb_conflict_olink "goldenmaster/tb_conflict/olink"
    tb_conflict_impl "goldenmaster/tb_conflict/olink/testbed"
    tb_data_olink "goldenmaster/tb_data/olink"
    tb_data_impl "goldenmaster/tb_data/olink/testbed"
    tb_enum_olink "goldenmaster/tb_enum/olink"
    tb_enum_impl "goldenmaster/tb_enum/olink/testbed"
    tb_same_olink "goldenmaster/tb_same/olink"
    tb_same_impl "goldenmaster/tb_same/olink/testbed"
    tb_again_olink "goldenmaster/tb_again/olink"
    tb_again_impl "goldenmaster/tb_again/olink/testbed"
    tb_simple_olink "goldenmaster/tb_simple/olink"
    tb_simple_impl "goldenmaster/tb_simple/olink/testbed"    
)


var addr = flag.String("addr", ":8080", "http service address")

var registry = remote.NewRegistry()
var node = remote.NewNode(registry)
var hub = ws.NewHub(registry)

func init() {
    {   // register tb_adv module
        source := tb_adv_olink.NewManyParamInterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_adv_impl.NewManyParamInterface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_adv module
        source := tb_adv_olink.NewNestedStruct1InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_adv_impl.NewNestedStruct1Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_adv module
        source := tb_adv_olink.NewNestedStruct2InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_adv_impl.NewNestedStruct2Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_adv module
        source := tb_adv_olink.NewNestedStruct3InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_adv_impl.NewNestedStruct3Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_conflict module
        source := tb_conflict_olink.NewConflict1Source()
        registry.AddObjectSource(source)
        impl := tb_conflict_impl.NewConflict1(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_conflict module
        source := tb_conflict_olink.NewConflict2Source()
        registry.AddObjectSource(source)
        impl := tb_conflict_impl.NewConflict2(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_conflict module
        source := tb_conflict_olink.NewConflict3Source()
        registry.AddObjectSource(source)
        impl := tb_conflict_impl.NewConflict3(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_conflict module
        source := tb_conflict_olink.NewConflict4Source()
        registry.AddObjectSource(source)
        impl := tb_conflict_impl.NewConflict4(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_data module
        source := tb_data_olink.NewStructInterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_data_impl.NewStructInterface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_data module
        source := tb_data_olink.NewStructArrayInterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_data_impl.NewStructArrayInterface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_enum module
        source := tb_enum_olink.NewEnumInterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_enum_impl.NewEnumInterface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_same module
        source := tb_same_olink.NewSameStruct1InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_same_impl.NewSameStruct1Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_same module
        source := tb_same_olink.NewSameStruct2InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_same_impl.NewSameStruct2Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_same module
        source := tb_same_olink.NewSameEnum1InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_same_impl.NewSameEnum1Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_same module
        source := tb_same_olink.NewSameEnum2InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_same_impl.NewSameEnum2Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_again module
        source := tb_again_olink.NewSameStruct1InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_again_impl.NewSameStruct1Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_again module
        source := tb_again_olink.NewSameStruct2InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_again_impl.NewSameStruct2Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_again module
        source := tb_again_olink.NewSameEnum1InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_again_impl.NewSameEnum1Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_again module
        source := tb_again_olink.NewSameEnum2InterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_again_impl.NewSameEnum2Interface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_simple module
        source := tb_simple_olink.NewSimpleInterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_simple_impl.NewSimpleInterface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
    {   // register tb_simple module
        source := tb_simple_olink.NewSimpleArrayInterfaceSource()
        registry.AddObjectSource(source)
        impl := tb_simple_impl.NewSimpleArrayInterface(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
}

func main() {
	flag.Parse()
    http.HandleFunc("/ws", hub.ServeHTTP)
    log.Printf("listen on %s/ws", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
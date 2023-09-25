package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	tb_enum_olink "testbed/tb_enum/olink"
	tb_enum_impl "testbed/tb_enum/olink/testbed"
	tb_same1_olink "testbed/tb_same1/olink"
	tb_same1_impl "testbed/tb_same1/olink/testbed"
	tb_same2_olink "testbed/tb_same2/olink"
	tb_same2_impl "testbed/tb_same2/olink/testbed"
	tb_simple_olink "testbed/tb_simple/olink"
	tb_simple_impl "testbed/tb_simple/olink/testbed"
	testbed1_olink "testbed/testbed1/olink"
	testbed1_impl "testbed/testbed1/olink/testbed"
	testbed2_olink "testbed/testbed2/olink"
	testbed2_impl "testbed/testbed2/olink/testbed"

	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

var addr = flag.String("addr", ":8080", "http service address")
var registry = remote.NewRegistry()
var node = remote.NewNode(registry)
var ctx = context.Background()
var hub = ws.NewHub(ctx, registry)

func init() {
	{ // register testbed2 module
		source := testbed2_olink.NewManyParamInterfaceSource()
		registry.AddObjectSource(source)
		impl := testbed2_impl.NewManyParamInterface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register testbed2 module
		source := testbed2_olink.NewNestedStruct1InterfaceSource()
		registry.AddObjectSource(source)
		impl := testbed2_impl.NewNestedStruct1Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register testbed2 module
		source := testbed2_olink.NewNestedStruct2InterfaceSource()
		registry.AddObjectSource(source)
		impl := testbed2_impl.NewNestedStruct2Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register testbed2 module
		source := testbed2_olink.NewNestedStruct3InterfaceSource()
		registry.AddObjectSource(source)
		impl := testbed2_impl.NewNestedStruct3Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_enum module
		source := tb_enum_olink.NewEnumInterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_enum_impl.NewEnumInterface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same1 module
		source := tb_same1_olink.NewSameStruct1InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same1_impl.NewSameStruct1Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same1 module
		source := tb_same1_olink.NewSameStruct2InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same1_impl.NewSameStruct2Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same1 module
		source := tb_same1_olink.NewSameEnum1InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same1_impl.NewSameEnum1Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same1 module
		source := tb_same1_olink.NewSameEnum2InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same1_impl.NewSameEnum2Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same2 module
		source := tb_same2_olink.NewSameStruct1InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same2_impl.NewSameStruct1Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same2 module
		source := tb_same2_olink.NewSameStruct2InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same2_impl.NewSameStruct2Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same2 module
		source := tb_same2_olink.NewSameEnum1InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same2_impl.NewSameEnum1Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_same2 module
		source := tb_same2_olink.NewSameEnum2InterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_same2_impl.NewSameEnum2Interface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_simple module
		source := tb_simple_olink.NewSimpleInterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_simple_impl.NewSimpleInterface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register tb_simple module
		source := tb_simple_olink.NewSimpleArrayInterfaceSource()
		registry.AddObjectSource(source)
		impl := tb_simple_impl.NewSimpleArrayInterface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register testbed1 module
		source := testbed1_olink.NewStructInterfaceSource()
		registry.AddObjectSource(source)
		impl := testbed1_impl.NewStructInterface(source)
		source.SetImplementation(impl)
		registry.LinkRemoteNode(source.ObjectId(), node)
	}
	{ // register testbed1 module
		source := testbed1_olink.NewStructArrayInterfaceSource()
		registry.AddObjectSource(source)
		impl := testbed1_impl.NewStructArrayInterface(source)
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

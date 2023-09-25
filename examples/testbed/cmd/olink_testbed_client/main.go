package main

import (
	"context"
	"flag"
	"log"
	tb_enum_olink "testbed/tb_enum/olink"
	tb_same1_olink "testbed/tb_same1/olink"
	tb_same2_olink "testbed/tb_same2/olink"
	tb_simple_olink "testbed/tb_simple/olink"
	testbed1_olink "testbed/testbed1/olink"
	testbed2_olink "testbed/testbed2/olink"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

var registry = client.NewRegistry()
var node = client.NewNode(registry)

func registerSinks() {
	{ // register testbed2 module
		sink := testbed2_olink.NewManyParamInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register testbed2 module
		sink := testbed2_olink.NewNestedStruct1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register testbed2 module
		sink := testbed2_olink.NewNestedStruct2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register testbed2 module
		sink := testbed2_olink.NewNestedStruct3InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_enum module
		sink := tb_enum_olink.NewEnumInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same1 module
		sink := tb_same1_olink.NewSameStruct1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same1 module
		sink := tb_same1_olink.NewSameStruct2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same1 module
		sink := tb_same1_olink.NewSameEnum1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same1 module
		sink := tb_same1_olink.NewSameEnum2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same2 module
		sink := tb_same2_olink.NewSameStruct1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same2 module
		sink := tb_same2_olink.NewSameStruct2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same2 module
		sink := tb_same2_olink.NewSameEnum1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same2 module
		sink := tb_same2_olink.NewSameEnum2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_simple module
		sink := tb_simple_olink.NewSimpleInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_simple module
		sink := tb_simple_olink.NewSimpleArrayInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register testbed1 module
		sink := testbed1_olink.NewStructInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register testbed1 module
		sink := testbed1_olink.NewStructArrayInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
}

func main() {
	var addr = flag.String("addr", "ws://127.0.0.1:8080/ws", "ws service addr")
	flag.Parse()
	ctx := context.Background()
	conn, err := ws.Dial(ctx, *addr)
	if err != nil {
		log.Fatalf("dial error: %s\n", err)
		return
	}
	defer conn.Close()
	node.SetOutput(conn)
	conn.SetOutput(node)
	registry.AttachClientNode(node)
	registerSinks()

}

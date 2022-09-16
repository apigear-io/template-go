package main
import (
    "flag"
	"log"
	"olink/pkg/client"
	"olink/pkg/ws"
    tb_adv_olink "goldenmaster/tb_adv/olink"
    tb_conflict_olink "goldenmaster/tb_conflict/olink"
    tb_data_olink "goldenmaster/tb_data/olink"
    tb_enum_olink "goldenmaster/tb_enum/olink"
    tb_same_olink "goldenmaster/tb_same/olink"
    tb_again_olink "goldenmaster/tb_again/olink"
    tb_simple_olink "goldenmaster/tb_simple/olink"    
)

var addr = flag.String("addr", "ws://127.0.0.1:8080/ws", "ws service addr")


var registry = client.NewRegistry()
var node = client.NewNode(registry)


func registerSinks() {
	{ // register tb_adv module
		sink := tb_adv_olink.NewManyParamInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_adv module
		sink := tb_adv_olink.NewNestedStruct1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_adv module
		sink := tb_adv_olink.NewNestedStruct2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_adv module
		sink := tb_adv_olink.NewNestedStruct3InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_conflict module
		sink := tb_conflict_olink.NewConflict1Sink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_conflict module
		sink := tb_conflict_olink.NewConflict2Sink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_conflict module
		sink := tb_conflict_olink.NewConflict3Sink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_conflict module
		sink := tb_conflict_olink.NewConflict4Sink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_data module
		sink := tb_data_olink.NewStructInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_data module
		sink := tb_data_olink.NewStructArrayInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_enum module
		sink := tb_enum_olink.NewEnumInterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same module
		sink := tb_same_olink.NewSameStruct1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same module
		sink := tb_same_olink.NewSameStruct2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same module
		sink := tb_same_olink.NewSameEnum1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_same module
		sink := tb_same_olink.NewSameEnum2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_again module
		sink := tb_again_olink.NewSameStruct1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_again module
		sink := tb_again_olink.NewSameStruct2InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_again module
		sink := tb_again_olink.NewSameEnum1InterfaceSink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
	{ // register tb_again module
		sink := tb_again_olink.NewSameEnum2InterfaceSink(node)
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
}

func main() {
    flag.Parse()
    conn, err := ws.Dial(*addr)
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
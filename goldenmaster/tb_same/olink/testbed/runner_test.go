package testbed

import (
    "fmt"
    "olink/pkg/client"
    "olink/pkg/ws"
    "goldenmaster/tb_same/api"
    "goldenmaster/tb_same/olink"
    "testing"
    
    "github.com/stretchr/testify/assert"
)

func getConnection(t *testing.T) (*client.Node, func(), error) {
	var addr = "ws://127.0.0.1:8080/ws"

	var registry = client.NewRegistry()
	var node = client.NewNode(registry)


	conn, err := ws.Dial(addr)
	if err != nil {
		return nil, nil, fmt.Errorf("dial error: %s\n", err)
	}
    done := func() {
        conn.Close()
    }
	node.SetOutput(conn)
	conn.SetOutput(node)
	registry.AttachClientNode(node)
    return node, done, nil
}

func TestSameStruct1Interface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewSameStruct1InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}

func TestSameStruct2Interface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewSameStruct2InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}

func TestSameEnum1Interface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewSameEnum1InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}

func TestSameEnum2Interface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewSameEnum2InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}
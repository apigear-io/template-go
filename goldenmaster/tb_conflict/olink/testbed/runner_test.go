package testbed

import (
	"fmt"
	"goldenmaster/tb_conflict/api"
	"goldenmaster/tb_conflict/olink"
	"olink/pkg/client"
	"olink/pkg/ws"
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

func TestConflict1(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewConflict1Sink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}

func TestConflict2(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewConflict2Sink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}

func TestConflict3(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewConflict3Sink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}

func TestConflict4(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewConflict4Sink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
}

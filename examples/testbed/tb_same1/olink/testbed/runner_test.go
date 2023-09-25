package testbed

import (
	"context"
	"fmt"
	"testbed/tb_same1/api"
	"testbed/tb_same1/olink"
	"testing"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/ws"

	"github.com/stretchr/testify/assert"
)

func getConnection(t *testing.T) (*client.Node, func(), error) {
	var addr = "ws://127.0.0.1:8080/ws"
	var ctx = context.Background()

	var registry = client.NewRegistry()
	var node = client.NewNode(registry)

	conn, err := ws.Dial(ctx, addr)
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
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := api.Struct1{}
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
}

func TestSameStruct2Interface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewSameStruct2InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := api.Struct1{}
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
	t.Run("func2", func(t *testing.T) {
		param1 := api.Struct1{}
		param2 := api.Struct2{}
		resultFunc2 := sink.Func2(param1, param2)
		assert.Equal(t, param1, resultFunc2)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
	})
}

func TestSameEnum1Interface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewSameEnum1InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := api.Enum1Value1
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
}

func TestSameEnum2Interface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewSameEnum2InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := api.Enum1Value1
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
	t.Run("func2", func(t *testing.T) {
		param1 := api.Enum1Value1
		param2 := api.Enum2Value1
		resultFunc2 := sink.Func2(param1, param2)
		assert.Equal(t, param1, resultFunc2)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
	})
}

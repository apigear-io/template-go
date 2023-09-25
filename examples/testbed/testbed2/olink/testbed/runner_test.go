package testbed

import (
	"context"
	"fmt"
	"testbed/testbed2/api"
	"testbed/testbed2/olink"
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

func TestManyParamInterface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewManyParamInterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := int32(0)
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
	t.Run("func2", func(t *testing.T) {
		param1 := int32(0)
		param2 := int32(0)
		resultFunc2 := sink.Func2(param1, param2)
		assert.Equal(t, param1, resultFunc2)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
	})
	t.Run("func3", func(t *testing.T) {
		param1 := int32(0)
		param2 := int32(0)
		param3 := int32(0)
		resultFunc3 := sink.Func3(param1, param2, param3)
		assert.Equal(t, param1, resultFunc3)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
		assert.Equal(t, param3, sink.Prop3)
	})
	t.Run("func4", func(t *testing.T) {
		param1 := int32(0)
		param2 := int32(0)
		param3 := int32(0)
		param4 := int32(0)
		resultFunc4 := sink.Func4(param1, param2, param3, param4)
		assert.Equal(t, param1, resultFunc4)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
		assert.Equal(t, param3, sink.Prop3)
		assert.Equal(t, param4, sink.Prop4)
	})
}

func TestNestedStruct1Interface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewNestedStruct1InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := api.NestedStruct1{}
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
}

func TestNestedStruct2Interface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewNestedStruct2InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := api.NestedStruct1{}
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
	t.Run("func2", func(t *testing.T) {
		param1 := api.NestedStruct1{}
		param2 := api.NestedStruct2{}
		resultFunc2 := sink.Func2(param1, param2)
		assert.Equal(t, param1, resultFunc2)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
	})
}

func TestNestedStruct3Interface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewNestedStruct3InterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry().AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func1", func(t *testing.T) {
		param1 := api.NestedStruct1{}
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
	t.Run("func2", func(t *testing.T) {
		param1 := api.NestedStruct1{}
		param2 := api.NestedStruct2{}
		resultFunc2 := sink.Func2(param1, param2)
		assert.Equal(t, param1, resultFunc2)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
	})
	t.Run("func3", func(t *testing.T) {
		param1 := api.NestedStruct1{}
		param2 := api.NestedStruct2{}
		param3 := api.NestedStruct3{}
		resultFunc3 := sink.Func3(param1, param2, param3)
		assert.Equal(t, param1, resultFunc3)
		assert.Equal(t, param1, sink.Prop1)
		assert.Equal(t, param2, sink.Prop2)
		assert.Equal(t, param3, sink.Prop3)
	})
}

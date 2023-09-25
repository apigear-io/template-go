package testbed

import (
	"fmt"
	"goldenmaster/tb_enum/api"
	"goldenmaster/tb_enum/olink"
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

func TestEnumInterface(t *testing.T) {
	node, done, err := getConnection(t)
	defer done()
	assert.NoError(t, err)
	sink := olink.NewEnumInterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
	t.Run("func0", func(t *testing.T) {
		param0 := api.Enum0Value0
		resultFunc0 := sink.Func0(param0)
		assert.Equal(t, param0, resultFunc0)
		assert.Equal(t, param0, sink.Prop0)
	})
	t.Run("func1", func(t *testing.T) {
		param1 := api.Enum1Value1
		resultFunc1 := sink.Func1(param1)
		assert.Equal(t, param1, resultFunc1)
		assert.Equal(t, param1, sink.Prop1)
	})
	t.Run("func2", func(t *testing.T) {
		param2 := api.Enum2Value2
		resultFunc2 := sink.Func2(param2)
		assert.Equal(t, param2, resultFunc2)
		assert.Equal(t, param2, sink.Prop2)
	})
	t.Run("func3", func(t *testing.T) {
		param3 := api.Enum3Value3
		resultFunc3 := sink.Func3(param3)
		assert.Equal(t, param3, resultFunc3)
		assert.Equal(t, param3, sink.Prop3)
	})
}

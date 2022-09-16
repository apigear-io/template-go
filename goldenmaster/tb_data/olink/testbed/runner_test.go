package testbed

import (
    "fmt"
    "olink/pkg/client"
    "olink/pkg/ws"
    "goldenmaster/tb_data/api"
    "goldenmaster/tb_data/olink"
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

func TestStructInterface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewStructInterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
    t.Run("funcBool", func(t *testing.T) {
        paramBool := api.StructBool{}
        resultFuncBool := sink.FuncBool(paramBool)
        assert.Equal(t, paramBool, resultFuncBool)
        assert.Equal(t, paramBool, sink.PropBool)
    })
    t.Run("funcInt", func(t *testing.T) {
        paramInt := api.StructInt{}
        resultFuncInt := sink.FuncInt(paramInt)
        assert.Equal(t, paramInt, resultFuncInt)
        assert.Equal(t, paramInt, sink.PropInt)
    })
    t.Run("funcFloat", func(t *testing.T) {
        paramFloat := api.StructFloat{}
        resultFuncFloat := sink.FuncFloat(paramFloat)
        assert.Equal(t, paramFloat, resultFuncFloat)
        assert.Equal(t, paramFloat, sink.PropFloat)
    })
    t.Run("funcString", func(t *testing.T) {
        paramString := api.StructString{}
        resultFuncString := sink.FuncString(paramString)
        assert.Equal(t, paramString, resultFuncString)
        assert.Equal(t, paramString, sink.PropString)
    })
}

func TestStructArrayInterface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewStructArrayInterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
    t.Run("funcBool", func(t *testing.T) {
        paramBool := []api.StructBool{}
        resultFuncBool := sink.FuncBool(paramBool)
        assert.Equal(t, paramBool, resultFuncBool)
        assert.Equal(t, paramBool, sink.PropBool)
    })
    t.Run("funcInt", func(t *testing.T) {
        paramInt := []api.StructInt{}
        resultFuncInt := sink.FuncInt(paramInt)
        assert.Equal(t, paramInt, resultFuncInt)
        assert.Equal(t, paramInt, sink.PropInt)
    })
    t.Run("funcFloat", func(t *testing.T) {
        paramFloat := []api.StructFloat{}
        resultFuncFloat := sink.FuncFloat(paramFloat)
        assert.Equal(t, paramFloat, resultFuncFloat)
        assert.Equal(t, paramFloat, sink.PropFloat)
    })
    t.Run("funcString", func(t *testing.T) {
        paramString := []api.StructString{}
        resultFuncString := sink.FuncString(paramString)
        assert.Equal(t, paramString, resultFuncString)
        assert.Equal(t, paramString, sink.PropString)
    })
}
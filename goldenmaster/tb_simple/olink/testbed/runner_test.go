package testbed

import (
    "fmt"
    "olink/pkg/client"
    "olink/pkg/ws"
    "goldenmaster/tb_simple/api"
    "goldenmaster/tb_simple/olink"
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

func TestSimpleInterface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewSimpleInterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
    t.Run("funcBool", func(t *testing.T) {
        paramBool := false
        resultFuncBool := sink.FuncBool(paramBool)
        assert.Equal(t, paramBool, resultFuncBool)
        assert.Equal(t, paramBool, sink.PropBool)
    })
    t.Run("funcInt", func(t *testing.T) {
        paramInt := int64(0)
        resultFuncInt := sink.FuncInt(paramInt)
        assert.Equal(t, paramInt, resultFuncInt)
        assert.Equal(t, paramInt, sink.PropInt)
    })
    t.Run("funcFloat", func(t *testing.T) {
        paramFloat := float64(0.0)
        resultFuncFloat := sink.FuncFloat(paramFloat)
        assert.Equal(t, paramFloat, resultFuncFloat)
        assert.Equal(t, paramFloat, sink.PropFloat)
    })
    t.Run("funcString", func(t *testing.T) {
        paramString := ""
        resultFuncString := sink.FuncString(paramString)
        assert.Equal(t, paramString, resultFuncString)
        assert.Equal(t, paramString, sink.PropString)
    })
}

func TestSimpleArrayInterface(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.NewSimpleArrayInterfaceSink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
    t.Run("funcBool", func(t *testing.T) {
        paramBool := []bool{}
        resultFuncBool := sink.FuncBool(paramBool)
        assert.Equal(t, paramBool, resultFuncBool)
        assert.Equal(t, paramBool, sink.PropBool)
    })
    t.Run("funcInt", func(t *testing.T) {
        paramInt := []int64{}
        resultFuncInt := sink.FuncInt(paramInt)
        assert.Equal(t, paramInt, resultFuncInt)
        assert.Equal(t, paramInt, sink.PropInt)
    })
    t.Run("funcFloat", func(t *testing.T) {
        paramFloat := []float64{}
        resultFuncFloat := sink.FuncFloat(paramFloat)
        assert.Equal(t, paramFloat, resultFuncFloat)
        assert.Equal(t, paramFloat, sink.PropFloat)
    })
    t.Run("funcString", func(t *testing.T) {
        paramString := []string{}
        resultFuncString := sink.FuncString(paramString)
        assert.Equal(t, paramString, resultFuncString)
        assert.Equal(t, paramString, sink.PropString)
    })
}
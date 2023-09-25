# Olink Testbed Client and Server

ApiGear comes with a Testbed API which validates most of ApiGears template facilities. The API is designed so that types and replies can be automatically generated from the API description. The API is defined in the apigear folder in each template.

For example

```yaml
schema: apigear.module/1.0
name: tb.simple
version: "1.0"

interfaces:
  - name: SimpleInterface
    properties:
      - { name: propBool, type: bool }
      - { name: propInt, type: int }
    operations:
      - name: funcBool
        params:
          - { name: paramBool, type: bool }
        return:
          type: bool
      - name: funcFloat
        params:
          - { name: paramFloat, type: float }
        return:
          type: float
    signals:
      - name: sigBool
        params:
          - { name: paramBool, type: bool }
      - name: sigInt
        params:
          - { name: paramInt, type: int }
```

You can see that a propBool is of type bool and funcBool receives as first param a paramBool and returns a bool. A signal sigBool has at first param a paramBool.

So we can design a test run where we call each operation (e.g. funcBool, funcFloat, ...) with a set of parameters and check the reply. We can also emit signals and check if the signal was received and trigger a property change and check if the property change was received.

call function with parameters and check reply
the function body is generated from the template and

- trigger property change and check if property change was received
- emit signals and check if signal was received

From the go code this should look like this:

```go
func (c *SimpleInterfaceClient) FuncBool(paramBool bool) (bool, error) {
    var reply bool
    err := c.Call("funcBool", paramBool, &reply)
    return reply, err
}
```

And the server side should look like this:

```go
func (s *SimpleInterfaceServer) OnFuncBool(paramBool bool) (bool, error) {
    s.SetPropBool(paramBool)
    s.NotifySigBool(paramBool)
    return paramBool, err
}
```

```go

```

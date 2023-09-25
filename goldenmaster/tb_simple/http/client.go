package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(c *http.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) Post(url string, data interface{}) (interface{}, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) SimpleInterfaceFuncBool(req *SimpleInterfaceFuncBoolRequest, reply *SimpleInterfaceFuncBoolReply) error {
	url := "tb.simple/SimpleInterface/funcBool"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleInterfaceFuncBoolReply)
	return nil
}

func (c *Client) SimpleInterfaceFuncInt(req *SimpleInterfaceFuncIntRequest, reply *SimpleInterfaceFuncIntReply) error {
	url := "tb.simple/SimpleInterface/funcInt"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleInterfaceFuncIntReply)
	return nil
}

func (c *Client) SimpleInterfaceFuncFloat(req *SimpleInterfaceFuncFloatRequest, reply *SimpleInterfaceFuncFloatReply) error {
	url := "tb.simple/SimpleInterface/funcFloat"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleInterfaceFuncFloatReply)
	return nil
}

func (c *Client) SimpleInterfaceFuncString(req *SimpleInterfaceFuncStringRequest, reply *SimpleInterfaceFuncStringReply) error {
	url := "tb.simple/SimpleInterface/funcString"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleInterfaceFuncStringReply)
	return nil
}

func (c *Client) SimpleArrayInterfaceFuncBool(req *SimpleArrayInterfaceFuncBoolRequest, reply *SimpleArrayInterfaceFuncBoolReply) error {
	url := "tb.simple/SimpleArrayInterface/funcBool"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleArrayInterfaceFuncBoolReply)
	return nil
}

func (c *Client) SimpleArrayInterfaceFuncInt(req *SimpleArrayInterfaceFuncIntRequest, reply *SimpleArrayInterfaceFuncIntReply) error {
	url := "tb.simple/SimpleArrayInterface/funcInt"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleArrayInterfaceFuncIntReply)
	return nil
}

func (c *Client) SimpleArrayInterfaceFuncFloat(req *SimpleArrayInterfaceFuncFloatRequest, reply *SimpleArrayInterfaceFuncFloatReply) error {
	url := "tb.simple/SimpleArrayInterface/funcFloat"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleArrayInterfaceFuncFloatReply)
	return nil
}

func (c *Client) SimpleArrayInterfaceFuncString(req *SimpleArrayInterfaceFuncStringRequest, reply *SimpleArrayInterfaceFuncStringReply) error {
	url := "tb.simple/SimpleArrayInterface/funcString"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*SimpleArrayInterfaceFuncStringReply)
	return nil
}

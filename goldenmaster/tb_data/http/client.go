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

func (c *Client) StructInterfaceFuncBool(req *StructInterfaceFuncBoolRequest, reply *StructInterfaceFuncBoolReply) error {
	url := "tb.data/StructInterface/funcBool"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructInterfaceFuncBoolReply)
	return nil
}

func (c *Client) StructInterfaceFuncInt(req *StructInterfaceFuncIntRequest, reply *StructInterfaceFuncIntReply) error {
	url := "tb.data/StructInterface/funcInt"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructInterfaceFuncIntReply)
	return nil
}

func (c *Client) StructInterfaceFuncFloat(req *StructInterfaceFuncFloatRequest, reply *StructInterfaceFuncFloatReply) error {
	url := "tb.data/StructInterface/funcFloat"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructInterfaceFuncFloatReply)
	return nil
}

func (c *Client) StructInterfaceFuncString(req *StructInterfaceFuncStringRequest, reply *StructInterfaceFuncStringReply) error {
	url := "tb.data/StructInterface/funcString"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructInterfaceFuncStringReply)
	return nil
}

func (c *Client) StructArrayInterfaceFuncBool(req *StructArrayInterfaceFuncBoolRequest, reply *StructArrayInterfaceFuncBoolReply) error {
	url := "tb.data/StructArrayInterface/funcBool"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructArrayInterfaceFuncBoolReply)
	return nil
}

func (c *Client) StructArrayInterfaceFuncInt(req *StructArrayInterfaceFuncIntRequest, reply *StructArrayInterfaceFuncIntReply) error {
	url := "tb.data/StructArrayInterface/funcInt"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructArrayInterfaceFuncIntReply)
	return nil
}

func (c *Client) StructArrayInterfaceFuncFloat(req *StructArrayInterfaceFuncFloatRequest, reply *StructArrayInterfaceFuncFloatReply) error {
	url := "tb.data/StructArrayInterface/funcFloat"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructArrayInterfaceFuncFloatReply)
	return nil
}

func (c *Client) StructArrayInterfaceFuncString(req *StructArrayInterfaceFuncStringRequest, reply *StructArrayInterfaceFuncStringReply) error {
	url := "tb.data/StructArrayInterface/funcString"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*StructArrayInterfaceFuncStringReply)
	return nil
}

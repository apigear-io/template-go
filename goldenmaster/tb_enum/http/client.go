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

func (c *Client) EnumInterfaceFunc0(req *EnumInterfaceFunc0Request, reply *EnumInterfaceFunc0Reply) error {
	url := "tb.enum/EnumInterface/func0"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*EnumInterfaceFunc0Reply)
	return nil
}

func (c *Client) EnumInterfaceFunc1(req *EnumInterfaceFunc1Request, reply *EnumInterfaceFunc1Reply) error {
	url := "tb.enum/EnumInterface/func1"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*EnumInterfaceFunc1Reply)
	return nil
}

func (c *Client) EnumInterfaceFunc2(req *EnumInterfaceFunc2Request, reply *EnumInterfaceFunc2Reply) error {
	url := "tb.enum/EnumInterface/func2"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*EnumInterfaceFunc2Reply)
	return nil
}

func (c *Client) EnumInterfaceFunc3(req *EnumInterfaceFunc3Request, reply *EnumInterfaceFunc3Reply) error {
	url := "tb.enum/EnumInterface/func3"
	resp, err := c.Post(url, req)
	if err != nil {
		return err
	}
	*reply = *resp.(*EnumInterfaceFunc3Reply)
	return nil
}

package client

import (
	"errors"
	"fmt"
)

// Config contains the configuration for the client
type Config struct {
	APIKey string
}

// Client is a struct representing the client to interact with the API
type Client struct {
	config *Config
	data   map[string]interface{}
}

// NewClient returns a new API client
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		data:   make(map[string]interface{}),
	}
}

// FooResource represents a Foo resource
type FooResource struct {
	ID          string
	Name        string
	Description string
}

// BarResource represents a Bar resource
type BarResource struct {
	ID      string
	Title   string
	Content string
}

// Error to represent a not found resource
var ErrNotFound = errors.New("resource not found")

// IsNotFound checks if the error is a not found error
func IsNotFound(err error) bool {
	return err == ErrNotFound
}

// CreateFoo simulates creating a Foo resource
func (c *Client) CreateFoo(name, description string) (*FooResource, error) {
	id := fmt.Sprintf("foo-%s", name)
	foo := &FooResource{
		ID:          id,
		Name:        name,
		Description: description,
	}
	c.data[id] = foo
	return foo, nil
}

// GetFoo simulates retrieving a Foo resource
func (c *Client) GetFoo(id string) (*FooResource, error) {
	if resource, ok := c.data[id]; ok {
		return resource.(*FooResource), nil
	}
	return nil, ErrNotFound
}

// UpdateFoo simulates updating a Foo resource
func (c *Client) UpdateFoo(id, name, description string) (*FooResource, error) {
	if resource, ok := c.data[id]; ok {
		foo := resource.(*FooResource)
		foo.Name = name
		foo.Description = description
		c.data[id] = foo
		return foo, nil
	}
	return nil, ErrNotFound
}

// DeleteFoo simulates deleting a Foo resource
func (c *Client) DeleteFoo(id string) error {
	if _, ok := c.data[id]; ok {
		delete(c.data, id)
		return nil
	}
	return ErrNotFound
}

// CreateBar simulates creating a Bar resource
func (c *Client) CreateBar(title, content string) (*BarResource, error) {
	id := fmt.Sprintf("bar-%s", title)
	bar := &BarResource{
		ID:      id,
		Title:   title,
		Content: content,
	}
	c.data[id] = bar
	return bar, nil
}

// GetBar simulates retrieving a Bar resource
func (c *Client) GetBar(id string) (*BarResource, error) {
	if resource, ok := c.data[id]; ok {
		return resource.(*BarResource), nil
	}
	return nil, ErrNotFound
}

// UpdateBar simulates updating a Bar resource
func (c *Client) UpdateBar(id, title, content string) (*BarResource, error) {
	if resource, ok := c.data[id]; ok {
		bar := resource.(*BarResource)
		bar.Title = title
		bar.Content = content
		c.data[id] = bar
		return bar, nil
	}
	return nil, ErrNotFound
}

// DeleteBar simulates deleting a Bar resource
func (c *Client) DeleteBar(id string) error {
	if _, ok := c.data[id]; ok {
		delete(c.data, id)
		return nil
	}
	return ErrNotFound
}

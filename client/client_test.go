package client_test

import (
	"github.com/gmllt/terraform-provider-template/client"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFoo(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	foo, err := c.CreateFoo("test_foo", "This is a test Foo resource")
	assert.NoError(t, err)
	assert.NotNil(t, foo)
	assert.Equal(t, "test_foo", foo.Name)
	assert.Equal(t, "This is a test Foo resource", foo.Description)
}

func TestGetFoo(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	// First, create a Foo resource
	createdFoo, _ := c.CreateFoo("test_foo", "This is a test Foo resource")

	// Now, retrieve the Foo resource
	foo, err := c.GetFoo(createdFoo.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foo)
	assert.Equal(t, createdFoo.ID, foo.ID)
}

func TestUpdateFoo(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	// First, create a Foo resource
	createdFoo, _ := c.CreateFoo("test_foo", "This is a test Foo resource")

	// Now, update the Foo resource
	updatedFoo, err := c.UpdateFoo(createdFoo.ID, "updated_foo", "Updated description")
	assert.NoError(t, err)
	assert.NotNil(t, updatedFoo)
	assert.Equal(t, "updated_foo", updatedFoo.Name)
	assert.Equal(t, "Updated description", updatedFoo.Description)
}

func TestDeleteFoo(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	// First, create a Foo resource
	createdFoo, _ := c.CreateFoo("test_foo", "This is a test Foo resource")

	// Now, delete the Foo resource
	err := c.DeleteFoo(createdFoo.ID)
	assert.NoError(t, err)

	// Try to retrieve the deleted Foo resource
	foo, err := c.GetFoo(createdFoo.ID)
	assert.Error(t, err)
	assert.Nil(t, foo)
	assert.True(t, client.IsNotFound(err))
}

func TestCreateBar(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	bar, err := c.CreateBar("test_bar", "This is a test Bar resource")
	assert.NoError(t, err)
	assert.NotNil(t, bar)
	assert.Equal(t, "test_bar", bar.Title)
	assert.Equal(t, "This is a test Bar resource", bar.Content)
}

func TestGetBar(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	// First, create a Bar resource
	createdBar, _ := c.CreateBar("test_bar", "This is a test Bar resource")

	// Now, retrieve the Bar resource
	bar, err := c.GetBar(createdBar.ID)
	assert.NoError(t, err)
	assert.NotNil(t, bar)
	assert.Equal(t, createdBar.ID, bar.ID)
}

func TestUpdateBar(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	// First, create a Bar resource
	createdBar, _ := c.CreateBar("test_bar", "This is a test Bar resource")

	// Now, update the Bar resource
	updatedBar, err := c.UpdateBar(createdBar.ID, "updated_bar", "Updated content")
	assert.NoError(t, err)
	assert.NotNil(t, updatedBar)
	assert.Equal(t, "updated_bar", updatedBar.Title)
	assert.Equal(t, "Updated content", updatedBar.Content)
}

func TestDeleteBar(t *testing.T) {
	c := client.NewClient(&client.Config{APIKey: "test_api_key"})

	// First, create a Bar resource
	createdBar, _ := c.CreateBar("test_bar", "This is a test Bar resource")

	// Now, delete the Bar resource
	err := c.DeleteBar(createdBar.ID)
	assert.NoError(t, err)

	// Try to retrieve the deleted Bar resource
	bar, err := c.GetBar(createdBar.ID)
	assert.Error(t, err)
	assert.Nil(t, bar)
	assert.True(t, client.IsNotFound(err))
}

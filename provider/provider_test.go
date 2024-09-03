package provider_test

import (
	"testing"

	"github.com/gmllt/terraform-provider-template/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	p := provider.Provider()

	if err := p.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_configure(t *testing.T) {
	t.Parallel()

	p := provider.Provider()

	resourceData := schema.TestResourceDataRaw(t, p.Schema, map[string]interface{}{
		"api_key": "test_api_key",
	})
	client, err := p.ConfigureFunc(resourceData)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

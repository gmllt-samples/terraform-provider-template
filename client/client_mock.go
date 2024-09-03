package client

// MockClient is a mock implementation of the Client for testing purposes
type MockClient struct {
	Client
}

// NewMockClient creates a new mock client
func NewMockClient() *MockClient {
	return &MockClient{
		Client{
			data: make(map[string]interface{}),
		},
	}
}

// Additional mock methods can be implemented here if necessary

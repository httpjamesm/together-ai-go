package togetherai

import "github.com/go-resty/resty/v2"

type Client struct {
	APIKey string
}

// NewClient creates a new instance of Client.
//
// It takes an apiKey string as a parameter and returns a pointer to a Client.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

func (c *Client) getRestyClient() *resty.Client {
	return resty.New().SetAuthScheme("Bearer").SetAuthToken(c.APIKey).SetBaseURL("https://api.together.xyz")
}

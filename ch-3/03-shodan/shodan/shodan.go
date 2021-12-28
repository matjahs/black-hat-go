package shodan

const BaseURL = "https://api.03-shodan.io"

type Client struct {
  apiKey string
}

func New(apiKey string) *Client {
  return &Client{apiKey}
}

package telegram

import (
	"net/http"
	"net/url"
	"path"
	"strconv"
	"telegram-bot/lib/e"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Update(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	// do request <- getUpdates
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() {err = e.WrapIfErr("can't do request", err)}()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {_=resp.Body.Close()}()
}

func (c *Client) SendMessage() {

}

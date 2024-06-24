package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host       string
	basePath   string
	clientHttp http.Client
}

func New(host, token string) Client {
	return Client{
		host:       host,
		basePath:   basePathGen(token),
		clientHttp: http.Client{},
	}
}

func basePathGen(token string) string {
	return fmt.Sprintf("bot%s", token)
}

func (c *Client) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessages(chatID int, text string) error {
	q := url.Values{}
	q.Add("chatID", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest("sendMessage", q)
	if err != nil {
		return err
	}
	return nil

}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.clientHttp.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

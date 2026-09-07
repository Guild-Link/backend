package compatlink

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		http:    http.Client{Timeout: 15 * time.Second},
	}
}

func (c *Client) Networth(ctx context.Context, profile, museum json.RawMessage, uuid string) (*NetworthResponse, error) {
	response := &NetworthResponse{}
	err := c.post(ctx, "/networth", networthRequest{
		Profile: profile,
		Museum:  museum,
		UUID:    uuid,
	}, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) FarmingWeight(ctx context.Context, profile json.RawMessage, uuid string) (*FarmingWeightResponse, error) {
	response := &FarmingWeightResponse{}
	err := c.post(ctx, "/farming-weight", farmingWeightRequest{
		Profile: profile,
		UUID:    uuid,
	}, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) post(ctx context.Context, path string, data, response any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encode compatlink request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create compatlink request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("send compatlink request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read compatlink response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("compatlink POST %s returned %s: %s", path, resp.Status, respBody)
	}

	if err := json.Unmarshal(respBody, response); err != nil {
		return fmt.Errorf("decode compatlink response: %w", err)
	}

	return nil
}

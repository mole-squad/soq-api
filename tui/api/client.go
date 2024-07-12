package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/burkel24/task-app/pkg/auth"
	"github.com/burkel24/task-app/pkg/focusareas"
	"github.com/burkel24/task-app/pkg/tasks"
)

type Client struct {
	apiHost    string
	httpClient *http.Client

	token string
}

func NewClient() *Client {
	apiHost := os.Getenv("API_HOST")

	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		apiHost:    apiHost,
		httpClient: c,
	}
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) Login(ctx context.Context, username, password string) (string, error) {
	var tokenResp auth.TokenResponseDTO

	reqUrl := url.URL{
		Scheme: "http",
		Host:   c.apiHost,
		Path:   "/auth/token",
	}

	req := auth.NewLoginRequestDTO(username, password)

	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshalling login request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error building login request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("error executing login request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading login response: %w", err)
	}

	if err = json.Unmarshal(respBody, &tokenResp); err != nil {
		return "", fmt.Errorf("error unmarshalling login response: %w", err)
	}

	return tokenResp.Token, nil
}

func (c *Client) ListTasks(ctx context.Context) ([]tasks.TaskDTO, error) {
	reqUrl := url.URL{
		Scheme: "http",
		Host:   c.apiHost,
		Path:   "/tasks",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error building list tasks request: %w", err)
	}

	req.Header = c.buildHeaders()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing list tasks request: %w", err)
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading list tasks response: %w", err)
	}

	var tasksResp []tasks.TaskDTO
	if err = json.Unmarshal(respBody, &tasksResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling list tasks response: %w", err)
	}

	return tasksResp, nil
}

func (c *Client) CreateTask(ctx context.Context, t *tasks.CreateTaskRequestDto) (tasks.TaskDTO, error) {
	var task tasks.TaskDTO

	reqUrl := url.URL{
		Scheme: "http",
		Host:   c.apiHost,
		Path:   "/tasks",
	}

	body, err := json.Marshal(t)
	if err != nil {
		return task, fmt.Errorf("error marshalling create task request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		return task, fmt.Errorf("error building create task request: %w", err)
	}

	req.Header = c.buildHeaders()

	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return task, fmt.Errorf("error executing create task request: %w", err)
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return task, fmt.Errorf("error reading create task response: %w", err)
	}

	if err = json.Unmarshal(respBody, &task); err != nil {
		return task, fmt.Errorf("error unmarshalling create task response: %w", err)
	}

	return task, nil
}

func (c *Client) UpdateTask(ctx context.Context, taskID uint, t *tasks.UpdateTaskRequestDto) (tasks.TaskDTO, error) {
	var task tasks.TaskDTO

	reqUrl := url.URL{
		Scheme: "http",
		Host:   c.apiHost,
		Path:   fmt.Sprintf("/tasks/%d", taskID),
	}

	body, err := json.Marshal(t)
	if err != nil {
		return task, fmt.Errorf("error marshalling update task request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, reqUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		return task, fmt.Errorf("error building update task request: %w", err)
	}

	req.Header = c.buildHeaders()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return task, fmt.Errorf("error executing update task request: %w", err)
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return task, fmt.Errorf("error reading update task response: %w", err)
	}

	if err = json.Unmarshal(respBody, &task); err != nil {
		return task, fmt.Errorf("error unmarshalling update task response: %w", err)
	}

	return task, nil
}

func (c *Client) DeleteTask(ctx context.Context, taskID uint) error {
	reqUrl := url.URL{
		Scheme: "http",
		Host:   c.apiHost,
		Path:   fmt.Sprintf("/tasks/%d", taskID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, reqUrl.String(), nil)
	if err != nil {
		return fmt.Errorf("error building delete task request: %w", err)
	}

	req.Header = c.buildHeaders()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing delete task request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return nil
}

func (c *Client) ListFocusAreas(ctx context.Context) ([]focusareas.FocusAreaDTO, error) {
	reqUrl := url.URL{
		Scheme: "http",
		Host:   c.apiHost,
		Path:   "/focusareas",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error building list focus areas request: %w", err)
	}

	req.Header = c.buildHeaders()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing list focus areas request: %w", err)
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading list focus areas response: %w", err)
	}

	var focusAreasResp []focusareas.FocusAreaDTO
	if err = json.Unmarshal(respBody, &focusAreasResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling list focus areas response: %w", err)
	}

	return focusAreasResp, nil
}

func (c *Client) buildHeaders() http.Header {
	headers := http.Header{}

	if c.token != "" {
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}

	headers.Set("Content-Type", "application/json")

	return headers
}

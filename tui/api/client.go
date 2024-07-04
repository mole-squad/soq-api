package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/burkel24/task-app/pkg/tasks"
)

const APIHost = "localhost:3000"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		httpClient: c,
	}
}

func (c *Client) ListTasks(ctx context.Context) ([]tasks.TaskDTO, error) {
	reqUrl := url.URL{
		Scheme: "http",
		Host:   APIHost,
		Path:   "/tasks",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error building list tasks request: %w", err)
	}

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
		Host:   APIHost,
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
		Host:   APIHost,
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

	req.Header.Set("Content-Type", "application/json")

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

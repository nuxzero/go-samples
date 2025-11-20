package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Todo struct {
	ID        string
	Text      string
	CreatedAt time.Time
}

type TodoStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewTodoStore(redisAddr string) (*TodoStore, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &TodoStore{
		client: rdb,
		ctx:    ctx,
	}, nil
}

func (ts *TodoStore) CreateTodo(text string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("todo text cannot be empty")
	}

	id := uuid.New().String()
	now := time.Now()

	key := fmt.Sprintf("todo:%s", id)
	err := ts.client.HSet(ts.ctx, key, map[string]interface{}{
		"id":         id,
		"text":       text,
		"created_at": now.Format(time.RFC3339),
	}).Err()
	if err != nil {
		return "", fmt.Errorf("failed to create todo: %w", err)
	}

	err = ts.client.SAdd(ts.ctx, "todos:all", id).Err()
	if err != nil {
		return "", fmt.Errorf("failed to add index: %w", err)
	}

	return id, nil
}

func (ts *TodoStore) GetTodo(id string) (*Todo, error) {
	key := fmt.Sprintf("todo:%s", id)

	result, err := ts.client.HGetAll(ts.ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("todo not found")
	}

	createdAt, err := time.Parse(time.RFC3339, result["created_at"])
	if err != nil {
		return nil, fmt.Errorf("invalid created_at format: %w", err)
	}

	return &Todo{
		ID:        result["id"],
		Text:      result["text"],
		CreatedAt: createdAt,
	}, nil
}

func (ts *TodoStore) ListTodos() ([]*Todo, error) {
	ids, err := ts.client.SMembers(ts.ctx, "todos:all").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get todo list: %w", err)
	}

	var todos []*Todo
	for _, id := range ids {
		todo, err := ts.GetTodo(id)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (ts *TodoStore) UpdateTodo(id, text string) error {
	if text == "" {
		return fmt.Errorf("todo text cannot be empty")
	}

	key := fmt.Sprintf("todo:%s", id)

	exists, err := ts.client.Exists(ts.ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to check todo existence: %w", err)
	}
	if exists == 0 {
		return fmt.Errorf("todo not found")
	}

	err = ts.client.HSet(ts.ctx, key, "text", text).Err()
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	return nil
}

func (ts *TodoStore) DeleteTodo(id string) error {
	key := fmt.Sprintf("todo:%s", id)

	err := ts.client.Del(ts.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	err = ts.client.SRem(ts.ctx, "todos:all", id).Err()
	if err != nil {
		return fmt.Errorf("failed to remove from index: %w", err)
	}

	return nil
}

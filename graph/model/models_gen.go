// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

type Node interface {
	IsNode()
}

type BackwardPagination struct {
	Last   int     `json:"last"`
	Before *string `json:"before"`
}

type CreateTaskInput struct {
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type EdgeOrder struct {
	Key       *OrderKey      `json:"key"`
	Direction OrderDirection `json:"direction"`
}

type FilterCondition struct {
	FilterWord  *string    `json:"filterWord"`
	User        *string    `json:"user"`
	CreatedAt   *time.Time `json:"createdAt"`
	Deadline    *time.Time `json:"deadline"`
	IsCompleted *bool      `json:"isCompleted"`
}

type ForwardPagination struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

type OrderKey struct {
	Task *TaskOrderKey `json:"task"`
}

type PageCondition struct {
	Backward   *BackwardPagination `json:"backward"`
	Forward    *ForwardPagination  `json:"forward"`
	PageNumber int                 `json:"pageNumber"`
	Limit      *int                `json:"limit"`
}

type PageInfo struct {
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
}

type TaskConnection struct {
	PageInfo   *PageInfo   `json:"pageInfo"`
	Edges      []*TaskEdge `json:"edges"`
	TotalCount int         `json:"totalCount"`
}

func (TaskConnection) IsConnection() {}

type TaskEdge struct {
	Cursor string `json:"cursor"`
	Node   *Task  `json:"node"`
}

func (TaskEdge) IsEdge() {}

type UpdateTaskInput struct {
	ID          string     `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Deadline    *time.Time `json:"deadline"`
	IsCompleted *bool      `json:"isCompleted"`
}

type UpdateUserInput struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Twitter *string `json:"twitter"`
}

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Twitter *string `json:"twitter"`
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskOrderKey string

const (
	TaskOrderKeyCreatedAt TaskOrderKey = "CREATED_AT"
	TaskOrderKeyDeadline  TaskOrderKey = "DEADLINE"
)

var AllTaskOrderKey = []TaskOrderKey{
	TaskOrderKeyCreatedAt,
	TaskOrderKeyDeadline,
}

func (e TaskOrderKey) IsValid() bool {
	switch e {
	case TaskOrderKeyCreatedAt, TaskOrderKeyDeadline:
		return true
	}
	return false
}

func (e TaskOrderKey) String() string {
	return string(e)
}

func (e *TaskOrderKey) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskOrderKey(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskOrderKey", str)
	}
	return nil
}

func (e TaskOrderKey) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

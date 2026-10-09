package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusIncomplete Status = "incomplete"
	StatusComplete   Status = "complete"
)

type Node struct {
	ID        uuid.UUID `json:"id"`
	GraphID   uuid.UUID `json:"graph_id"`
	ParentID  *uuid.UUID `json:"parent_id"`
	Name      string    `json:"name"`
	Status    Status    `json:"status"`
	Position  uint16    `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNode(graphID uuid.UUID, parentID *uuid.UUID, name string, position uint16) *Node {
	return &Node{
		ID:        uuid.New(),
		GraphID:   graphID,
		ParentID:  parentID,
		Name:      name,
		Status:    StatusIncomplete,
		Position:  position,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (node *Node) Complete() {
	node.Status = StatusComplete
	node.Touch()
}

func (node *Node) Reopen() {
	node.Status = StatusIncomplete
	node.Touch()
}

func (node *Node) Rename(name string) {
	node.Name = name
	node.Touch()
}

func (node *Node) Touch() {
	node.UpdatedAt = time.Now()
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type KnowledgeTree struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewKnowledgeTree(name string) *KnowledgeTree {
	return &KnowledgeTree{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *KnowledgeTree) Touch() {
	t.UpdatedAt = time.Now()
}

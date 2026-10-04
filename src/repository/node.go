package repository

import (
	"jakegodsall/knowledge-graph/src/domain"

	"github.com/google/uuid"
)

type NodeRepository interface {
	Create(node *domain.Node) error
	FindByID(id uuid.UUID) (*domain.Node, error)
	FindChildren(parentID uuid.UUID) ([]*domain.Node, error)
	Update(node *domain.Node) error
	DeleteByID(id uuid.UUID) error
}

package repository

import (
	"jakegodsall/knowledge-graph/src/domain"

	"github.com/google/uuid"
)

type NodeRepository interface {
	Create(node *domain.Node) error
	FindByID(id uuid.UUID) (*domain.Node, error)
	// FindByIDPrefix returns every node whose ID starts with prefix, so short
	// IDs can be used on the command line.
	FindByIDPrefix(prefix string) ([]*domain.Node, error)
	FindByTree(graphID uuid.UUID) ([]*domain.Node, error)
	FindRoots(graphID uuid.UUID) ([]*domain.Node, error)
	FindChildren(parentID uuid.UUID) ([]*domain.Node, error)
	Update(node *domain.Node) error
	// DeleteByID deletes the node and all of its descendants.
	DeleteByID(id uuid.UUID) error
}

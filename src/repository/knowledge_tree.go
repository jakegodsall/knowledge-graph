package repository

import (
	"jakegodsall/knowledge-graph/src/domain"

	"github.com/google/uuid"
)

type KnowledgeTreeRepository interface {
	GetAll() (trees []*domain.KnowledgeTree, error)
	Create(tree *domain.KnowledgeTree) error
	FindByID(id uuid.UUID) (*domain.KnowledgeTree, error)
	DeleteByID(id uuid.UUID) error
}

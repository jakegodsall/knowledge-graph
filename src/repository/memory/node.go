package memory

import (
	"fmt"
	"jakegodsall/knowledge-graph/src/domain"

	"github.com/google/uuid"
)

type NodeRepository struct {
	nodes []*domain.Node
}

func NewNodeRepository() *NodeRepository {
	return &NodeRepository{}
}

func (r *NodeRepository) Create(node *domain.Node) error {
	r.nodes = append(r.nodes, node)
	return nil
}

func (r *NodeRepository) FindByID(id uuid.UUID) (*domain.Node, error) {
	for _, n := range r.nodes {
		if n.ID == id {
			return n, nil
		}
	}

	return nil, fmt.Errorf("no node found for id %s", id.String())
}

func (r *NodeRepository) FindChildren(parentID uuid.UUID) ([]*domain.Node, error) {
	children := []*domain.Node{}

	for _, n := range r.nodes {
		if n.ParentID != nil && *n.ParentID == parentID {
			children = append(children, n)
		}
	}

	return children, nil
}

func (r *NodeRepository) Update(node *domain.Node) error {
	for i, n := range r.nodes {
		if n.ID == node.ID {
			r.nodes[i] = node
			return nil
		}
	}

	return fmt.Errorf("no node found for id %s", node.ID.String())
}

func (r *NodeRepository) DeleteByID(id uuid.UUID) error {
	idx := -1

	for i, n := range r.nodes {
		if n.ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		return fmt.Errorf("no node found for id %s", id.String())
	}

	r.nodes = append(r.nodes[:idx], r.nodes[idx+1:]...)
	return nil
}

package memory

import (
	"fmt"
	"jakegodsall/knowledge-graph/src/domain"
	"sort"
	"strings"

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

func (r *NodeRepository) FindByIDPrefix(prefix string) ([]*domain.Node, error) {
	prefix = strings.ToLower(prefix)
	found := []*domain.Node{}

	for _, n := range r.nodes {
		if strings.HasPrefix(n.ID.String(), prefix) {
			found = append(found, n)
		}
	}

	return found, nil
}

func (r *NodeRepository) FindByTree(graphID uuid.UUID) ([]*domain.Node, error) {
	return r.filter(func(n *domain.Node) bool {
		return n.GraphID == graphID
	}), nil
}

func (r *NodeRepository) FindRoots(graphID uuid.UUID) ([]*domain.Node, error) {
	return r.filter(func(n *domain.Node) bool {
		return n.GraphID == graphID && n.ParentID == nil
	}), nil
}

func (r *NodeRepository) FindChildren(parentID uuid.UUID) ([]*domain.Node, error) {
	return r.filter(func(n *domain.Node) bool {
		return n.ParentID != nil && *n.ParentID == parentID
	}), nil
}

// filter returns matching nodes ordered by position, like the SQLite queries.
func (r *NodeRepository) filter(match func(*domain.Node) bool) []*domain.Node {
	found := []*domain.Node{}

	for _, n := range r.nodes {
		if match(n) {
			found = append(found, n)
		}
	}

	sort.SliceStable(found, func(i, j int) bool {
		return found[i].Position < found[j].Position
	})

	return found
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
	if _, err := r.FindByID(id); err != nil {
		return err
	}

	subtree := map[uuid.UUID]bool{id: true}

	// keep sweeping until no more descendants are found
	for added := true; added; {
		added = false

		for _, n := range r.nodes {
			if n.ParentID != nil && subtree[*n.ParentID] && !subtree[n.ID] {
				subtree[n.ID] = true
				added = true
			}
		}
	}

	remaining := []*domain.Node{}

	for _, n := range r.nodes {
		if !subtree[n.ID] {
			remaining = append(remaining, n)
		}
	}

	r.nodes = remaining
	return nil
}

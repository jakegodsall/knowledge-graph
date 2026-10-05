package sqlite

import (
	"database/sql"
	"fmt"
	"jakegodsall/knowledge-graph/src/domain"

	"github.com/google/uuid"
)

type NodeRepository struct {
	db *sql.DB
}

func NewNodeRepository(db *sql.DB) *NodeRepository {
	return &NodeRepository{
		db: db,
	}
}

func (r *NodeRepository) Create(node *domain.Node) error {
	res, err := r.db.Exec(`
			INSERT INTO nodes (id, graph_id, parent_id, name, status, position, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`,
		&node.ID,
		&node.GraphID,
		&node.ParentID,
		&node.Name,
		&node.Status,
		&node.Position,
		&node.CreatedAt,
		&node.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("could not create the node %s", node.ID.String())
	}

	return nil
}

func (r *NodeRepository) FindByID(id uuid.UUID) (*domain.Node, error) {
	row := r.db.QueryRow(`
			SELECT id, graph_id, parent_id, name, status, position, created_at, updated_at
			FROM nodes
			WHERE id = ?
		`,
		id,
	)

	node := &domain.Node{}

	if err := row.Scan(
		&node.ID,
		&node.GraphID,
		&node.ParentID,
		&node.Name,
		&node.Status,
		&node.Position,
		&node.CreatedAt,
		&node.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no node found for id %s", id)
		}
		return nil, err
	}

	return node, nil
}

func (r *NodeRepository) FindChildren(parentID uuid.UUID) ([]*domain.Node, error) {
	rows, err := r.db.Query(`
			SELECT id, graph_id, parent_id, name, status, position, created_at, updated_at
			FROM nodes
			WHERE parent_id = ?
			ORDER BY position
		`,
		parentID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	nodes := []*domain.Node{}

	for rows.Next() {
		node := &domain.Node{}

		if err := rows.Scan(
			&node.ID,
			&node.GraphID,
			&node.ParentID,
			&node.Name,
			&node.Status,
			&node.Position,
			&node.CreatedAt,
			&node.UpdatedAt,
		); err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nodes, nil
}

func (r *NodeRepository) Update(node *domain.Node) error {
	res, err := r.db.Exec(`
			UPDATE nodes
			SET graph_id = ?, parent_id = ?, name = ?, status = ?, position = ?, updated_at = ?
			WHERE id = ?
		`,
		&node.GraphID,
		&node.ParentID,
		&node.Name,
		&node.Status,
		&node.Position,
		&node.UpdatedAt,
		&node.ID,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("node %s could not be updated", node.ID.String())
	}

	return nil
}

func (r *NodeRepository) DeleteByID(id uuid.UUID) error {
	res, err := r.db.Exec("DELETE FROM nodes WHERE id = ?", id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("node %s could not be deleted", id.String())
	}

	return nil
}

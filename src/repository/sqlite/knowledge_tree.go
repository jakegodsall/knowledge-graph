package sqlite

import (
	"database/sql"
	"fmt"
	"jakegodsall/knowledge-graph/src/domain"

	"github.com/google/uuid"
)

type KnowledgeTreeRepository struct {
	db *sql.DB
}

func NewKnowledgeTreeRepository(db *sql.DB) *KnowledgeTreeRepository {
	return &KnowledgeTreeRepository{
		db: db,
	}
}

func (r *KnowledgeTreeRepository) GetAll() ([]*domain.KnowledgeTree, error) {
	rows, err := r.db.Query(`
		SELECT id, name, created_at, updated_at
		FROM knowledge_trees
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	trees := []*domain.KnowledgeTree{}

	for rows.Next() {
		tree := &domain.KnowledgeTree{}

		if err := rows.Scan(
			&tree.ID,
			&tree.Name,
			&tree.CreatedAt,
			&tree.UpdatedAt,
		); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trees, nil
}

func (r *KnowledgeTreeRepository) Create(tree *domain.KnowledgeTree) error {
	res, err := r.db.Exec(`
			INSERT INTO knowledge_trees (id, name, created_at, updated_at)
			VALUES (?, ?, ?, ?)
		`,
		tree.ID,
		tree.Name,
		tree.CreatedAt,
		tree.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("error creating knowledge tree %s", tree.ID.String())
	}

	return nil
}

func (r *KnowledgeTreeRepository) FindByID(id uuid.UUID) (*domain.KnowledgeTree, error) {
	row := r.db.QueryRow(`
			SELECT id, name, created_at, updated_at
			FROM knowledge_trees
			WHERE id = ?
		`,
		id,
	)

	tree := &domain.KnowledgeTree{}

	if err := row.Scan(
		&tree.ID,
		&tree.Name,
		&tree.CreatedAt,
		&tree.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no knowledge tree fund for id %s", id.String())
		}
		return nil, err
	}

	return tree, nil
}

func (r *KnowledgeTreeRepository) DeleteByID(id uuid.UUID) error {
	res, err := r.db.Exec("DELETE FROM knowledge_trees WHERE id = ?", id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("error deleting knowledge tree %s", id.String())
	}

	return nil
}

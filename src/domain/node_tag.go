package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type NodeTag struct {
	NodeID    uuid.UUID `json:"node_id"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNodeTag(nodeID uuid.UUID, tag string) (*NodeTag, error) {
	tag = NormaliseTag(tag)

	if tag == "" {
		return nil, fmt.Errorf("tag for node %s cannot be empty", nodeID.String())
	}

	return &NodeTag{
		NodeID:    nodeID,
		Tag:       tag,
		CreatedAt: time.Now(),
	}, nil
}

// NormaliseTag lowercases the tag and joins its words with hyphens, so
// " Exam  Topic " becomes "exam-topic".
func NormaliseTag(tag string) string {
	return strings.ToLower(strings.Join(strings.Fields(tag), "-"))
}

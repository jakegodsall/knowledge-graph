package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var ErrPrerequisiteCycle = errors.New("prerequisite would create a cycle")

type Prerequisite struct {
	NodeID         uuid.UUID `json:"node_id"`
	RequiresNodeID uuid.UUID `json:"requires_node_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewPrerequisite(nodeID uuid.UUID, requiresNodeID uuid.UUID) (*Prerequisite, error) {
	if nodeID == requiresNodeID {
		return nil, fmt.Errorf("node %s cannot be a prerequisite of itself", nodeID.String())
	}

	return &Prerequisite{
		NodeID:         nodeID,
		RequiresNodeID: requiresNodeID,
		CreatedAt:      time.Now(),
	}, nil
}

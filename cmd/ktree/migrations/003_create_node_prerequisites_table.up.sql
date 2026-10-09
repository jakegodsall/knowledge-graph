CREATE TABLE node_prerequisites (
    node_id TEXT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
    requires_node_id TEXT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
    created_at DATETIME NOT NULL,
    PRIMARY KEY (node_id, requires_node_id),
    CHECK (node_id <> requires_node_id)
);

CREATE INDEX idx_node_prerequisites_requires_node_id ON node_prerequisites (requires_node_id);

CREATE TABLE node_tags (
    node_id TEXT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
    tag TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    PRIMARY KEY (node_id, tag)
);

CREATE INDEX idx_node_tags_tag ON node_tags (tag);

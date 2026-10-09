CREATE TABLE nodes (
    id TEXT PRIMARY KEY NOT NULL,
    graph_id TEXT NOT NULL,
    parent_id TEXT,
    name TEXT NOT NULL,
    status TEXT NOT NULL,
    position INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

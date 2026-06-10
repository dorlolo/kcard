CREATE TABLE IF NOT EXISTS knowledge_relationships (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    learner_workspace_id uuid NOT NULL REFERENCES learner_workspaces(id),
    source_knowledge_point_id uuid REFERENCES knowledge_points(id),
    target_knowledge_point_id uuid REFERENCES knowledge_points(id),
    relationship_type text NOT NULL,
    label text,
    weight numeric NOT NULL DEFAULT 1,
    source_type text NOT NULL DEFAULT 'system_derived',
    source_material_id uuid REFERENCES source_materials(id),
    tag_id uuid REFERENCES tags(id),
    card_id uuid,
    confidence numeric,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    archived_at timestamptz
);

CREATE INDEX IF NOT EXISTS idx_knowledge_relationships_workspace_type
    ON knowledge_relationships(learner_workspace_id, relationship_type);

CREATE INDEX IF NOT EXISTS idx_knowledge_relationships_source_target
    ON knowledge_relationships(source_knowledge_point_id, target_knowledge_point_id);

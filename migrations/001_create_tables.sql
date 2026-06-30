CREATE TABLE IF NOT EXISTS form_definitions
(
    id UUID PRIMARY KEY,

    name TEXT NOT NULL,

    description TEXT,

    created_at TIMESTAMP NOT NULL,

    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS form_template_versions
(
    id UUID PRIMARY KEY,

    definition_id UUID NOT NULL
        REFERENCES form_definitions(id)
        ON DELETE CASCADE,

    version INTEGER NOT NULL,

    status VARCHAR(20) NOT NULL,

    ui_schema JSONB NOT NULL,

    validation_schema JSONB NOT NULL,

    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS form_submissions
(
    id UUID PRIMARY KEY,

    template_version_id UUID NOT NULL
        REFERENCES form_template_versions(id)
        ON DELETE CASCADE,

    data JSONB NOT NULL,

    created_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_template_definition
ON form_template_versions(definition_id);

CREATE INDEX idx_submission_template
ON form_submissions(template_version_id);

CREATE INDEX idx_submission_data
ON form_submissions
USING GIN(data);
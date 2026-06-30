-- ============================================
-- Seed Form Definition
-- ============================================

INSERT INTO form_definitions
(
    id,
    name,
    description,
    created_at,
    updated_at
)
VALUES
(
    '11111111-1111-1111-1111-111111111111',
    'Equipment Registration',
    'Demo dynamic form for the assessment',
    NOW(),
    NOW()
);

-- ============================================
-- Seed Published Template
-- ============================================

INSERT INTO form_template_versions
(
    id,
    definition_id,
    version,
    status,
    ui_schema,
    validation_schema,
    created_at
)
VALUES
(
    '22222222-2222-2222-2222-222222222222',

    '11111111-1111-1111-1111-111111111111',

    1,

    'PUBLISHED',

    '{
        "title":"Equipment Registration",
        "fields":[
            {
                "id":"equipmentName",
                "label":"Equipment Name",
                "type":"text",
                "placeholder":"Generator"
            },
            {
                "id":"serialNumber",
                "label":"Serial Number",
                "type":"text"
            },
            {
                "id":"capacity",
                "label":"Capacity (kVA)",
                "type":"number"
            },
            {
                "id":"location",
                "label":"Location",
                "type":"text"
            }
        ]
    }'::jsonb,

    '{
        "$schema":"https://json-schema.org/draft/2020-12/schema",
        "type":"object",
        "required":[
            "equipmentName",
            "serialNumber",
            "capacity"
        ],
        "properties":{
            "equipmentName":{
                "type":"string"
            },
            "serialNumber":{
                "type":"string"
            },
            "capacity":{
                "type":"number"
            },
            "location":{
                "type":"string"
            }
        }
    }'::jsonb,

    NOW()
);
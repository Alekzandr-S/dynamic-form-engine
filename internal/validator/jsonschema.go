package validator

import (
	"bytes"
	"encoding/json"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Validator interface {
	Validate(schema json.RawMessage, data map[string]any) error
}

type JSONSchemaValidator struct{}

func NewJSONSchemaValidator() *JSONSchemaValidator {
	return &JSONSchemaValidator{}
}

func (v *JSONSchemaValidator) Validate(schema json.RawMessage, data map[string]any) error {
	compiler := jsonschema.NewCompiler()

	if err := compiler.AddResource(
		"schema.json",
		bytes.NewReader(schema),
	); err != nil {
		return err
	}

	compiled, err := compiler.Compile("schema.json")
	if err != nil {
		return err
	}

	return compiled.Validate(data)
}

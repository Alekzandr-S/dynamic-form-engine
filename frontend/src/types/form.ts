export type FieldType =
  | "text"
  | "number"
  | "checkbox"
  | "select"
  | "textarea";

export interface Field {
  id: string;
  label: string;
  type: FieldType;
  placeholder?: string;
  required?: boolean;
  options?: string[];
}

export interface UISchema {
  title: string;
  fields: Field[];
}

export interface FormVersion {
  id: string;
  definitionId: string
  version: number;
  status: string;
  uiSchema: UISchema;
  validationSchema: unknown;
  createdAt: string;
}

export interface FormDefinition {
  id: string;
  name: string;
  description: string;
}
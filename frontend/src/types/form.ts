export interface Field {
  id: string;
  label: string;
  type: string;
  placeholder?: string;
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
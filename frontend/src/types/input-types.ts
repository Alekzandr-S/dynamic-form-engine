import type { Field } from "./form";

export interface FieldProps<T> {
  field: Field;
  value: T;
  onChange: (value: T) => void;
}

export interface Props {
  field: Field;
  value: unknown;
  onChange: (value: unknown) => void;
}
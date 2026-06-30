// import type { Props } from "../types/input-types";
import type { Field } from "../types/form";
import { fieldRegistry } from "./fields";

interface Props {
  field: Field;
  value: unknown;
  onChange: (value: unknown) => void;
}

export default function DynamicField({field, value, onChange,}: Props) {
  const Component = fieldRegistry[field.type];
  
  if (!Component) {
    return <p>Unsupported field type: {field.type}</p>;
  }

  return (
    <Component field={field} value={value} onChange={onChange} />
  )
}
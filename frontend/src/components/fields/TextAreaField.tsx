import type { Props } from "../../types/input-types";

export default function TextArea({field, value, onChange,}: Props) {
  return (
    <div>
      <label>{field.label}</label>
      <input
        type="text"
        value={(value as string) ?? ""}
        placeholder={field.placeholder}
        onChange={(e) => onChange(String(e.target.value))}
      />
    </div>
  )
}
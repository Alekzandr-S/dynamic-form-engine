import type {Props } from "../../types/input-types";


export default function Numberfield({field, value, onChange,}: Props) {
  return (
    <div>
      <label>{field.label}</label>
      <input
        type="number"
        value={(value as number) ?? 0}
        placeholder={field.placeholder}
        onChange={(e) => onChange(Number(e.target.value))}
      />
    </div>
  )
} 
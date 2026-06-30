import type { Props } from "../../types/input-types";

export default function CheckboxField({field, value, onChange}: Props) {
  return (
    <div>
      <label>
        <input 
          type="boolean" 
          checked={(value as boolean) ?? false}
          onChange={(e) => onChange(Boolean(e.target.checked))}
        />
        {field.label}
      </label>
    </div>
  )
}
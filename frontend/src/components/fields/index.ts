import CheckboxField from "./CheckboxField";
import Numberfield from "./NumberField";
import SelectField from "./SelectField";
import TextArea from "./TextAreaField";
import TextField from "./TextField";


export const fieldRegistry = {
  text: TextField,
  number: Numberfield,
  checkbox: CheckboxField,
  select: SelectField,
  textarea: TextArea,
} as const;
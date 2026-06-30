import { useState } from "react";
import type { UISchema } from "../types/form";
import api from "../api/api";
import DynamicField from "./DynamicField";

interface Props {
  definitionId: string;
  schema: UISchema;
}

export default function DynamicForm({definitionId, schema}: Props) {
  const [formData, setFormData] = useState<Record<string, unknown>>({});
  // const [submitted, setSubmitted] = useState(false);

  const handleChange = (id: string, value: unknown,) => {
    setFormData(prev => ({
      ...prev, [id]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent,) => {
    e.preventDefault();
    try {
      await api.post(`/forms/${definitionId}/submissions`, formData,);

      alert("Submitted successfully!");
      // setSubmitted(true)
      setFormData({});
    } catch (err) {
      console.error(err);

      alert("Submission failed.");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>{schema.title}</h2>
      {schema.fields.map(field => (
        <DynamicField
          key={field.id}
          field={field}
          value={formData[field.id]}
          onChange={(value) => handleChange(field.id, value)}
        />
        // {submitted && (
          // <p style={{color: "green"}}>
            // Form Submitted successfully.
          // </p>
        // )}
      ))}
      <button type="submit">Submit</button>
    </form>
  )
}
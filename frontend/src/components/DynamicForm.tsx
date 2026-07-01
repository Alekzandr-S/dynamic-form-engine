import { useState } from "react";
import type { UISchema } from "../types/form";
import api from "../api/api";
import DynamicField from "./DynamicField";
import { Button } from "./ui/button";
import { Link } from "react-router-dom";

interface Props {
  definitionId: string;
  schema: UISchema;
}

export default function DynamicForm({definitionId, schema}: Props) {
  const [formData, setFormData] = useState<Record<string, unknown>>({});
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState("");
  
  const handleChange = (id: string, value: unknown,) => {
    setFormData(prev => ({
      ...prev, [id]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent,) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setSuccess(false);
    try {
      await api.post(`/forms/${definitionId}/submissions`, formData,);

      setSuccess(true);
      setFormData({});
    } catch (err) {
      console.error(err);

      setError("Submission failed.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-slate-100">
      <div className="mx-auto max-w-2xl p-8">
        {
          success && (
            <div
            className="mb-4 rounded bg-green-100 border border-green-300 p-3 text-green-700"
            >
              Submission successful!.
              <Link to={"/"}>Return to Forms</Link>
            </div>
          )
        }
        {
          error && (
            <div
              className="mb-4 rounded bg-red-100 border border-red-300 p-3 text-red 700"
            >
              {error}
            </div>
          )
        }
        <form className="rounded-lg bg-white shadow p-8" noValidate onSubmit={handleSubmit}>
          <h2>{schema.title}</h2>
          {schema.fields.map(field => (
            <DynamicField
              key={field.id}
              field={field}
              value={formData[field.id] ?? ""}
              onChange={(value) => handleChange(field.id, value)}
            />
          ))}
          <Button type="submit" disabled={loading} className="w-full">
            {loading ? "Submitting..." : "Submit"}
          </Button>
        </form>
      </div>
    </div>
  )
}
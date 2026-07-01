import api from "@/api/api";
import type { FormDefinition } from "@/types/form";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
// import DynamicForm from "../components/DynamicForm";

export default function Home() {
  const [definitions, setDefinitions] = useState<FormDefinition[]>([])
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    api.get("definitions")
      .then(res => {
        setDefinitions(res.data)
      })
      .finally(() => setLoading(false));
  }, [])

  if (loading) {
    return <h2>Loading forms...</h2>
  }

  return (
    <div>
      <h1>Available Forms</h1>
      {definitions.map(definition => (
        <Link
          key={definition.id}
          to={`/forms/${definition.id}`}
          className="block rounded border p-4 mb-4 hover:bg-gray-100"
        >
          <h2 className="font-semibold">{definition.name}</h2>
          <p>{definition.description}</p>
        </Link>
      ))}
    </div>
  )
}
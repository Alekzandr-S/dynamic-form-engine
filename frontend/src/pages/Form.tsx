import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import type { FormVersion } from "../types/form";
import api from "../api/api";
import DynamicForm from "../components/DynamicForm";


export default function Form() {
  const {id} = useParams();
  const [form, setForm] = useState<FormVersion | null>(null);

  useEffect(() =>{

    api.get(`/forms/${id}`)
      .then(res => {
        console.log(res.data)
        setForm(res.data)});
        // setForm(res.data.uiSchema)});
  }, [id])

  if (!form)
    return <h2>Loading...</h2>;

  return (
    <div>
      <DynamicForm definitionId={id!} schema={form.uiSchema} />
    </div>
  )
}
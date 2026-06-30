import { Link } from "react-router-dom";
// import DynamicForm from "../components/DynamicForm";

export default function Home() {
  const demoId = import.meta.env.VITE_DEMO_FORM_ID;

  return (
    <div>
      {/* <DynamicForm definitionId={demoId} schema={form.uiSchema} /> */}
      <h1>Dynamic Form Engine</h1>
      <Link to={`/forms/${demoId}`}>
        Open Demo Form
      </Link>
    </div>
  )
}
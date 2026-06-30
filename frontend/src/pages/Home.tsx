import { Link } from "react-router-dom";

export default function Home() {
  const demoId = import.meta.env.VITE_DEMO_FORM_ID;

  return (
    <div>
      <h1>Dynamic Form Engine</h1>
      <Link to={`/forms${demoId}`}>
        Open Demo Form
      </Link>
    </div>
  )
}
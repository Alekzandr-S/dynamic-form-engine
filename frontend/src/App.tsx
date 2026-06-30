import {Routes, Route} from "react-router-dom";
import Home from "./pages/Home";
import Form from "./pages/Form";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/forms/:id" element={<Form />} />
    </Routes>
  )
}

export default App;
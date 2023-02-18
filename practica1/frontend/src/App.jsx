import { BrowserRouter, Route, Routes } from "react-router-dom";
import Calculadora from "./components/Calculadora/Calculadora";
import Logs from "./components/Reports/Logs";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Calculadora/>}></Route>
        <Route path="logs" element={<Logs/>}></Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;

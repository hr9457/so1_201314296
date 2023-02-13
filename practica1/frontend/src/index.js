import react from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import ReactDOM from "react-dom/client";
// import AutoLayoutExample from "./components/Conteiner"
import App from "./App";

const root = ReactDOM.createRoot(document.getElementById("root"));

function Greeting() {
  const name = "SOPES1 - CALCULADORA";
  return <h1>{name}</h1>;
}

root.render(
  <>
    <Greeting></Greeting>
    <App/>
  </>
);

import react from "react";
import ReactDOM from "react-dom/client";

const root = ReactDOM.createRoot(document.getElementById("root"));

function Greeting() {
  const name = "Hector";
  return <h1>{name}</h1>;
}

root.render(
  <>
    <Greeting></Greeting>
  </>
);

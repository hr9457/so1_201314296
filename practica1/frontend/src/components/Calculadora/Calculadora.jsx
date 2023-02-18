import "./Calculadora.css";
import { evaluate, prodDependencies } from "mathjs";
import Pantalla from "../Pantalla/Pantalla";
import Boton from "../Button/Boton";
import Clear from "../Clear/Clear";
import NavbarMenu from "../Menu/Menu";

// api
import { suma, operar } from "../../services/api";

// hook
import { useState } from "react";

function Calculadora() {
  // estado de la pantalla
  const [input, setInput] = useState("");
  const [value1, setValue1] = useState("");
  const [value2, setValue2] = useState("");
  const [operation, setOperation] = useState("");

  // para agregar valores a la pantalla
  const agregarInput = (val) => {
    setInput(input + val);
  };

  // funcion para calcular resultado por medio del =
  const calcularResultado = async () => {
    if (input) {
      const responseJson = await operar(value1, input, operation);
      setInput(responseJson);
    } else {
      alert("Por favor ingrese valores para realizar los calculos");
    }
  };

  // puedo recibir el valor del boton
  const botonOperacion = (val) => {
    setOperation(val);
    if (value1 == "") setValue1(input);
    setInput("");
  };

  // manejar boton de clear
  const botonClear = () => {
    setInput("");
    setValue1("");
    setValue2("");
    setOperation("");
  };

  return (
    <>
      <NavbarMenu></NavbarMenu>

      <div className="App">
        <div className="contenedor-calculadora">
          <Pantalla input={input} />

          <div className="fila">
            <Boton manejarClic={agregarInput}>1</Boton>
            <Boton manejarClic={agregarInput}>2</Boton>
            <Boton manejarClic={agregarInput}>3</Boton>
            <Boton manejarClic={botonOperacion}>+</Boton>
          </div>

          <div className="fila">
            <Boton manejarClic={agregarInput}>4</Boton>
            <Boton manejarClic={agregarInput}>5</Boton>
            <Boton manejarClic={agregarInput}>6</Boton>
            <Boton manejarClic={botonOperacion}>-</Boton>
          </div>

          <div className="fila">
            <Boton manejarClic={agregarInput}>7</Boton>
            <Boton manejarClic={agregarInput}>8</Boton>
            <Boton manejarClic={agregarInput}>9</Boton>
            <Boton manejarClic={botonOperacion}>*</Boton>
          </div>

          <div className="fila">
            <Boton manejarClic={calcularResultado}>=</Boton>
            <Boton manejarClic={agregarInput}>0</Boton>
            <Boton manejarClic={agregarInput}>.</Boton>
            <Boton manejarClic={botonOperacion}>/</Boton>
          </div>

          <div className="fila">
            <Clear manejarClear={botonClear}>Clear</Clear>
          </div>
        </div>
      </div>
    </>
  );
}

export default Calculadora;

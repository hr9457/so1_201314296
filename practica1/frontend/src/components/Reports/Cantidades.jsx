import React from "react";
import { useState, useEffect } from "react";

import NavbarMenu from "../Menu/Menu";
import Table from 'react-bootstrap/Table';
import { historial } from "../../services/api";

export default function Cantidades() {

  const [datos, setDatos] = useState([]);


  const consumo = async () => {
    const response = await historial();
    setDatos(response);
  };


  return (
    <>
      <NavbarMenu></NavbarMenu>

      <h1>Reporte de Logs</h1>

      {/* TABLA DE REPORTES */}
      <Table striped bordered hover>
      <thead>
        <tr>
          <th>#</th>
          <th>Numero 1</th>
          <th>Numero 2</th>
          <th>Operador</th>
          <th>Resultado</th>
          <th>Fecha</th>
        </tr>
      </thead>

      
    </Table>
    </>
  );
}

import React from "react";
import { useState, useEffect } from "react";

import NavbarMenu from "../Menu/Menu";
import Table from 'react-bootstrap/Table';
import { historial } from "../../services/api";

export default function Logs() {

  const [datos, setDatos] = useState([]);


  const consumo = async () => {
    const response = await historial();
    setDatos(response);
  };


  useEffect( () => {
    consumo();
  },[]);



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

      <tbody>
        {datos.map(log => {
          return(
            <tr key={log.id_historial}>
              <td>{log.id_historial}</td>
              <td>{log.numero1}</td>
              <td>{log.numero2}</td>
              <td>{log.operador}</td>
              <td>{log.resultado}</td>
              <td>{log.fecha}</td>              
            </tr>
          )
        })}
      </tbody>
      
    </Table>
    </>
  );
}

import React, { useEffect, useState } from "react";
import Accordion from "react-bootstrap/Accordion";
import Table from "react-bootstrap/Table";

const AcordinInfo = () => {
  const [datos, setDatos] = useState([]);
  const [padre, setPadre] = useState("root");

  const cpu = async () => {
    const response = await fetch(`${process.env.REACT_APP_API_URL}/api/cpu`, {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    });
    const responseJson = await response.json();
    const data = responseJson.data;
    setDatos(data);
  };

  useEffect(() => {
    setTimeout(() => {
      cpu();
    }, 5000);
  }, []);

  return (
    <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
        <Accordion.Header>Informacion General</Accordion.Header>
        <Accordion.Body>
          <Table striped bordered hover>
            <thead>
              <tr>
                <th>#</th>
                <th>Padre</th>
                <th>PID</th>
                <th>Usuario</th>
                <th>Estado</th>
                <th>%RAM (MB)</th>
              </tr>
            </thead>

            <tbody>
              {datos.map((log) => {
                return (
                  
                  log.hijos.map((child) =>{
                    return(
                      <tr>
                        <td>{log.PID}</td>
                        <td>{log.nombre}</td>
                        <td>{child.PID}</td>
                        <td>{child.nombre}</td>
                        <td>{child.estado}</td>
                        <td>{child.vmrss}mb</td>
                      </tr>
                    )})                                 

                );
              })}
              {/* <tr>
                <td>Numero</td>
                <td>root</td>
                <td>0</td>
                <td>1</td>
                <td>systemd</td>
                <td>sleeping</td>
                <td>13209600</td>
              </tr> */}
            </tbody>
          </Table>
        </Accordion.Body>
      </Accordion.Item>
    </Accordion>
  );
};

export default AcordinInfo;

import React, { memo, useEffect, useState } from 'react'
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Pie } from 'react-chartjs-2';


import Card from "react-bootstrap/Card";

ChartJS.register(ArcElement, Tooltip, Legend);


// initial data state
export const dataState = {
    labels: ['Libre', 'Usada'],
    datasets: [
      {
        label: '%',
        data: [50, 50],
        backgroundColor: [
            'rgba(54, 162, 235, 0.2)',
            'rgba(255, 99, 132, 0.2)',          
          ],
          borderColor: [
            'rgba(54, 162, 235, 1)',
            'rgba(255, 99, 132, 1)',          
          ],
        borderWidth: 3,
      },
    ],
  };


const PieChartCpu = () => {

  const [dataPie, setdataPie ] = useState(dataState);
  const [ejecuccion, setEjecuccion] = useState(0);
  const [suspendidos, setSuspendidos] = useState(0);
  const [detenidos, setDetenidos] = useState(0);
  const [zombies, setZombies] = useState(0);
  const [totales, setTotales] = useState(0);


  const memory = async () => {
    const response = await fetch(`${process.env.REACT_APP_API_URL}/api/cpu`, {
      method: 'GET',
      headers: {
        'Accept':'application/json',
        'Content-Type': 'application/json',
      }    
    });
    const responseJson = await response.json();
    setEjecuccion(responseJson.ejecuccion);
    setSuspendidos(responseJson.suspendidos);
    setDetenidos(responseJson.detenidos);
    setZombies(responseJson.zombies);
    setTotales(responseJson.total);
    setdataPie(
      {
        labels: ['Libre', 'Usada'],
        datasets: [
          {
            label: '%',
            data: [100-responseJson.porcentaje, responseJson.porcentaje],
            backgroundColor: [
                'rgba(54, 162, 235, 0.2)',
                'rgba(255, 99, 132, 0.2)',          
              ],
              borderColor: [
                'rgba(54, 162, 235, 1)',
                'rgba(255, 99, 132, 1)',          
              ],
            borderWidth: 3,
          },
        ],
      }
    );
  };
  
  useEffect(() => {
    setTimeout(() => {
      memory();
    }, 5000);    
  },[dataPie]);


  return (
    <>
      <Card style={{ width: "18rem" }}>
        <Pie data={dataPie} />
        <Card.Body>
          <Card.Title>RAM</Card.Title>
          <Card.Text>{`Ejecuccion: ${ejecuccion} Suspendidos: ${suspendidos} Detenidos: ${detenidos} Zombies: ${zombies} Totales: ${totales}`}</Card.Text>
        </Card.Body>
      </Card>
    </>
  )
}

export default PieChartCpu
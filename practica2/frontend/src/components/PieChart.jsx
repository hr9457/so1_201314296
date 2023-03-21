import React, { memo, useEffect, useState } from "react";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Pie } from "react-chartjs-2";


import Card from "react-bootstrap/Card";

ChartJS.register(ArcElement, Tooltip, Legend);

// initial data state
export const dataState = {
  labels: ["Libre", "Usada"],
  datasets: [
    {
      label: "%",
      data: [50, 50],
      backgroundColor: ["rgba(54, 162, 235, 0.2)", "rgba(255, 99, 132, 0.2)"],
      borderColor: ["rgba(54, 162, 235, 1)", "rgba(255, 99, 132, 1)"],
      borderWidth: 3,
    },
  ],
};




const PieChart = () => {
  const [dataPie, setdataPie] = useState(dataState);
  const [total, setTotal] = useState(0);
  const [consumida, setConsumidad] = useState(0);
  const [usada, setUsada] = useState(0);

  const memory = async () => {
    const response = await fetch(`${process.env.REACT_APP_API_URL}/api/memory`, {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    });
    const responseJson = await response.json();
    
    // actualizaciones
    setTotal(responseJson.total_memory);
    setConsumidad(responseJson.consumed_memory);
    setUsada(responseJson.use_memory);
    setdataPie({
      labels: ["Libre", "Usada"],
      datasets: [
        {
          label: "%",
          data: [responseJson.use_memory, 100 - responseJson.use_memory],
          backgroundColor: [
            "rgba(54, 162, 235, 0.2)",
            "rgba(255, 99, 132, 0.2)",
          ],
          borderColor: ["rgba(54, 162, 235, 1)", "rgba(255, 99, 132, 1)"],
          borderWidth: 3,
        },
      ],
    });
  };

  useEffect(() => {
    setTimeout(() => {
      memory();
    }, 5000);
  }, [dataPie]);

  return (
    <>
      <Card style={{ width: "18rem" }}>
        <Pie data={dataPie} />
        <Card.Body>
          <Card.Title>RAM</Card.Title>
          <Card.Text>{`Total: ${total} Consumida: ${consumida} Usada: ${usada}`}</Card.Text>
        </Card.Body>
      </Card>
    </>
  );
};

export default PieChart;

import React from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import PieChart from "./PieChart";

import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import PieChartCpu from "./PieChartCpu";

const CardChart = () => {
  return (
    <>
      <Row className="justify-content-center" >
        <PieChart></PieChart>
        <PieChartCpu></PieChartCpu>
      </Row>
    </>
  );
};

export default CardChart;

import React from "react";
import NavBarMenu from "../components/NavBarMenu";
import PieChart from "../components/PieChart";
import CardChart from "../components/CardChart";
import AcordinInfo from "../components/AcordinInfo";

const Home = () => {
  return (
    <>
      <NavBarMenu />
      <CardChart />
      <AcordinInfo />
    </>
  );
};

export default Home;

import React from "react";
import '../../App.css';
import BarChart from "../BarChart.jsx";
import NavigationBar from "../NavigationBar.jsx";

function Home() {
    return (
        <div>
            <NavigationBar />
            <BarChart />
        </div>
    );
}

export default Home;
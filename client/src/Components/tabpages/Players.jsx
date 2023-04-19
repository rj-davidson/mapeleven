import React from "react";
import '../../App.css';
import NavigationBar from "../NavigationBar.jsx";
import BarChart from "../BarChart.jsx";
import SpiderChart from "../SpiderChart.jsx";
import ScatterPlot from "../ScatterPlot.jsx";

function Players() {
    return (
        <div>
            <NavigationBar />
            <BarChart />
            <ScatterPlot />
            <SpiderChart />
            <h1 className='players'>PLAYERS</h1>;
        </div>
    );
}

export default Players

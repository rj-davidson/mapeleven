import React from "react";
import '../../../App.css';
import NavigationBar from "../../Layout/NavigationBar.tsx";
import BarChart from "../../DataVis/BarChart.jsx";
import SpiderChart from "../../DataVis/SpiderChart.jsx";
import ScatterPlot from "../../DataVis/ScatterPlot.jsx";

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

import React from "react";
import '../../../App.css';
import BarChart from "../../DataVis/BarChart.jsx";
import SpiderChart from "../../DataVis/SpiderChart.jsx";
import ScatterPlot from "../../DataVis/ScatterPlot.jsx";
import Layout from "../../Layout/Layout.tsx";

function Players() {
    return (
        <div>
            <Layout>
                <BarChart />
                <ScatterPlot />
                <SpiderChart />
            </Layout>
        </div>
    );
}

export default Players

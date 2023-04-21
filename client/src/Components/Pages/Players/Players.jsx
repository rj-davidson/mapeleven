import React from "react";
import '../../../App.css';
import BarChart from "../../DataVis/BarChart.jsx";
import SpiderChart from "../../DataVis/SpiderChart.jsx";
import ScatterPlot from "../../DataVis/ScatterPlot.jsx";
import Layout from "../../Layout/Layout.tsx";
import './Players.css'

function Players() {
    return (
        <div>
            <Layout>
                <div className="chartWrapper1">
                    <BarChart />
                    <ScatterPlot />
                </div>
                <SpiderChart />
            </Layout>
        </div>
    );
}

export default Players

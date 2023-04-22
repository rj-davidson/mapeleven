import React from 'react';
import '../../../App.css';
import BarChart from '../../DataVis/BarChart.jsx';
import SpiderChart from '../../DataVis/SpiderChart.jsx';
import ScatterPlot from '../../DataVis/ScatterPlot.jsx';
import Layout from '../../Layout/Layout.tsx';
import TestDisplayPlayers from './TestDisplayPlayers.jsx';
import './Players.css';

function Players() {
    return (
        <div>
            <TestDisplayPlayers />
            <div className='chartWrapper1'>
                <BarChart />
                <ScatterPlot />
            </div>

            <SpiderChart />
        </div>
    );
}

export default Players;

import React from 'react';
import '../../../App.css';
import BarChart from '../../DataVis/BarChart.jsx';
import SpiderChart from '../../DataVis/SpiderChart.jsx';
import ScatterPlot from '../../DataVis/ScatterPlot.jsx';
import TestDisplayPlayers from './TestDisplayPlayers.jsx';
import './Players.css';
import PlayerCards from './PlayerCard.tsx';

function Players() {
    return (
        <>
            <PlayerCards />
            <BarChart />
            <ScatterPlot />
            <SpiderChart />
        </>
    );
}

export default Players;

import React from 'react';
import '../../../App.css';
import BarChart from '../../DataVis/BarChart.jsx';
import SpiderChart from '../../DataVis/SpiderChart.jsx';
import ScatterPlot from '../../DataVis/ScatterPlot.jsx';
import TestDisplayPlayers from './TestDisplayPlayers.jsx';
import './Players.css';
import PlayerCards from './PlayerCard.tsx';
import {Flex} from "../../Util/Flex.jsx";
import {Tile} from "../../Util/Tile.jsx";

function Players() {
    return (
        <Flex style={{width: '100%', gap: '24px'}}>
            <Flex style={{width: '100%', flexDirection: 'column', gap: '12px'}}>
                <Tile>
                    <PlayerCards/>
                </Tile>

                <Tile>
                    <BarChart/>
                </Tile>

                <Tile>
                    <ScatterPlot/>
                </Tile>

                <Tile>
                    <SpiderChart/>
                </Tile>
            </Flex>
        </Flex>
    );
}

export default Players;

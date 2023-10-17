import React from 'react';
import '../App.css';
import BarChart from '../Util/BarChart.jsx';
import SpiderChart from '../Util/SpiderChart.jsx';
import ScatterPlot from '../Util/ScatterPlot.jsx';
import './Players.css';
import PlayerCards from './PlayerCard.tsx';
import { Tile } from '../Util/Tile.jsx';
import { Grid } from '@mui/material';
import DisplayPlayers from './DisplayPlayersCopy.tsx';

function Players() {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <DisplayPlayers skeleton={15} />
                </Tile>
            </Grid>
        </Grid>
    );
}

export default Players;

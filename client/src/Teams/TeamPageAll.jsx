import React from 'react';
import '../App.css';
import { Grid } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import ScatterPlot from '../Util/ScatterPlot.jsx';
import SpiderChart from '../Util/SpiderChart.jsx';
import ListTeams from './ListTeams.tsx';

function TeamPageAll() {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <ListTeams limit={30} skeleton={5} />
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <ScatterPlot />
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <SpiderChart />
                </Tile>
            </Grid>
        </Grid>
    );
}

export default TeamPageAll;

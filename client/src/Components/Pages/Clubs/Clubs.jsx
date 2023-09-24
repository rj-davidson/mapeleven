import React from 'react';
import '../../../App.css';
import DisplayClubs from './DisplayClubs.jsx';
import ClubCard from './ClubCard.tsx';
import { Grid } from '@mui/material';
import { Tile } from '../../Util/Tile.jsx';
import ScatterPlot from '../../DataVis/ScatterPlot.jsx';
import SpiderChart from '../../DataVis/SpiderChart.jsx';
import DisplayTeams from './DisplayTeams.tsx';

function Clubs() {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <DisplayTeams limit={30} skeleton={5} />
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

export default Clubs;

import React from 'react';
import '../../../App.css';
import BarChart from '../../DataVis/BarChart.jsx';
import SpiderChart from '../../DataVis/SpiderChart.jsx';
import ScatterPlot from '../../DataVis/ScatterPlot.jsx';
import './Players.css';
import PlayerCards from './PlayerCard.tsx';
import {Tile} from "../../Util/Tile.jsx";
import { Grid } from "@mui/material";

function Players() {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <PlayerCards/>
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <BarChart/>
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <ScatterPlot/>
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <SpiderChart/>
                </Tile>
            </Grid>
        </Grid>
    );
}

export default Players;

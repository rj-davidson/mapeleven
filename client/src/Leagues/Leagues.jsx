import React from 'react';
import '../App.css';
import DisplayLeagues from './DisplayLeagues.jsx';
import Layout from '../Layout/Layout.tsx';
import LeagueCard from './LeagueCard.tsx';
import {Grid} from "@mui/material";
import {Tile} from "../Util/Tile.jsx";
//import PlayerCards from "../Players/PlayerCard.js";
//import BarChart from "../../DataVis/BarChart.jsx";
import ScatterPlot from "../Util/ScatterPlot.jsx";
import SpiderChart from "../Util/SpiderChart.jsx";

function Leagues() {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <LeagueCard/>
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <DisplayLeagues/>
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
//<div>
//    <LeagueCard />
//</div>
export default Leagues;

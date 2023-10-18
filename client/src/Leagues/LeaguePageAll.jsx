import React from 'react';
import '../App.css';
import { Grid } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import ListLeagues from './ListLeagues.tsx';

function LeaguePageAll() {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile>
                    <ListLeagues skeleton={5} />
                </Tile>
            </Grid>
        </Grid>
    );
}

export default LeaguePageAll;

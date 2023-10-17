import * as React from 'react';
import { useState } from 'react';
import { Grid } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import TeamIDStatBox from './TeamIDStatBox';

interface TeamStatsProps {
    goals: number;
    clean: number;
    gamesPlayed: number;
    wins: number;
    draws: number;
    loses: number;
    averageGoals: number;
    gamesScored: number;
    failedToScore: number;
}

const TeamStats: React.FC<TeamStatsProps> = ({
    goals,
    clean,
    gamesPlayed,
    wins,
    draws,
    loses,
    averageGoals,
    gamesScored,
    failedToScore,
}) => {
    return (
        <Tile
            style={{
                flexDirection: 'column',
                gap: '24px',
            }}
        >
            <Grid container spacing={1}>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Wins' value={wins} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Draws' value={draws} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Loses' value={loses} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Goals' value={goals} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Average Goals' value={averageGoals} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Matches Scored' value={gamesScored} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Failed To Score' value={failedToScore} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                    <TeamIDStatBox stat='Clean Sheets' value={clean} />
                </Grid>
            </Grid>
        </Tile>
    );
};

export default TeamStats;

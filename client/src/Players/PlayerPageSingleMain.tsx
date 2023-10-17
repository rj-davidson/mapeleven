import * as React from 'react';
import { Grid } from '@mui/material';
import PlayerRadarChart from './PlayerRadarChart';

interface PlayerPageSingleMainProps {
    name: string;
}

const PlayerPageSingleMain: React.FC<PlayerPageSingleMainProps> = ({ name }) => {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <PlayerRadarChart name={name} />
            </Grid>
        </Grid>
    );
};

export default PlayerPageSingleMain;

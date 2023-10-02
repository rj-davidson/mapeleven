import * as React from 'react';
import { Grid } from '@mui/material';
import { APITSGoalMinuteSplit } from '../Models/api_types';
import TeamGoalSplitCharts from './TeamGoalSplitCharts';

interface TeamPageSingleMainProps {
    goalMinuteSplit: APITSGoalMinuteSplit;
}

const TeamPageSingleMain: React.FC<TeamPageSingleMainProps> = ({ goalMinuteSplit }) => {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <TeamGoalSplitCharts goalMinuteSplit={goalMinuteSplit} />
            </Grid>
        </Grid>
    );
};

export default TeamPageSingleMain;

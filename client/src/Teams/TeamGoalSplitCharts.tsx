import * as React from 'react';
import { Box, Grid, Typography } from '@mui/material';
import { Tile } from '../Util/TileTS';
import { Flex } from '../Util/Flex.jsx';
import { APITSGoalMinuteSplit } from '../Models/api_types';
import GoalSplitBarChart from './GoalSplitBarChart';
import GoalSplitPieChart from './GoalSplitPieChart';

interface TeamGoalSplitChartsProps {
    goalMinuteSplit: APITSGoalMinuteSplit;
}

const TeamGoalSplitCharts: React.FC<TeamGoalSplitChartsProps> = ({ goalMinuteSplit }) => {
    return (
        <Tile sx={{ flexDirection: 'column' }}>
            <Typography sx={{ fontSize: { xs: '20px', sm: '20px' }, textAlign: 'left' }}>
                Goals by the Minute
            </Typography>
            <Grid container spacing={1}>
                <Grid item xs={12} sm={12} md={12} lg={3} width='100%'>
                    <Flex
                        style={{
                            width: '100%',
                            justifyContent: 'center',
                            alignItems: 'center',
                            flexDirection: 'column',
                        }}
                    >
                        <Box sx={{ width: '100%', height: '200px' }}>
                            <GoalSplitPieChart data={goalMinuteSplit} />
                        </Box>
                    </Flex>
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={9} width='100%'>
                    <Flex
                        style={{
                            width: '100%',
                            justifyContent: 'center',
                            alignItems: 'center',
                            flexDirection: 'column',
                        }}
                    >
                        <Box
                            sx={{
                                width: '100%',
                                height: { xs: '100px', sm: '100px', md: '300px', lg: '200px' },
                            }}
                        >
                            <GoalSplitBarChart data={goalMinuteSplit} />
                        </Box>
                    </Flex>
                </Grid>
            </Grid>
        </Tile>
    );
};

export default TeamGoalSplitCharts;

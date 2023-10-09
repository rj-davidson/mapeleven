import * as React from 'react';
import { Box, Typography } from '@mui/material';

interface ListFixtureScoreProps {
    homeScore: number;
    awayScore: number;
    status: string;
}

const ListFixtureScore: React.FC<ListFixtureScoreProps> = ({ homeScore, awayScore, status }) => {
    return (
        <Box
            sx={{
                width: '20%',
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
            }}
        >
            {status === 'Not Started' ? (
                <Typography variant='h6' component='h2' sx={{ fontSize: '16px', textAlign: 'center' }}>
                    -
                </Typography>
            ) : (
                <Typography variant='h6' component='h2' sx={{ fontSize: '16px' }}>
                    {homeScore} - {awayScore}
                </Typography>
            )}
        </Box>
    );
};

export default ListFixtureScore;

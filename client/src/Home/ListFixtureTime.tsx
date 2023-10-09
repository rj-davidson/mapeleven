import * as React from 'react';
import { Box, Typography } from '@mui/material';

interface ListFixtureTimeProps {
    status: string;
    elapsed: number;
    timezone: string;
}

const ListFixtureTime: React.FC<ListFixtureTimeProps> = ({ status, elapsed, timezone }) => {
    let content = '';

    if (status === 'Match Finished') {
        content = 'FT';
    } else if (elapsed) {
        content = `${elapsed}'`;
    }

    return (
        <Box sx={{ width: '5%' }}>
            <Typography
                variant='h6'
                component='h2'
                sx={{ fontSize: '16px', color: status !== 'Match Finished' ? 'var(--green)' : 'inherit' }}
            >
                {content}
            </Typography>
        </Box>
    );
};

export default ListFixtureTime;

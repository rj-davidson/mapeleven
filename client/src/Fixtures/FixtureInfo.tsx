import * as React from 'react';
import { Grid, Typography } from '@mui/material';

type FixtureInfoProps = {
    live: boolean;
    time: string;
    kickoff: string;
    location: string;
};

const FixtureInfo: React.FC<FixtureInfoProps> = ({ live, time, kickoff, location }) => {
    return (
        <Grid container direction='column' sx={{ height: '100%' }}>
            {live ? (
                <Grid item container sx={{ height: '100%' }} alignItems='center' justifyContent='center'>
                    <Typography variant='body1' color='var(--green)' fontWeight='700'>
                        {time}
                    </Typography>
                </Grid>
            ) : (
                <Grid
                    container
                    direction='column'
                    spacing={3}
                    sx={{ height: '100%' }}
                    alignItems='center'
                    justifyContent='center'
                >
                    <Grid item>
                        <Typography variant='body2' color='var(--light1)'>
                            {kickoff}
                        </Typography>
                    </Grid>
                    <Grid item>
                        <Typography variant='body2' color='var(--light1)'>
                            {location}
                        </Typography>
                    </Grid>
                </Grid>
            )}
        </Grid>
    );
};

export default FixtureInfo;

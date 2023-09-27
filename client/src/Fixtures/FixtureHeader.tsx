import * as React from 'react';
import { Grid, Typography } from '@mui/material';
import Football from '@mui/icons-material/SportsSoccer';
import Trophy from '@mui/icons-material/EmojiEvents';

type FixtureHeaderProps = {
    flag: string;
    league: string;
};

const FixtureHeader: React.FC<FixtureHeaderProps> = ({ flag, league }) => {
    return (
        <Grid container direction='row' sx={{ width: '100%' }}>
            <Grid
                container
                direction='row'
                spacing={2}
                sx={{ height: '100%' }}
                alignItems='center'
                justifyContent='center'
            >
                <Grid item>
                    <img src={flag} alt={league + ' flag'} height='12px' />
                </Grid>
                <Grid item>
                    <Typography variant='body2' color='var(--light1)'>
                        {league}
                    </Typography>
                </Grid>
            </Grid>
        </Grid>
    );
};

export default FixtureHeader;

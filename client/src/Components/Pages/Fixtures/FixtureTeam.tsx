import * as React from 'react';
import { Box, Grid, Typography } from '@mui/material';

type FixtureTeamProps = {
    logo: string;
    name: string;
    score?: number;
};

const styles = {
    teamWrapper: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        height: 44,
        color: 'black',
    },
    teamName: {
        flexGrow: 1,
        paddingLeft: 10,
        color: 'var(--light1)',
    },
    teamScore: {
        fontSize: 20,
        fontWeight: 'bold',
        color: 'var(--light1)',
    },
};

const FixtureTeam: React.FC<FixtureTeamProps> = ({ logo, name, score }) => {
    if (typeof logo !== 'string' || typeof name !== 'string' || (typeof score !== 'number' && score !== undefined)) {
        return null; // or handle the error in some other way
    }

    return (
            <Box sx={{ width: '100%' }} id="teamBox">
                <Grid container justifyContent="space-between" alignItems="center">
                    <Grid item>
                        <Grid container alignItems="center" spacing={3}>
                            <Grid item>
                                <img src={logo} alt={`${name} logo`} width="44" height="44" />
                            </Grid>
                            <Grid item>
                                <Typography variant="body2">{name}</Typography>
                            </Grid>
                        </Grid>
                    </Grid>
                    <Grid item>
                        <Grid container alignItems="center">
                            <Typography variant="body1" sx={styles.teamScore}>
                                {score !== undefined && score !== null ? score : ''}
                            </Typography>
                        </Grid>
                    </Grid>
                </Grid>
            </Box>
    );
};

export default FixtureTeam;

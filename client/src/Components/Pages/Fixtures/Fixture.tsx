import * as React from 'react';
import { Box, Card, CardContent, Grid, Typography } from '@mui/material';
import FixtureHeader from './FixtureHeader';

type FixtureProps = {
    homeTeam: React.ReactNode;
    awayTeam: React.ReactNode;
    fixtureInfo: React.ReactNode;
    fixtureHeader: React.ReactNode;
};

const styles = {
    fixtureCard: {
        width: '100%',
        backgroundColor: 'var(--dark1)',
        border: '3px solid var(--dark1)',
        margin: '0',
        padding: '0',
        '&:hover': {
            border: '3px solid var(--focus)',
            backgroundColor: 'var(--dark2)',
        },
    },
    teamName: {
        color: 'var(--light1)',
        fontSize: 24,
    },
    teamScore: {
        color: 'var(--light1)',
        fontSize: 48,
        fontWeight: 'bold',
    },
};

const Fixture: React.FC<FixtureProps> = ({ homeTeam, awayTeam, fixtureInfo, fixtureHeader }) => {
    return (
        <Card sx={styles.fixtureCard}>
            <CardContent>
                <Grid container display='flex' spacing={1}>
                    <Grid item xs={12}>
                        {fixtureHeader}
                    </Grid>
                    <Grid item xs={9}>
                        <Box>
                            <Typography marginBottom={2} sx={styles.teamName}>
                                {awayTeam}
                            </Typography>
                            <Typography sx={styles.teamName}>{homeTeam}</Typography>
                        </Box>
                    </Grid>
                    <Grid item xs={3}>
                        {fixtureInfo}
                    </Grid>
                </Grid>
            </CardContent>
        </Card>
    );
};

export default Fixture;

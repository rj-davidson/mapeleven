import * as React from 'react';
import { Box, Card, CardContent, Grid, Typography } from '@mui/material';

type FixtureProps = {
    homeTeam: React.ReactNode;
    awayTeam: React.ReactNode;
    fixtureInfo: React.ReactNode;
};

const styles = {
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

const Fixture: React.FC<FixtureProps> = ({ homeTeam, awayTeam, fixtureInfo }) => {
    return (
        <Grid container sx={{ width: '100%', margin: '16px' }}>
            <Grid item xs={12}>
                <Card sx={{ width: '100%' }}>
                    <CardContent>
                        <Grid container id={"fixtureRow"}>
                            <Grid item xs={9} sm={9} md={9} lg={9}>
                                <Box>
                                    <Typography marginBottom={2} sx={styles.teamName}>
                                        {awayTeam}
                                    </Typography>
                                    <Typography sx={styles.teamName}>
                                        {homeTeam}
                                    </Typography>
                                </Box>
                            </Grid>
                            <Grid item xs={3} sm={3} md={3} lg={3}>
                                {fixtureInfo}
                            </Grid>
                        </Grid>
                    </CardContent>
                </Card>
            </Grid>
        </Grid>
    );
};

export default Fixture;

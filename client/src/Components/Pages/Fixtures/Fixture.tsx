import * as React from 'react';
import { Box, Card, CardContent, Grid, Typography } from '@mui/material';

type FixtureProps = {
    homeTeam: React.ReactNode;
    awayTeam: React.ReactNode;
    fixtureInfo: React.ReactNode;
};

const styles = {
    root: {
        minWidth: 400,
        width: '100%',
        margin: '16px',
        backgroundColor: 'var(--dark2)',
        color: 'var(--light1)',
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

const Fixture: React.FC<FixtureProps> = ({ homeTeam, awayTeam, fixtureInfo }) => {
    return (
        <Card sx={styles.root}>
            <CardContent>
                <Grid item xs={8} sm={8} md={8} lg={8}>
                    <Box display='flex' flexDirection='row'>
                        <Box flexDirection='column'>
                            <Box marginBottom={2} sx={styles.teamName}>
                                {homeTeam}
                            </Box>
                            <Box sx={styles.teamName}>{awayTeam}</Box>
                        </Box>
                    </Box>
                </Grid>

                <Box paddingTop={12}>
                    <Typography>{fixtureInfo}</Typography>
                </Box>
            </CardContent>
        </Card>
    );
};

export default Fixture;

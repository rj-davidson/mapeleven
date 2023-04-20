import * as React from 'react';
import FixtureTeam from './FixtureTeam';
import { Box, Card, CardContent, Grid, Typography } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

type FixtureProps = {
    homeTeam: React.ReactNode;
    awayTeam: React.ReactNode;
    fixtureInfo: React.ReactNode;
};

const useStyles = makeStyles({
    root: {
        minWidth: 400,
        width: '100%',
        margin: '16px',
    },
    teamName: {
        fontSize: 24,
    },
    teamScore: {
        fontSize: 48,
        fontWeight: 'bold',
    },
});

const Fixture: React.FC<FixtureProps> = ({ homeTeam, awayTeam, fixtureInfo }) => {
    const classes = useStyles();

    return (
        <Card className={classes.root}>
            <CardContent>
                <Grid item xs={12} sm={6} md={4}>
                    <Box display='flex' flexDirection='row'>
                        <Box flexDirection='column'>
                            <Box marginBottom={2}>{homeTeam}</Box>
                            <Box>{awayTeam}</Box>
                        </Box>
                    </Box>
                </Grid>

                <Box marginTop={2}>
                    <Typography>{fixtureInfo}</Typography>
                </Box>
            </CardContent>
        </Card>
    );
};

export default Fixture;

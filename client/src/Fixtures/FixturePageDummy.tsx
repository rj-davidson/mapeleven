import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import {Box, Card, CardContent, CardMedia, Grid, Paper, Skeleton, Typography} from '@mui/material';
import Field from '../Teams/FieldSVG'

const url = import.meta.env.VITE_API_URL;

function FixturePageDummy() {
    return(
        <Grid container spacing={3} justifyContent="center">
            <Grid item xs={12}>
                <Card>
                    <CardContent>
                        <Grid container spacing={2} justifyContent="center">
                            <Grid item xs={4}>
                                <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                                    <CardMedia
                                        component="img"
                                        alt="Chelsea"
                                        style={{ width: '50px' }}
                                        image="http://64.176.215.29:8080/api/images/teams/49.png"
                                    />
                                    <Typography variant="h6" align="center" sx={{ fontSize: '42px' }}>
                                        Chelsea
                                    </Typography>
                                </div>

                                <Typography variant="body2" align="center">
                                    Didier Drogba
                                </Typography>
                                <Typography variant="body2" align="center">
                                    Eden Hazard
                                </Typography>
                            </Grid>
                            <Grid item xs={4}>
                                <Typography variant="h4" align="center">
                                    2 - 3
                                </Typography>
                                <Typography variant="h6" align="center">
                                    67'
                                </Typography>
                            </Grid>
                            <Grid item xs={4}>
                                <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                                    <CardMedia
                                        component="img"
                                        alt="Arsenal"
                                        style={{ width: '50px' }}
                                        image="http://64.176.215.29:8080/api/images/teams/42.png"
                                    />
                                    <Typography variant="h6" align="center" sx={{ fontSize: '42px' }}>
                                        Arsenal
                                    </Typography>
                                </div>

                                <Typography variant="body2" align="center">
                                    Robin van Persie
                                </Typography>
                                <Typography variant="body2" align="center">
                                    Robin van Persie
                                </Typography>
                                <Typography variant="body2" align="center">
                                    Robin van Persie
                                </Typography>
                            </Grid>
                        </Grid>
                        <Box mt={2}>
                            <Typography variant="body1" align="center">
                                Referee: Ethan
                            </Typography>
                        </Box>
                    </CardContent>
                </Card>
                <Card>

                <Grid item xs={6}>
                    <Card>

                    </Card>
                    <Grid container spacing={2} justifyContent="center">
                        <Grid item xs={12}>
                            <Typography variant="h6" align="center">
                                Away Formation
                            </Typography>
                        </Grid>
                        <Grid item xs={12}>
                            <Field /> {/* Away team field */}
                        </Grid>
                    </Grid>
                    <Grid container spacing={2} justifyContent="center">
                        <Grid item xs={12}>
                            <Typography variant="h6" align="center">
                                Home Formation
                            </Typography>
                        </Grid>
                        <Grid item xs={12}>
                            <Field /> {/* Home team field */}
                        </Grid>
                    </Grid>
                </Grid>
                </Card>
            </Grid>
        </Grid>
    );
}

export default FixturePageDummy;

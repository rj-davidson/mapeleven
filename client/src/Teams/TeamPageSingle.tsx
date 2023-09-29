import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import { Box, Grid, Skeleton, Typography } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import { Flex } from '../Util/Flex.jsx';
import { APITSGoalMinuteSplit } from '../Models/api_types';
import GoalSplitBarChart from './GoalSplitBarChart';
import GoalSplitPieChart from './GoalSplitPieChart';
import FormCard from './FormCard.jsx';
import { Groups } from '@mui/icons-material';

function TeamPageSingle() {
    const navigate = useNavigate();
    const { slug } = useParams();

    const [name, setName] = useState<string>('');
    const [logo, setLogo] = useState<string>('');
    const [flag, setFlag] = useState<string>('');
    const [country, setCountry] = useState<string>('');
    const [goalMinuteSplit, setGoalMinuteSplit] = useState<APITSGoalMinuteSplit>(null);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        // Send a GET request to the API.
        fetch(`http://localhost:8080/teams/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                setName(jsonData.name.long);
                setLogo(jsonData.badge);
                setFlag(jsonData.country.flag);
                setCountry(jsonData.country.name);
                setGoalMinuteSplit(jsonData.stats.goals.for.minute);
                setLoading(false);
            })
            .catch(error => console.error(error));
    }, [slug]);

    return loading ? (
        <Grid container spacing={2} direction='column'>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Skeleton
                    variant='rectangular'
                    height='150px'
                    width='100%'
                    sx={{ background: 'gray', borderRadius: '12px' }}
                />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Skeleton
                    variant='rectangular'
                    height='600px'
                    width='100%'
                    sx={{ background: 'gray', borderRadius: '12px' }}
                />
            </Grid>
        </Grid>
    ) : (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile style={{ justifyContent: 'space-between' }}>
                    <Flex style={{ gap: '32px', alignItems: 'center' }}>
                        <Box
                            component='img'
                            sx={{
                                maxHeight: { xs: 64, sm: 96 },
                                maxWidth: { xs: 64, sm: 96 },
                            }}
                            alt={name}
                            src={'http://localhost:8080/' + logo}
                        />
                        <Flex style={{ flexDirection: 'column' }}>
                            <Typography sx={{ fontSize: { xs: '24px', sm: '32px' } }}>{name}</Typography>
                            <Flex
                                style={{ justifyContent: 'left', alignItems: 'center', gap: '8px', marginLeft: '2px' }}
                            >
                                <Box
                                    component='img'
                                    sx={{
                                        maxHeight: 20,
                                        maxWidth: 25,
                                    }}
                                    alt={country}
                                    src={'http://localhost:8080/' + flag}
                                />
                                <Typography sx={{ fontSize: '18px' }}>{country}</Typography>
                            </Flex>
                        </Flex>
                    </Flex>
                    <Flex style={{ justifyContent: 'center', alignItems: 'center'}}>
                        <Groups style={{ fontSize: 50 }} />
                    </Flex>
                </Tile>
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={2.5} width='100%'>
                <Flex style={{ justifyContent: 'center', alignItems: 'center', flexDirection: 'column' }}>
                    <FormCard />
                </Flex>
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={9.5} width='100%'>
                <Tile style={{ flexDirection: 'column' }}>
                    <Typography sx={{ fontSize: { xs: '20px', sm: '20px' }, textAlign: 'left' }}>
                        Goals by the Minute
                    </Typography>
                    <Grid container spacing={1}>
                        <Grid item xs={12} sm={12} md={12} lg={3} width='100%'>
                            <Flex
                                style={{
                                    width: '100%',
                                    justifyContent: 'center',
                                    alignItems: 'center',
                                    flexDirection: 'column',
                                }}
                            >
                                <Box sx={{ width: '100%', height: '200px' }}>
                                    <GoalSplitPieChart data={goalMinuteSplit} />
                                </Box>
                            </Flex>
                        </Grid>
                        <Grid item xs={12} sm={12} md={12} lg={9} width='100%'>
                            <Flex
                                style={{
                                    width: '100%',
                                    justifyContent: 'center',
                                    alignItems: 'center',
                                    flexDirection: 'column',
                                }}
                            >
                                <Box
                                    sx={{
                                        width: '100%',
                                        height: { xs: '100px', sm: '100px', md: '300px', lg: '200px' },
                                    }}
                                >
                                    <GoalSplitBarChart data={goalMinuteSplit} />
                                </Box>
                            </Flex>
                        </Grid>
                    </Grid>
                </Tile>
            </Grid>
        </Grid>
    );
}

export default TeamPageSingle;

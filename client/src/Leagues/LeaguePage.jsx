import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Box, Grid, Skeleton, Typography, useMediaQuery } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import { Flex } from '../Util/Flex.jsx';
import Standings from './Standings.jsx';

const url = import.meta.env.VITE_API_URL;

function LeaguePage() {
    const { slug } = useParams();

    // Declare state variables using the `useState` hook.
    const [name, setName] = useState();
    const [logo, setLogo] = useState();
    const [flag, setFlag] = useState('');
    const [country, setCountry] = useState('');
    const [teams, setTeams] = useState(new Set());
    const [loading, setLoading] = useState(true);

    const isSmallerThanLG = useMediaQuery('(max-width: 1260px)');

    useEffect(() => {
        setLoading(true);

        // Send a GET request to the API.
        fetch(`${url}/leagues/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                setName(jsonData.name);
                setLogo(jsonData.logo);
                setFlag(jsonData.country.flag);
                setCountry(jsonData.country.name);

                const standings = jsonData.standings;
                const newTeams = new Set();
                for (let i = 0; i < standings.length; i++) {
                    if (standings[i].rank > 0) {
                        newTeams.add({
                            rank: standings[i].rank,
                            name: standings[i].team.name.long,
                            logo: standings[i].team.badge,
                            points: standings[i].points,
                            played: standings[i].overall.played,
                            wins: standings[i].overall.won,
                            draw: standings[i].overall.draw,
                            losses: standings[i].overall.lost,
                            slug: standings[i].team.slug,
                            form: standings[i].form,
                        });
                    }
                }
                // Update the state variables with the new data.
                setTeams(newTeams);
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
                <Tile style={{ justifyContent: 'left', gap: isSmallerThanLG ? '12px' : '18px' }}>
                    <Flex style={{ alignItems: 'center' }}>
                        <Box
                            component='img'
                            sx={{
                                maxHeight: { xs: 60, sm: 96 },
                                maxWidth: { xs: 60, sm: 96 },
                            }}
                            alt={name}
                            src={`${url}/` + logo}
                        />
                    </Flex>
                    <Flex style={{ alignItems: 'left', flexDirection: 'column' }}>
                        <Typography sx={{ fontSize: { xs: '24px', sm: '48px' } }}>{name}</Typography>
                        <Flex
                            style={{
                                justifyContent: 'left',
                                alignItems: 'center',
                                flexDirection: 'row',
                                gap: '6px',
                            }}
                        >
                            {country !== 'World' && (
                                <Box
                                    component='img'
                                    sx={{
                                        maxHeight: { xs: 16, sm: 24 },
                                        maxWidth: { xs: 20, sm: 30 },
                                    }}
                                    alt={country}
                                    src={`${url}/` + flag}
                                />
                            )}
                            <Typography sx={{ fontSize: { xs: '14px', sm: '24px' } }}>{country}</Typography>
                        </Flex>
                    </Flex>
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile style={{ justifyContent: 'left' }}>
                    <Standings data={teams} />
                </Tile>
            </Grid>
        </Grid>
    );
}

export default LeaguePage;

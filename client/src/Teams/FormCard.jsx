import React, { useState, useRef, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import * as d3 from 'd3';
import { Box, Grid, Skeleton, Typography } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import { useNavigate } from 'react-router-dom';
import { Flex } from '../Util/Flex.jsx';
import Standings from '../Leagues/Standings.jsx';

const url = import.meta.env.VITE_API_URL;

function FormCard() {
    const navigate = useNavigate();
    const { slug } = useParams();

    // Declare state variables using the `useState` hook.
    const [form, setForm] = useState();
    const [stats, setStats] = useState();
    const [loading, setLoading] = useState(true);

    const characterColors = {
        W: 'var(--green)',
        L: 'var(--red)',
    };

    useEffect(() => {
        setLoading(true);

        // Send a GET request to the API.
        fetch(`${url}/teams/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                setForm(jsonData.stats.form);
                setLoading(false);
            })
            .catch(error => console.error(error));
    }, [slug]);

    return loading ? (
        <Grid container direction='column'>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Skeleton
                    variant='rectangular'
                    height='50px'
                    width='100%'
                    sx={{ background: 'gray', borderRadius: '12px' }}
                />
            </Grid>
        </Grid>
    ) : (
        <Tile style={{ justifyContent: 'left' }}>
            <Flex style={{ justifyContent: 'left', alignItems: 'left', gap: '8px', flexDirection: 'column' }}>
                <Typography sx={{fontSize: { xs: '20px', sm: '20px' }}}>
                    Form
                </Typography>
                <Flex style={{ justifyContent: 'left', alignItems: 'left', gap: '15px', flexDirection: 'row' }}>
                    {form
                        .slice(0, 5)
                        .split('')
                        .map((character, index) => (
                            <Box
                                sx={{
                                    display: 'flex',
                                    justifyContent: 'center',
                                    alignItems: 'center',
                                    background: characterColors[character] ?? 'gray',
                                    width: '32px',
                                    borderRadius: '6px'
                                }}
                            >
                                <Typography key={index} variant='h6' component='span' style={{ color: 'white' }}>
                                    {character}
                                </Typography>
                            </Box>
                        ))}
                </Flex>
            </Flex>
        </Tile>
    );
}

export default FormCard;

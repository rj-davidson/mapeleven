import { useEffect, useState } from 'react';
import { Card, CardContent, CardMedia, Grid, Typography } from '@mui/material';
import * as React from 'react';

function PlayerCards() {
    const [players, setPlayers] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/stats/topscorers')
            .then(response => response.json())
            .then(jsonData => {
                const response = jsonData.response;
                const newPlayers = [];
                for (let i = 0; i < response.length; i++) {
                    const player = response[i].player;
                    newPlayers.push(player);
                }
                setPlayers(newPlayers);
            })
            .catch(error => console.error(error));
    }, []);

    const playerCards = players.map(player => {
        return (
            <Grid item xs={12} sm={12} md={12} lg={12} key={player.id} width='100%'>
                <Card sx={{ height: '100%', width: '100%' }}>
                    <CardContent
                        sx={{
                            display: 'flex',
                            flexDirection: 'row',
                            flexWrap: 'wrap',
                            width: '100%',
                            alignItems: 'center',
                            justifyContent: 'left',
                            paddingLeft: '12px',
                            paddingRight: '32px',
                            paddingTop: '24px',
                            paddingBottom: '24px',
                        }}
                    >
                        <div
                            style={{
                                display: 'flex',
                                flexDirection: 'row',
                                marginLeft: '16px',
                                alignItems: 'center',
                                textAlign: 'left',
                                justifyContent: 'space-between',
                                width: '100%',
                            }}
                        >
                            <img
                                src={player.photo}
                                alt={player.name + ' logo'}
                                height={32}
                                style={{ borderRadius: '8px', marginLeft:'-20px', marginRight:'10px', border: '2px solid #D8DEE9' }}
                            />
                            <Typography variant='body1' color='text.primary'>
                                {player.name}
                            </Typography>

                        </div>
                    </CardContent>
                </Card>
            </Grid>
        );
    });

    return (
        <div>
            <Typography variant='h4' component='h2' gutterBottom>
                Players
            </Typography>
            <Grid container spacing={2} sx={{ justifyContent: 'space-between' }}>
                {playerCards}
            </Grid>
        </div>
    );
}

export default PlayerCards;

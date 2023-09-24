import * as React from 'react';
import { useEffect, useState } from 'react';
import { APITeam } from '../../../Models/api_types';
import { Box, Grid, Skeleton, Typography } from '@mui/material';
import { Flex } from '../../Util/Flex.jsx';
import { Groups, Person } from '@mui/icons-material';
import PlayerCard from './PlayerCardCopy';

interface DisplayPlayersProps {
    limit?: number;
    skeleton: number;
}

const DisplayPlayers: React.FC<DisplayPlayersProps> = ({ limit, skeleton }) => {
    const [playerData, setPlayerData] = useState([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        // Send a GET request to the API.
        fetch(`http://localhost:8080/stats/topscorers`)
            .then(response => response.json())
            .then(jsonData => {
                const response = jsonData.response;
                const newPlayerData = [];
                for (let i = 0; i < response.length; i++) {
                    const player = response[i].player;
                    newPlayerData.push(player);
                }
                if (limit) {
                    setPlayerData(newPlayerData.slice(0, limit));
                } else {
                    setPlayerData(newPlayerData);
                }
                setLoading(false);
            })
            .catch(error => console.error(error));
    });

    const skeletonCards: JSX.Element[] = [];

    for (let i = 0; i < skeleton; i++) {
        skeletonCards.push(
            <Grid item xs={12} sm={12} md={12} lg={12} key={i}>
                <Skeleton
                    variant='rectangular'
                    height='52px'
                    width='100%'
                    sx={{ background: 'var(--dark1)', borderRadius: '12px' }}
                />
            </Grid>,
        );
    }

    return (
        <Box sx={{ width: '100%' }}>
            <Flex
                style={{
                    flexDirection: 'row',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    margin: '0 4px 16px 2px',
                }}
            >
                <Typography variant='h6' component='h2' gutterBottom>
                    Players
                </Typography>
                <Person />
            </Flex>
            {loading ? (
                <Grid container spacing={1}>
                    {skeletonCards}
                </Grid>
            ) : (
                <Grid container spacing={1}>
                    {playerData.map(player => (
                        <Grid item xs={12} sm={12} md={12} lg={12} key={player.slug}>
                            <PlayerCard name={player.name} photo={player.photo} />
                        </Grid>
                    ))}
                </Grid>
            )}
        </Box>
    );
};

export default DisplayPlayers;

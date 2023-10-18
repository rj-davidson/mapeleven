import * as React from 'react';
import { useEffect, useState } from 'react';
import { APITeam } from '../Models/api_types';
import { Box, Grid, Skeleton, Typography } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import { Groups } from '@mui/icons-material';
import LeagueCard from './LeagueCard';

const url = import.meta.env.VITE_API_URL;

interface ListLeaguesProps {
    limit?: number;
    skeleton: number;
}

const ListLeagues: React.FC<ListLeaguesProps> = ({ limit, skeleton }) => {
    const [leagueData, setLeagueData] = useState<APITeam[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        // Send a GET request to the API.
        fetch(`${url}/leagues`)
            .then(response => response.json())
            .then(jsonData => {
                const newTeamData = jsonData;
                if (limit) {
                    setLeagueData(newTeamData.slice(0, limit));
                } else {
                    setLeagueData(newTeamData);
                }
                setLoading(false);
                console.log(leagueData)
            })
            .catch(error => console.error(error));
    }, []);

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

    // @ts-ignore
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
                    Leagues
                </Typography>
                <Groups />
            </Flex>
            {loading ? (
                <Grid container spacing={1}>
                    {skeletonCards}
                </Grid>
            ) : (
                <Grid container spacing={1}>
                    {leagueData.map(league => (
                        <Grid item xs={12} sm={12} md={12} lg={12} key={league.slug}>
                            <LeagueCard
                                slug={league.slug}
                                leagueName={league.name}
                                logo={league.logo}
                                countryCode={league.country.code}
                                countryName={league.country.name}
                                countryFlag={league.country.flag}
                            />
                        </Grid>
                    ))}
                </Grid>
            )}
        </Box>
    );
};

export default ListLeagues;

import * as React from 'react';
import { useEffect, useState } from 'react';
import { APITeam } from '../Models/api_types';
import { Box, Grid, Skeleton, Typography } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import { Groups } from '@mui/icons-material';
import TeamCard from './TeamCard';

const url = import.meta.env.VITE_API_URL;

interface ListTeamsProps {
    limit?: number;
    skeleton: number;
}

const ListTeams: React.FC<ListTeamsProps> = ({ limit, skeleton }) => {
    const [teamData, setTeamData] = useState<APITeam[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        // Send a GET request to the API.
        fetch(`${url}/teams`)
            .then(response => response.json())
            .then(jsonData => {
                const newTeamData = jsonData;
                if (limit) {
                    setTeamData(newTeamData.slice(0, limit));
                } else {
                    setTeamData(newTeamData);
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
                    Teams
                </Typography>
                <Groups />
            </Flex>
            {loading ? (
                <Grid container spacing={1}>
                    {skeletonCards}
                </Grid>
            ) : (
                <Grid container spacing={1}>
                    {teamData.map(team => (
                        <Grid item xs={12} sm={12} md={12} lg={12} key={team.slug}>
                            <TeamCard
                                slug={team.slug}
                                nameLong={team.name.long}
                                nameShort={team.name.short}
                                badge={team.badge}
                                countryCode={team.country.code}
                                countryName={team.country.name}
                                countryFlag={team.country.flag}
                            />
                        </Grid>
                    ))}
                </Grid>
            )}
        </Box>
    );
};

export default ListTeams;

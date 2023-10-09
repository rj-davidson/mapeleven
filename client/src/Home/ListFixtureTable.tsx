import * as React from 'react';
import { APIScoreboardDate } from '../Models/api_types';
import { Grid, Typography } from '@mui/material';
import ListFixtureCard from './ListFixtureCard';
import { Flex } from '../Util/Flex';
import DisplayImage from '../Util/DisplayImage';
import { Link } from 'react-router-dom';

interface ListFixtureTableProps {
    date: APIScoreboardDate;
}

const ListFixtureTable: React.FC<ListFixtureTableProps> = ({ date }) => {
    // Sort the leagues array by name in ascending order
    const sortedLeagues = [...date.leagues].sort((a, b) => b.name.localeCompare(a.name));

    return (
        <Grid container spacing={1}>
            {sortedLeagues.map((league, leagueIndex) => (
                <Grid item xs={12} sm={12} md={12} lg={12} key={leagueIndex}>
                    <Flex
                        style={{
                            flexDirection: 'row',
                            justifyContent: 'left',
                            gap: '8px',
                            marginTop: '24px',
                            marginBottom: '10px',
                        }}
                    >
                        <DisplayImage src={league.logo} alt={league.name} />
                        <Link to={`/leagues/${league.slug}`}>
                            <Typography
                                variant='h6'
                                component='h2'
                                gutterBottom
                                sx={{
                                    '&:hover': {
                                        textDecoration: 'underline', // Add underline on hover
                                    },
                                }}
                            >
                                {league.name}
                            </Typography>
                        </Link>
                    </Flex>
                    <Grid container spacing={1}>
                        {league.fixtures.map((fixture, fixtureIndex) => (
                            <Grid item xs={12} sm={12} md={12} lg={12} key={fixtureIndex}>
                                <ListFixtureCard
                                    slug={fixture.slug}
                                    status={fixture.status}
                                    elapsed={fixture.elapsed}
                                    homeName={fixture.homeTeam.name}
                                    homeSlug={fixture.homeTeam.slug}
                                    homeLogo={fixture.homeTeam.logo}
                                    homeScore={fixture.score.home}
                                    awayName={fixture.awayTeam.name}
                                    awaySlug={fixture.awayTeam.slug}
                                    awayLogo={fixture.awayTeam.logo}
                                    awayScore={fixture.score.away}
                                    timezone={fixture.timezone}
                                />
                            </Grid>
                        ))}
                    </Grid>
                </Grid>
            ))}
        </Grid>
    );
};

export default ListFixtureTable;

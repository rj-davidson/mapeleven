import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Grid, Skeleton, Typography } from '@mui/material';
import { APIFixtureEvent, APIFixtureTeam } from '../Models/api_types';
import FixtureMatchEvents from './FixtureMatchEvents';
import FixturePageSingleMain from './FixturePageSingleMain';
import TeamIDCardFixturePage from '../Teams/TeamIDCardFixturePage';
import { Tile } from '../Util/Tile';

const url = import.meta.env.VITE_API_URL;

function FixturePageSingle() {
    const { slug } = useParams();

    const [homeTeamSlug, setHomeTeamSlug] = useState<string>('');
    const [awayTeamSlug, setAwayTeamSlug] = useState<string>('');
    const [homeTeam, setHomeTeam] = useState<string>('');
    const [awayTeam, setAwayTeam] = useState<string>('');
    const [homeLogo, setHomeLogo] = useState<string>('');
    const [awayLogo, setAwayLogo] = useState<string>('');
    const [referee, setReferee] = useState<string>('');
    const [homeTeamScore, setHomeTeamScore] = useState<number>(0);
    const [awayTeamScore, setAwayTeamScore] = useState<number>(0);
    const [minute, setMinute] = useState<number>(0);
    const [date, setDate] = useState<Date>(new Date());
    const [timeZone, setTimeZone] = useState<string>('');
    const [loading, setLoading] = useState<boolean>(true);
    const [events, setEvents] = useState<APIFixtureEvent[]>([]);
    const [teams, setTeams] = useState<APIFixtureTeam>();

    useEffect(() => {
        const fetchData = () => {
            fetch(`${url}/fixtures-demo/${slug}`)
                .then(response => response.json())
                .then(jsonData => {
                    setHomeTeamSlug(jsonData.teams.home.slug);
                    setAwayTeamSlug(jsonData.teams.away.slug);
                    setHomeTeam(jsonData.teams.home.name);
                    setAwayTeam(jsonData.teams.away.name);
                    setHomeLogo(jsonData.teams.home.logo);
                    setAwayLogo(jsonData.teams.away.logo);
                    setReferee(jsonData.fixture.referee);
                    setMinute(jsonData.fixture.elapsed);
                    setHomeTeamScore(jsonData.fixture.homeTeamScore);
                    setAwayTeamScore(jsonData.fixture.awayTeamScore);
                    setLoading(false);
                    setEvents(jsonData.events);
                    setDate(jsonData.fixture.date);
                    setTimeZone(jsonData.fixture.timezone);
                    setTeams(jsonData.teams);
                    console.log('Got data!');
                })
                .catch(error => console.error(error));
        };

        fetchData();
        const intervalId = setInterval(fetchData, 30000); // 30 seconds in milliseconds

        // Cleanup the interval on component unmount to avoid memory leaks
        return () => clearInterval(intervalId);
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
            <Grid item xs={12} sm={12}>
                <FixturePageSingleMain
                    homeTeam={homeTeam}
                    awayTeam={awayTeam}
                    homeLogo={homeLogo}
                    awayLogo={awayLogo}
                    referee={referee}
                    homeTeamScore={homeTeamScore}
                    awayTeamScore={awayTeamScore}
                    minute={minute}
                    events={events}
                    date={date}
                    timeZone={timeZone}
                />
            </Grid>
            <Grid item xs={12} sm={12}>
                <Grid container spacing={2}>
                    <Grid item xs={12} sm={3}>
                        <TeamIDCardFixturePage slug={homeTeamSlug} name={homeTeam} badge={homeLogo} />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Tile style={{ flexDirection: 'column' }}>
                            {events ? (
                                events.map((event, index) => (
                                    <Grid item xs={12} key={index}>
                                        <FixtureMatchEvents
                                            event={event}
                                            homeTeamSlug={homeTeamSlug}
                                            awayTeamSlug={awayTeamSlug}
                                        />
                                    </Grid>
                                ))
                            ) : (
                                // Handle the case when events is null or an empty array
                                <Typography variant='body2'>No events yet.</Typography>
                            )}
                        </Tile>
                    </Grid>
                    <Grid item xs={12} sm={3}>
                        <TeamIDCardFixturePage slug={awayTeamSlug} name={awayTeam} badge={awayLogo} />
                    </Grid>
                </Grid>
            </Grid>
        </Grid>
    );
}

export default FixturePageSingle;

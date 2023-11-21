import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Grid, Skeleton, Typography, useMediaQuery } from '@mui/material';
import { APIFixtureEvent, APIFixtureTeam } from '../Models/api_types';
import FixtureMatchEvents from './FixtureMatchEvents';
import FixturePageSingleMain from './FixturePageSingleMain';
import TeamIDCardFixturePage from '../Teams/TeamIDCardFixturePage';
import { Tile } from '../Util/Tile';
import FixtureLineup from './FixtureLineup';
import FixtureLineupMobile from './FixtureLineupMobile';
import FixturePlayerStepper from "./FixturePlayerStepper";
import FixturePlayerStepperDummy from "./FixturePlayerStepperDummy";

const url = import.meta.env.VITE_API_URL;

function FixturePageSingle() {
    const { slug } = useParams();
    const isSmallerThanLG = useMediaQuery('(max-width: 1260px)');

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
    const [homeLineup, setHomeLineup] = useState([]);
    const [awayLineup, setAwayLineup] = useState([]);

    useEffect(() => {
        fetch(`${url}/fixtures/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                setHomeTeamSlug(jsonData.teams.home.slug);
                setAwayTeamSlug(jsonData.teams.away.slug);
                setHomeTeam(jsonData.teams.home.name);
                setAwayTeam(jsonData.teams.away.name);
                setHomeLogo(jsonData.teams.home.logo);
                setAwayLogo(jsonData.teams.away.logo);
                setReferee(jsonData.fixture.referee);
                setHomeTeamScore(jsonData.fixture.homeTeamScore);
                setAwayTeamScore(jsonData.fixture.awayTeamScore);
                setMinute(jsonData.fixture.elapsed);
                setLoading(false);
                setEvents(jsonData.events);
                setDate(jsonData.fixture.date);
                setTimeZone(jsonData.fixture.timezone);
                setTeams(jsonData.teams);
                if (jsonData.lineups && jsonData.lineups.home && jsonData.lineups.away) {
                    setHomeLineup(jsonData.lineups.home.matchPlayers);
                    setAwayLineup(jsonData.lineups.away.matchPlayers);
                }
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
            {homeLineup && homeLineup.length > 0 && awayLineup && awayLineup.length > 0 && (
                <Grid item xs={12} sm={12}>
                    {isSmallerThanLG ? (
                        <FixtureLineupMobile homePlayers={awayLineup} awayPlayers={homeLineup} />
                    ) : (
                        <FixtureLineup homePlayers={awayLineup} awayPlayers={homeLineup} />
                    )}
                </Grid>
            )}
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
            <Grid item xs={12} sm={12}>
                <FixturePlayerStepperDummy/>
            </Grid>
        </Grid>
    );
}

export default FixturePageSingle;

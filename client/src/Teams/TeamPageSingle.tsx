import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Grid, Skeleton } from '@mui/material';
import { APITSGoalMinuteSplit } from '../Models/api_types';
import TeamPageSingleColumn from './TeamPageSingleColumn';
import TeamPageSingleMain from './TeamPageSingleMain';

const url = import.meta.env.VITE_API_URL;

function TeamPageSingle() {
    const { slug } = useParams();

    const [name, setName] = useState<string>('');
    const [badge, setBadge] = useState<string>('');
    const [flag, setFlag] = useState<string>('');
    const [country, setCountry] = useState<string>('');
    const [founded, setFounded] = useState<string>('');
    const [goals, setGoals] = useState<number>(0);
    const [averageGoals, setAverageGoals] = useState<number>(0);
    const [clean, setClean] = useState<number>(0);
    const [gamesPlayed, setGamesPlayed] = useState<number>(0);
    const [wins, setWins] = useState<number>(0);
    const [failedToScore, setFailedToScore] = useState<number>(0);
    const [gamesScored, setGamesScored] = useState<number>(0);
    const [goalMinuteSplit, setGoalMinuteSplit] = useState<APITSGoalMinuteSplit>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [draws, setDraws] = useState<number>(0);
    const [loses, setLoses] = useState<number>(0);

    useEffect(() => {
        fetch(`${url}/teams/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                setName(jsonData.name.long);
                setBadge(jsonData.badge);
                setFlag(jsonData.country.flag);
                setCountry(jsonData.country.name);
                setFounded(jsonData.founded);
                setGoals(jsonData.competitions[0].stats.goals.for.total.total);
                setAverageGoals(jsonData.competitions[0].stats.goals.for.average.total);
                setClean(jsonData.competitions[0].stats.clean_sheet.total);
                setGamesPlayed(
                    jsonData.competitions[0].stats.fixtures.played.home +
                        jsonData.competitions[0].stats.fixtures.played.away,
                );
                setWins(
                    jsonData.competitions[0].stats.fixtures.wins.home +
                        jsonData.competitions[0].stats.fixtures.wins.away,
                );
                setFailedToScore(jsonData.competitions[0].stats.failed_to_score.total);
                setGamesScored(
                    jsonData.competitions[0].stats.fixtures.played.home +
                        jsonData.competitions[0].stats.fixtures.played.away -
                        jsonData.competitions[0].stats.failed_to_score.total,
                );
                setGoalMinuteSplit(jsonData.competitions[0].stats.goals.for.minute);
                setLoading(false);
                setDraws(
                    jsonData.competitions[0].stats.fixtures.draws.home +
                        jsonData.competitions[0].stats.fixtures.draws.away,
                );
                setLoses(
                    jsonData.competitions[0].stats.fixtures.loses.home +
                        jsonData.competitions[0].stats.fixtures.loses.away,
                );
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
            <Grid item xs={12} sm={12} md={12} lg={3} width='100%'>
                <TeamPageSingleColumn
                    slug={slug}
                    name={name}
                    badge={badge}
                    country={country}
                    flag={flag}
                    founded={founded}
                    goals={goals}
                    clean={clean}
                    gamesPlayed={gamesPlayed}
                    wins={wins}
                    draws={draws}
                    loses={loses}
                    failedToScore={failedToScore}
                    averageGoals={averageGoals}
                    gamesScored={gamesScored}
                />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={9} width='100%'>
                <TeamPageSingleMain
                    goalMinuteSplit={goalMinuteSplit}
                    name={name}
                    goals={goals}
                    clean={clean}
                    fixtures={gamesPlayed}
                    wins={wins}
                    averageGoals={averageGoals}
                    gamesScored={gamesScored}
                    failedToScore={failedToScore}
                />
            </Grid>
        </Grid>
    );
}

export default TeamPageSingle;

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
    const [clean, setClean] = useState<number>(0);
    const [gamesPlayed, setGamesPlayed] = useState<number>(0);
    const [goalMinuteSplit, setGoalMinuteSplit] = useState<APITSGoalMinuteSplit>(null);
    const [loading, setLoading] = useState<boolean>(true);

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
                setClean(jsonData.competitions[0].stats.clean_sheet.total);
                setGamesPlayed(jsonData.competitions[0].stats.fixtures.played.home + jsonData.competitions[0].stats.fixtures.played.away);
                setGoalMinuteSplit(jsonData.competitions[0].stats.goals.for.minute);
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
            <Grid item xs={12} sm={12} md={12} lg={3} width='100%'>
                <TeamPageSingleColumn
                    name={name}
                    badge={badge}
                    country={country}
                    flag={flag}
                    founded={founded}
                    goals={goals}
                    clean={clean}
                    gamesPlayed={gamesPlayed}
                />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={9} width='100%'>
                <TeamPageSingleMain goalMinuteSplit={goalMinuteSplit} />
            </Grid>
        </Grid>
    );
}

export default TeamPageSingle;

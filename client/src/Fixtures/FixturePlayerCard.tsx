import {useEffect, useState} from "react";
import {Grid, Skeleton} from "@mui/material";
import * as React from "react";
import PlayerIDCard from "../Players/PlayerIDCard";
import {APIPlayer} from "../Models/api_types";

const url = import.meta.env.VITE_API_URL;

interface FixturePlayerCardProps {
    slug: string;
}

const FixturePlayerCard: React.FC<FixturePlayerCardProps> = ({ slug }) => {

    const [name, setName] = useState<string>('');
    const [firstName, setFirstName] = useState<string>('');
    const [lastName, setLastName] = useState<string>('');
    const [photo, setPhoto] = useState<string>('');
    const [age, setAge] = useState<number>(0);
    const [height, setHeight] = useState<string>('');
    const [weight, setWeight] = useState<string>('');
    const [injured, setInjured] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(true);
    const [playerData, setPlayerData] = useState<APIPlayer>();

    useEffect(() => {
        fetch(`${url}/players/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                setPlayerData(jsonData as APIPlayer);
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
                <Grid container spacing={2}>
                    <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                        <PlayerIDCard
                            slug={slug} playerData={playerData}
                        />
                    </Grid>
                </Grid>
            </Grid>
        </Grid>
    );
}

export default FixturePlayerCard;
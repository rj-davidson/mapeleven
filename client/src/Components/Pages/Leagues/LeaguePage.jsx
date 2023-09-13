import React, {useState, useRef, useEffect} from "react";
import {useParams} from "react-router-dom";
import * as d3 from 'd3';
import {Box, Grid, Skeleton, Typography} from "@mui/material";
import {Tile} from "../../Util/Tile.jsx";
import {useNavigate} from "react-router-dom";
import {Flex} from "../../Util/Flex.jsx";
import Standings from "../Fixtures/Standings.jsx";

function LeaguePage() {

    const navigate = useNavigate();
    const {slug} = useParams();

    // Declare state variables using the `useState` hook.
    const [name, setName] = useState();
    const [logo, setLogo] = useState();
    const [flag, setFlag] = useState("");
    const [country, setCountry] = useState("");
    const [teams, setTeams] = useState(new Set());
    const [loading, setLoading] = useState(true);

    useEffect(() => {

        setLoading(true);

        // Send a GET request to the API.
        fetch(`http://localhost:8080/leagues/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                console.log(jsonData);
                setName(jsonData.name);
                setLogo(jsonData.logo);
                setFlag(jsonData.country.flag);
                setCountry(jsonData.country.name);

                const standings = jsonData.standings;
                const newTeams = new Set();
                for (let i = 0; i < standings.length; i++) {
                    if (standings[i].rank > 0) {
                        newTeams.add({
                            rank: standings[i].rank,
                            name: standings[i].team.name.long,
                            logo: standings[i].team.badge,
                            points: standings[i].points,
                            played: standings[i].overall.played,
                            wins: standings[i].overall.won,
                            draw: standings[i].overall.draw,
                            losses: standings[i].overall.lost,
                        });
                    }
                }
                // Update the state variables with the new data.
                setTeams(newTeams);
                setLoading(false);
            })
            .catch(error => console.error(error));
    }, [slug]);

    return (

        loading ? (<Grid container spacing={2} direction="column">
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Skeleton variant="rectangular" height="150px" width="100%"
                          sx={{background: 'gray', borderRadius: '12px'}}/>
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Skeleton variant="rectangular" height="600px" width="100%"
                          sx={{background: 'gray', borderRadius: '12px'}}/>
            </Grid>
        </Grid>) : (<Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile style={{justifyContent: 'space-between'}}>
                    <Flex style={{justifyContent: 'center', alignItems: 'center', gap: '32px'}}>
                        <Box
                            component="img"
                            sx={{
                                maxHeight: {xs: 64, sm: 96}, maxWidth: {xs: 64, sm: 96},
                            }}
                            alt={name}
                            src={'http://localhost:8080/' + logo}
                        />
                        <Typography sx={{fontSize: {xs: '24px', sm: '48px'},}}>
                            {name}
                        </Typography>
                    </Flex>
                    <Flex style={{justifyContent: 'center', alignItems: 'center', flexDirection: 'column'}}>
                        <Box
                            component="img"
                            sx={{
                                maxHeight: {xs: 48, sm: 48}, maxWidth: {xs: 60, sm: 60},
                            }}
                            alt={country}
                            src={'http://localhost:8080/' + flag}
                        />
                    </Flex>
                </Tile>
            </Grid>

            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Tile style={{justifyContent: 'left'}}>
                    <Standings data={teams}/>
                </Tile>
            </Grid>
        </Grid>));
}

export default LeaguePage;


import React, {useState, useRef, useEffect} from "react";
import {useParams} from "react-router-dom";
import * as d3 from 'd3';
import {Box, Grid, Skeleton, Typography} from "@mui/material";
import {Tile} from "../../Util/Tile.jsx";
import {useNavigate} from "react-router-dom";
import {Flex} from "../../Util/Flex.jsx";
import Standings from "../Fixtures/Standings.jsx";

function FormCard() {

    const navigate = useNavigate();
    const {slug} = useParams();

    // Declare state variables using the `useState` hook.
    const [form, setForm] = useState();
    const [stats, setStats] = useState();
    const [loading, setLoading] = useState(true);

    const characterColors = {
        W: 'green',
        L: 'red',
    };

    useEffect(() => {

        setLoading(true);

        // Send a GET request to the API.
        fetch(`http://localhost:8080/teams/${slug}`)
            .then(response => response.json())
            .then(jsonData => {
                console.log(jsonData);
                setForm(jsonData.stats.form)


                setLoading(false);
            })
            .catch(error => console.error(error));
    }, [slug]);

    return (

        loading ? (<Grid direction="column">
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Skeleton variant="rectangular" height="50px" width="100%"
                          sx={{background: 'gray', borderRadius: '12px'}}/>
            </Grid>
        </Grid>) : (<Grid >
           <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>

                <Tile style = {{justifyContent: 'space-between',  height:'50px'}}>
                    <Flex style={{justifyContent: 'center', alignItems: 'center', gap: '32px'}}>

                        <Typography variant='h6' component='h2' style={{ marginBottom: '20 px' }}>
                            Form:
                        </Typography>
                    </Flex>
                    <Flex style={{justifyContent: 'center', alignItems: 'center', gap: '32px'}}>

                        <Typography variant='h6' component='h2' sx={{ mr: 0 }}>
                            {/* IS IT FIRST 5 OR LAST 5*/}
                            {form.slice(0,5).split('').map((character, index) => (
                                <Typography key={index} variant='h6' component='span' style={{ color: characterColors[character] || 'white' }}>
                                    {character}
                                </Typography>
                            ))}
                        </Typography>
                    </Flex>

                </Tile>
            </Grid>
        </Grid>));
}

export default FormCard;


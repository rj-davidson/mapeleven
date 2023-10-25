import * as React from 'react';
import { Box, Typography, Avatar, Grid, Divider } from '@mui/material';
import { APIFixtureEvent } from '../Models/api_types';
import DisplayImage from '../Util/DisplayImage';
import {RectangleRounded, SportsSoccer, Support, ArrowUpward, ArrowDownward} from '@mui/icons-material';
import {Link} from "react-router-dom";

interface MatchEventProps {
    event: APIFixtureEvent;
}

const FixtureGoalScorerCard: React.FC<MatchEventProps> = ({ event }) => {
    const getIcon = (event:APIFixtureEvent)=>{
        if(event.type == 'Card' && event.detail == 'Yellow Card'){
            return <RectangleRounded style={{ fill: 'var(--yellow)' , transform:'rotate(90deg)'}} />
        }
        else if(event.type == 'Card' && event.detail == 'Red Card'){
            return <RectangleRounded style={{ fill: 'var(--red)' , transform:'rotate(90deg)'}} />
        }

    };

    return (
        <Grid container alignItems="left" spacing={2} >
            <Grid item>
                {event.team.slug === 'home' && (
                    <DisplayImage src={event.team.logo} alt={event.team.name} />
                )}
            </Grid>

            <Grid item xs={6}>

                <Typography variant="body1">
                    {event.time.elapsed} { event.time.extra > 0 && `+${event.time.extra}'`}
                </Typography>

                {event.type === 'subst' ? (

                    <div>
                        <Link to={`/players/${event.player.slug}`}>
                            <Typography
                                sx={{
                                    '&:hover': {
                                        textDecoration: 'underline', // Add underline on hover
                                    },
                                }}
                            >
                                <ArrowDownward style={{ fill: 'var(--red)' }} />
                                {event.player.name}
                            </Typography>
                        </Link>

                        <br />
                        <Link to={`/players/${event.assist.slug}`}>
                            <Typography
                                sx={{
                                    '&:hover': {
                                        textDecoration: 'underline', // Add underline on hover
                                    },
                                }}
                            >
                                <ArrowUpward style={{ fill: 'var(--green)' }} />
                                {event.assist.name}
                            </Typography>
                        </Link>
                    </div>
                ) : null}

                {event.type === 'Goal' ? (
                        <div>
                            <Link to={`/players/${event.player.slug}`}>
                                <Typography
                                    sx={{
                                        '&:hover': {
                                            textDecoration: 'underline', // Add underline on hover
                                        },
                                    }}
                                >
                                    <SportsSoccer />
                                    {event.player.name}
                                </Typography>
                            </Link>
                            {event.assist && (
                                <>
                                    <br />
                                    <Link to={`/players/${event.player.slug}`}>
                                        <Typography
                                            sx={{
                                                '&:hover': {
                                                    textDecoration: 'underline', // Add underline on hover
                                                },
                                            }}
                                        >
                                            <Support />
                                            {event.assist.name}
                                        </Typography>
                                    </Link>
                                </>
                            )}
                        </div>
                    ) :
                    null
                }

                {getIcon(event) && (
                    <Typography variant="body2">
                        <Link to={`/players/${event.player.slug}`}>
                            <Typography
                                sx={{
                                    '&:hover': {
                                        textDecoration: 'underline', // Add underline on hover
                                    },
                                }}
                            >
                                {getIcon(event)} {event.player.name}
                            </Typography>
                        </Link>
                    </Typography>
                )}




                {event.comments && (
                    <Typography variant="caption">
                        {event.comments}
                    </Typography>
                )}
            </Grid>

            <Grid item>
                {event.team.slug === 'away' && (
                    <DisplayImage src={event.team.logo} alt={event.team.name} />
                )}
            </Grid>

            {/* Add a horizontal divider for separation */}
            <Grid item xs={12}>
                <Divider />
            </Grid>
        </Grid>
    );
};

export default FixtureGoalScorerCard;

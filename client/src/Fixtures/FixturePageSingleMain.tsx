import * as React from 'react';
import { Grid, Typography } from '@mui/material';
import { APIFixtureEvent } from '../Models/api_types';
import DisplayImage from '../Util/DisplayImage';
import { Tile } from '../Util/TileTS';

interface FixturePageSingleMainProps {
    homeTeam: string;
    awayTeam: string;
    homeLogo: string;
    awayLogo: string;
    referee: string;
    homeTeamScore: number;
    awayTeamScore: number;
    minute: number;
    events: APIFixtureEvent[];
    date: Date;
    timeZone: string;
}

const FixturePageSingleMain: React.FC<FixturePageSingleMainProps> = ({
    homeTeam,
    awayTeam,
    homeLogo,
    awayLogo,
    referee,
    homeTeamScore,
    awayTeamScore,
    minute,
    events,
    date,
    timeZone,
}) => {
    return (
        <Grid item xs={12} sm={12}>
            <Tile>
                <Grid container spacing={2} justifyContent='center'>
                    <Grid item xs={12} sm={4}>
                        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                            <DisplayImage src={homeLogo} alt={homeTeam} width={'100px'} height={'100px'}></DisplayImage>
                            <Typography variant='h6' align='center' sx={{ fontSize: '42px' }}>
                                {homeTeam}
                            </Typography>
                        </div>

                        {events ? (
                            events
                                .filter(event => event.type === 'Goal' && event.team.name === homeTeam)
                                .map((event, index) => (
                                    <Typography
                                        variant='body2'
                                        align='center'
                                        key={index}
                                        sx={{
                                            fontSize: '20px', // Adjust the font size as needed
                                        }}
                                    >
                                        {event.player.name} {event.time.elapsed}'{' '}
                                        {event.time.extra > 0 && `+${event.time.extra}'`}
                                    </Typography>
                                ))
                        ) : (
                            <Typography variant='body2'>No events available.</Typography>
                        )}
                    </Grid>
                    <Grid item xs={12} sm={4}>
                        <Typography
                            variant='h4'
                            align='center'
                            sx={{
                                fontSize: '42px', // Adjust the font size as needed
                            }}
                        >
                            {homeTeamScore} - {awayTeamScore}
                        </Typography>
                        <Typography
                            variant='h6'
                            align='center'
                            sx={{
                                fontSize: '20px', // Adjust the font size as needed
                            }}
                        >
                            {minute}'
                        </Typography>
                    </Grid>
                    <Grid item xs={12} sm={4}>
                        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                            <DisplayImage src={awayLogo} alt={awayTeam} width={'100px'} height={'100px'}></DisplayImage>
                            <Typography variant='h6' align='center' sx={{ fontSize: '42px' }}>
                                {awayTeam}
                            </Typography>
                        </div>
                        {events ? (
                            events
                                .filter(event => event.type === 'Goal' && event.team.name === awayTeam)
                                .map((event, index) => (
                                    <Typography
                                        variant='body2'
                                        align='center'
                                        key={index}
                                        sx={{
                                            fontSize: '20px', // Adjust the font size as needed
                                        }}
                                    >
                                        {event.player.name} {event.time.elapsed}'{' '}
                                        {event.time.extra > 0 && `+${event.time.extra}'`}
                                    </Typography>
                                ))
                        ) : (
                            <Typography variant='body2'>No events available.</Typography>
                        )}
                    </Grid>
                    <Grid item xs={12} sm={12} sx={{ textAlign: 'center' }}>
                        <Typography
                            variant='body1'
                            align='center'
                            sx={{
                                fontSize: '20px', // Adjust the font size as needed
                            }}
                        >
                            {new Date(date.toString()).toDateString()} {timeZone}
                        </Typography>
                    </Grid>
                    <Grid item xs={12} sm={12}>
                        <Typography
                            variant='body1'
                            align='center'
                            sx={{
                                fontSize: '20px', // Adjust the font size as needed
                            }}
                        >
                            Referee: {referee}
                        </Typography>
                    </Grid>
                </Grid>
            </Tile>
        </Grid>
    );
};

export default FixturePageSingleMain;

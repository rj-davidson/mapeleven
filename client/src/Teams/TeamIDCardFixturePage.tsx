import * as React from 'react';
import { useEffect, useState } from 'react';
import { Box, Grid, Typography } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import { Flex } from '../Util/Flex.jsx';
import Tilt from 'react-parallax-tilt';
import DisplayImage from '../Util/DisplayImage';
import TeamIDTitle from '../Teams/TeamIDTitle';
import TeamIDStatBox from '../Teams/TeamIDStatBox';
import { useMediaQuery } from '@mui/material';
import { useParams } from 'react-router-dom';
import { APITSGoalMinuteSplit } from '../Models/api_types';

const url = import.meta.env.VITE_API_URL;

interface TeamIDCardProps {
    slug: string;
    name: string;
    badge: string;
}

const TeamIDCardFixturePage: React.FC<TeamIDCardProps> = ({ slug, name, badge }) => {
    const [isHovered, setIsHovered] = useState<boolean>(false);
    const [inspect, setInspect] = useState<boolean>(false);

    const isSmallerThanLG = useMediaQuery('(max-width: 1260px)');
    const [country, setCountry] = useState<string>('');
    const [flag, setFlag] = useState<string>('');
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

    return (
        <Box>
            {!isSmallerThanLG ? (
                <Box
                    onMouseEnter={() => setIsHovered(true)}
                    onMouseLeave={() => setIsHovered(false)}
                    onClick={() => setInspect(!inspect)}
                >
                    <Tilt
                        reset={inspect}
                        glareEnable={true}
                        tiltMaxAngleX={10}
                        tiltMaxAngleY={10}
                        tiltAngleXManual={inspect ? null : 0}
                        tiltAngleYManual={inspect ? null : 0}
                        perspective={1000}
                        glareMaxOpacity={0.2}
                        glareBorderRadius={'10px'}
                        tiltReverse={true}
                        style={{
                            cursor: isHovered ? 'pointer' : 'default',
                            borderRadius: '10px',
                            boxShadow: isHovered ? 'rgba(0,0,0,0.2) 0 0 10px' : 'none',
                        }}
                    >
                        <Tile
                            style={{
                                flexDirection: 'column',
                                gap: '24px',
                            }}
                        >
                            <TeamIDTitle slug={slug} name={name} country={country} flag={flag} />

                            <Flex style={{ justifyContent: 'center', alignItems: 'center', marginTop: '-12px' }}>
                                <DisplayImage src={badge} alt={name} width={'108px'} height={'108'} />
                            </Flex>

                            <Grid container spacing={1}>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <TeamIDStatBox stat='Points' value={3 * wins + draws} />
                                </Grid>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <TeamIDStatBox stat='Games Played' value={gamesPlayed} />
                                </Grid>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <TeamIDStatBox stat='Wins' value={wins} />
                                </Grid>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <TeamIDStatBox stat='Draws' value={draws} />
                                </Grid>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <TeamIDStatBox stat='Losses' value={loses} />
                                </Grid>
                            </Grid>

                            <Typography sx={{ fontSize: '16px', textAlign: 'right', fontStyle: 'italic' }}>
                                Founded {founded}
                            </Typography>
                        </Tile>
                    </Tilt>
                </Box>
            ) : (
                <Tile
                    style={{
                        flexDirection: 'column',
                        gap: '24px',
                        border: '5px solid var(--red)',
                    }}
                >
                    <TeamIDTitle slug={slug} name={name} country={country} flag={flag} />

                    <Flex style={{ justifyContent: 'center', alignItems: 'center', marginTop: '-12px' }}>
                        <DisplayImage src={badge} alt={name} width={'108px'} height={'108'} />
                    </Flex>

                    <Grid container spacing={1}>
                        <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                            <TeamIDStatBox stat='Position' value={'2nd'} />
                        </Grid>
                        <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                            <TeamIDStatBox stat='Points' value={3 * wins + draws} />
                        </Grid>
                        <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                            <TeamIDStatBox stat='Games Played' value={gamesPlayed} />
                        </Grid>
                    </Grid>

                    <Typography sx={{ fontSize: '16px', textAlign: 'right', fontStyle: 'italic' }}>
                        Founded {founded}
                    </Typography>
                </Tile>
            )}
        </Box>
    );
};

export default TeamIDCardFixturePage;

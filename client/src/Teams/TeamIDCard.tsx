import * as React from 'react';
import { useEffect, useState } from 'react';
import { Box, Grid, Typography } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import { Flex } from '../Util/Flex.jsx';
import Tilt from 'react-parallax-tilt';
import DisplayImage from '../Util/DisplayImage';
import TeamIDTitle from './TeamIDTitle';
import TeamIDStatBox from './TeamIDStatBox';
import { useMediaQuery } from '@mui/material';

interface TeamIDCardProps {
    slug: string;
    name: string;
    badge: string;
    country: string;
    flag: string;
    founded: string;
    goals: number;
    clean: number;
    gamesPlayed: number;
    wins: number;
    draws: number;
    loses: number;
    leagueSlug: string;
    leagueName: string;
    leagueLogo: string;
    rank: string;
}

const url = import.meta.env.VITE_API_URL;

const TeamIDCard: React.FC<TeamIDCardProps> = ({
    slug,
    name,
    badge,
    country,
    flag,
    founded,
    goals,
    clean,
    gamesPlayed,
    wins,
    draws,
    loses,
    leagueSlug,
    leagueName,
    leagueLogo,
    rank,
}) => {
    const [isHovered, setIsHovered] = useState<boolean>(false);
    const [inspect, setInspect] = useState<boolean>(false);

    const isSmallerThanLG = useMediaQuery('(max-width: 1260px)');

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
                                gap: '32px',
                                border: '5px solid var(--red)',
                            }}
                        >
                            <TeamIDTitle
                                slug={slug}
                                name={name}
                                country={country}
                                flag={flag}
                                leagueSlug={leagueSlug}
                                leagueName={leagueName}
                                leagueLogo={leagueLogo}
                            />

                            <Flex style={{ justifyContent: 'center', alignItems: 'center', marginTop: '-12px' }}>
                                <DisplayImage src={badge} alt={name} width={'108px'} height={'108'} />
                            </Flex>

                            <Grid container spacing={1}>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <TeamIDStatBox stat='Position' value={rank} />
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
                    </Tilt>
                </Box>
            ) : (
                <Tile
                    style={{
                        flexDirection: 'column',
                        gap: '32px',
                        border: '5px solid var(--red)',
                    }}
                >
                    <TeamIDTitle
                        slug={slug}
                        name={name}
                        country={country}
                        flag={flag}
                        leagueSlug={leagueSlug}
                        leagueName={leagueName}
                        leagueLogo={leagueLogo}
                    />

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

export default TeamIDCard;

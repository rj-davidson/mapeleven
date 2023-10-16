import * as React from 'react';
import { useState } from 'react';
import { Box, Grid, Typography } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import { Flex } from '../Util/Flex.jsx';
import Tilt from 'react-parallax-tilt';
import DisplayImage from '../Util/DisplayImage';
import TeamIDTitle from './TeamIDTitle';
import TeamIDStatBox from './TeamIDStatBox';

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
    wins: number,
    draws: number,
    loses: number,
}

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
    loses
}) => {
    const [isHovered, setIsHovered] = useState<boolean>(false);
    const [inspect, setInspect] = useState<boolean>(false);

    return (
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
                            <TeamIDStatBox stat='Position' value={0} />
                        </Grid>
                        <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                            <TeamIDStatBox stat='Points' value={0} />
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
                            <TeamIDStatBox stat='Loses' value={loses} />
                        </Grid>
                    </Grid>

                    <Typography sx={{ fontSize: '16px', textAlign: 'right', fontStyle: 'italic' }}>
                        Founded {founded}
                    </Typography>
                </Tile>
            </Tilt>
        </Box>
    );
};

export default TeamIDCard;

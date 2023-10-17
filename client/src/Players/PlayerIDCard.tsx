import * as React from 'react';
import { useState } from 'react';
import { Box, Grid, Typography, useMediaQuery } from '@mui/material';
import { Tile } from '../Util/Tile.jsx';
import { Flex } from '../Util/Flex.jsx';
import Tilt from 'react-parallax-tilt';
import DisplayImage from '../Util/DisplayImage';
import PlayerIDTitle from './PlayerIDTitle';
import PlayerIDStatBox from './PlayerIDStatBox';

interface PlayerIDCardProps {
    slug: string;
    name: string;
    firstName: string;
    lastName: string;
    photo: string;
    age: number;
    height: string;
    weight: string;
    injured: boolean;
}

const PlayerIDCard: React.FC<PlayerIDCardProps> = ({
    slug,
    name,
    firstName,
    lastName,
    photo,
    age,
    height,
    weight,
    injured,
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
                                gap: '48px',
                                border: '5px solid var(--red)',
                            }}
                        >
                            <PlayerIDTitle slug={slug} name={name} country={'Country'} flag={''} injured={injured} />

                            <Flex style={{ justifyContent: 'center', alignItems: 'center', marginTop: '-12px' }}>
                                <DisplayImage
                                    src={photo}
                                    alt={name}
                                    width={'108px'}
                                    height={'108'}
                                    sx={{
                                        borderRadius: '10px',
                                        boxShadow: 'rgba(0,0,0,0.2) 0 0 10px',
                                    }}
                                />
                            </Flex>

                            <Grid container spacing={1}>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <PlayerIDStatBox stat='Age' value={age} />
                                </Grid>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <PlayerIDStatBox stat='Height' value={height} />
                                </Grid>
                                <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                                    <PlayerIDStatBox stat='Weight' value={weight} />
                                </Grid>
                            </Grid>
                        </Tile>
                    </Tilt>
                </Box>
            ) : (
                <Tile
                    style={{
                        flexDirection: 'column',
                        gap: '48px',
                        border: '5px solid var(--red)',
                    }}
                >
                    <PlayerIDTitle slug={slug} name={name} country={'Country'} flag={''} injured={injured} />

                    <Flex style={{ justifyContent: 'center', alignItems: 'center', marginTop: '-12px' }}>
                        <DisplayImage
                            src={photo}
                            alt={name}
                            width={'108px'}
                            height={'108'}
                            sx={{
                                borderRadius: '10px',
                                boxShadow: 'rgba(0,0,0,0.2) 0 0 10px',
                            }}
                        />
                    </Flex>

                    <Grid container spacing={1}>
                        <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                            <PlayerIDStatBox stat='Age' value={age} />
                        </Grid>
                        <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                            <PlayerIDStatBox stat='Height' value={height} />
                        </Grid>
                        <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                            <PlayerIDStatBox stat='Weight' value={weight} />
                        </Grid>
                    </Grid>
                </Tile>
            )}
        </Box>
    );
};

export default PlayerIDCard;

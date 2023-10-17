import * as React from 'react';
import { Grid } from '@mui/material';
import PlayerIDCard from './PlayerIDCard';

interface PlayerPageSingleColumnProps {
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

const PlayerPageSingleColumn: React.FC<PlayerPageSingleColumnProps> = ({
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
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <PlayerIDCard
                    slug={slug}
                    name={name}
                    firstName={firstName}
                    lastName={lastName}
                    photo={photo}
                    age={age}
                    height={height}
                    weight={weight}
                    injured={injured}
                />
            </Grid>
        </Grid>
    );
};

export default PlayerPageSingleColumn;

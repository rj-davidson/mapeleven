import * as React from 'react';
import { Typography } from '@mui/material';
import { Tile } from '../Util/TileTS';

interface PlayerIDStatBoxProps {
    stat: string;
    value: number | string;
}

const PlayerIDStatBox: React.FC<PlayerIDStatBoxProps> = ({ stat, value }) => {
    return (
        <Tile
            sx={{
                justifyContent: 'space-between',
                background: 'var(--dark1)',
                border: 'none',
            }}
        >
            <Typography sx={{ fontSize: '18px' }}>{stat}:</Typography>
            <Typography sx={{ fontSize: '18px' }}>{value}</Typography>
        </Tile>
    );
};

export default PlayerIDStatBox;

import * as React from 'react';
import { Box, Card, CardContent, Typography } from '@mui/material';
import DisplayImage from '../../Util/DisplayImage';
import { Flex } from '../../Util/Flex.jsx';
import { Link } from 'react-router-dom';

interface PlayerCardProps {
    name: string;
    photo: string;
}

const PlayerCard: React.FC<PlayerCardProps> = ({
    name,
    photo,
}) => {
    return (
        <Link to={`/players/${name}`}>
            <Card>
                <CardContent sx={{ padding: '8px' }}>
                    <Flex style={{ flexDirection: 'row', gap: '16px', justifyContent: 'left', alignItems: 'center' }}>
                        <Box sx={{ display: 'flex', justifyContent: 'left', alignItems: 'center'}}>
                            <DisplayImage src={photo} alt={name} location={""} sx={{borderRadius: '8px'}}/>
                        </Box>
                        <Typography variant='h6' component='h2' sx={{ fontSize: '16px' }}>
                            {name}
                        </Typography>
                    </Flex>
                </CardContent>
            </Card>
        </Link>
    );
};

export default PlayerCard;

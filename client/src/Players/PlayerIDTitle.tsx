import * as React from 'react';
import { Typography } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import DisplayImage from '../Util/DisplayImage';
import { Link } from 'react-router-dom';
import HealingIcon from '@mui/icons-material/Healing';

interface PlayerIDTitleProps {
    slug: string;
    name: string;
    country: string;
    flag: string;
    injured: boolean;
}

const PlayerIDTitle: React.FC<PlayerIDTitleProps> = ({ name, country, flag, slug, injured }) => {
    return (
        <Flex style={{ justifyContent: 'space-between' }}>
            <Flex style={{ flexDirection: 'column' }}>
                <Link to={`/players/${slug}`}>
                    <Typography
                        sx={{
                            fontSize: '24px',
                            '&:hover': {
                                textDecoration: 'underline', // Add underline on hover
                            },
                        }}
                    >
                        {name}
                    </Typography>
                </Link>

                <Flex style={{ flexDirection: 'row', gap: '6px', alignItems: 'center' }}>
                    {/*<DisplayImage src={flag} alt={country} width={'20px'} height={'16px'} />*/}
                    <Typography sx={{ fontSize: '14px', color: 'var(--light0)' }}>{country}</Typography>
                </Flex>
            </Flex>
            <Flex>{injured && <HealingIcon style={{ fill: 'var(--red)' }} />}</Flex>
        </Flex>
    );
};

export default PlayerIDTitle;

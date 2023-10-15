import * as React from 'react';
import { Typography } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import DisplayImage from '../Util/DisplayImage';
import { Link } from 'react-router-dom';

interface TeamIDTitleProps {
    slug: string;
    name: string;
    country: string;
    flag: string;
}

const TeamIDTitle: React.FC<TeamIDTitleProps> = ({ name, country, flag, slug }) => {
    return (
        <Flex style={{ flexDirection: 'column' }}>
            <Link to={`/teams/${slug}`}>
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
                <DisplayImage src={flag} alt={country} width={'20px'} height={'16px'} />
                <Typography sx={{ fontSize: '14px', color: 'var(--light0)' }}>{country}</Typography>
            </Flex>
        </Flex>
    );
};

export default TeamIDTitle;

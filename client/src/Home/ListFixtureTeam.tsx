import * as React from 'react';
import { Typography } from '@mui/material';
import DisplayImage from '../Util/DisplayImage';
import { Flex } from '../Util/Flex.jsx';
import { Link } from 'react-router-dom';

interface ListFixtureTeamProps {
    slug: string;
    name: string;
    logo: string;
    home: boolean;
}

const ListFixtureTeam: React.FC<ListFixtureTeamProps> = ({ slug, name, logo, home }) => {
    return home ? (
        <Flex
            style={{
                justifyContent: 'right',
                alignItems: 'center',
                width: '80%',
                gap: '8px',
            }}
        >
            <Link to={`/teams/${slug}`}>
                <Typography
                    variant='h6'
                    component='h2'
                    sx={{
                        fontSize: '16px',
                        '&:hover': {
                            textDecoration: 'underline', // Add underline on hover
                        },
                    }}
                >
                    {name}
                </Typography>
            </Link>
            <DisplayImage src={logo} alt={name} />
        </Flex>
    ) : (
        <Flex
            style={{
                justifyContent: 'left',
                alignItems: 'center',
                width: '80%',
                gap: '8px',
            }}
        >
            <DisplayImage src={logo} alt={name} />
            <Link to={`/teams/${slug}`}>
                <Typography
                    variant='h6'
                    component='h2'
                    sx={{
                        fontSize: '16px',
                        '&:hover': {
                            textDecoration: 'underline', // Add underline on hover
                        },
                    }}
                >
                    {name}
                </Typography>
            </Link>
        </Flex>
    );
};

export default ListFixtureTeam;

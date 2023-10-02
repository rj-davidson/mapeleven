import * as React from 'react';
import { Typography } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import DisplayImage from '../Util/DisplayImage';

interface TeamIDTitleProps {
    name: string;
    country: string;
    flag: string;
}

const TeamIDTitle: React.FC<TeamIDTitleProps> = ({ name, country, flag }) => {
    return (
        <Flex style={{ flexDirection: 'column' }}>
            <Typography sx={{ fontSize: '24px' }}>{name}</Typography>

            <Flex style={{ flexDirection: 'row', gap: '6px', alignItems: 'center' }}>
                <DisplayImage src={flag} alt={country} width={'20px'} height={'16px'} />
                <Typography sx={{ fontSize: '14px', color: 'var(--light0)' }}>{country}</Typography>
            </Flex>
        </Flex>
    );
};

export default TeamIDTitle;

import * as React from 'react';
import { Box, Typography } from '@mui/material';

type FixtureTeamProps = {
    logo: string;
    name: string;
    score: number;
};

const styles = {
    teamWrapper: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        height: 44,
    },
    teamName: {
        flexGrow: 1,
        paddingLeft: 10,
    },
    teamScore: {
        fontSize: 20,
        fontWeight: 'bold',
    },
};

const FixtureTeam: React.FC<FixtureTeamProps> = ({ logo, name, score }) => {
    if (typeof logo !== 'string' || typeof name !== 'string' || typeof score !== 'number') {
        return null; // or handle the error in some other way
    }

    return (
        <Box sx={styles.teamWrapper}>
            <img src={logo} alt={`${name} logo`} width='44' height='44' />
            <Typography variant='body1' sx={styles.teamName}>
                {name}
            </Typography>
            <Typography variant='body1' sx={styles.teamScore}>
                {score}
            </Typography>
        </Box>
    );
};

export default FixtureTeam;

import * as React from 'react';
import { Grid } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import FormCard from './FormCard.jsx';
import TeamIDCard from './TeamIDCard';

interface TeamPageSingleColumnProps {
    name: string;
    badge: string;
    country: string;
    flag: string;
    founded: string;
    goals: number;
    clean: number;
    gamesPlayed: number;
}

const TeamPageSingleColumn: React.FC<TeamPageSingleColumnProps> = ({
    name,
    badge,
    country,
    flag,
    founded,
    goals,
    clean,
    gamesPlayed,
}) => {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <TeamIDCard
                    name={name}
                    badge={badge}
                    country={country}
                    flag={flag}
                    founded={founded}
                    goals={goals}
                    clean={clean}
                    gamesPlayed={gamesPlayed}
                />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Flex style={{ justifyContent: 'center', alignItems: 'center', flexDirection: 'column' }}>
                    <FormCard />
                </Flex>
            </Grid>
        </Grid>
    );
};

export default TeamPageSingleColumn;

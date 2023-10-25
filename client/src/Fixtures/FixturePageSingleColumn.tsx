import * as React from 'react';
import { Grid } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import FormCard from '../Teams/FormCard.jsx';
import TeamIDCard from '../Teams/TeamIDCard';
import TeamStats from '../Teams/TeamStats';

interface FixturePageSingleColumnProps {
    slug: string;
    homeTeam: string;
    awayTeam: string;
    awayLogo: string;
    homeLogo: string;
    date: string;
    timeZone: string;
    referee: string;
    awayTeamScore: number;
    homeTeamScore: number;
    minute: number;
    events: string;
    lineups: string;
}

const FixturePageSingleColumn: React.FC<FixturePageSingleColumnProps> = ({
                                                                             slug,
                                                                             homeTeam,
                                                                             awayTeam,
                                                                             homeLogo,
                                                                             awayLogo,
                                                                             date,
                                                                             timeZone,
                                                                             referee,
                                                                             homeTeamScore,
                                                                             awayTeamScore,
                                                                             minute,
                                                                             events,
                                                                             lineups,
}) => {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>

            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Flex style={{ justifyContent: 'center', alignItems: 'center', flexDirection: 'column' }}>
                    <FormCard />
                </Flex>
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Flex style={{ justifyContent: 'center', alignItems: 'center', flexDirection: 'column' }}>

                </Flex>
            </Grid>
        </Grid>
    );
};

export default FixturePageSingleColumn;

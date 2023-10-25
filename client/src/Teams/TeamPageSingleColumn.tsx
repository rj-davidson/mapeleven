import * as React from 'react';
import { Grid } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import FormCard from './FormCard.jsx';
import TeamIDCard from './TeamIDCard';
import TeamStats from './TeamStats';

interface TeamPageSingleColumnProps {
    slug: string;
    name: string;
    badge: string;
    country: string;
    flag: string;
    founded: string;
    goals: number;
    clean: number;
    gamesPlayed: number;
    wins: number;
    draws: number;
    loses: number;
    averageGoals: number;
    gamesScored: number;
    failedToScore: number;
    leagueSlug: string;
    leagueName: string;
    leagueLogo: string;
    rank: string;
}

const TeamPageSingleColumn: React.FC<TeamPageSingleColumnProps> = ({
    slug,
    name,
    badge,
    country,
    flag,
    founded,
    goals,
    clean,
    gamesPlayed,
    wins,
    draws,
    loses,
    averageGoals,
    gamesScored,
    failedToScore,
    leagueSlug,
    leagueName,
    leagueLogo,
    rank,
}) => {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <TeamIDCard
                    slug={slug}
                    name={name}
                    badge={badge}
                    country={country}
                    flag={flag}
                    founded={founded}
                    goals={goals}
                    clean={clean}
                    gamesPlayed={gamesPlayed}
                    wins={wins}
                    draws={draws}
                    loses={loses}
                    leagueSlug={leagueSlug}
                    leagueName={leagueName}
                    leagueLogo={leagueLogo}
                    rank={rank}
                />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Flex style={{ justifyContent: 'center', alignItems: 'center', flexDirection: 'column' }}>
                    <FormCard />
                </Flex>
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width='100%'>
                <Flex style={{ justifyContent: 'center', alignItems: 'center', flexDirection: 'column' }}>
                    <TeamStats
                        goals={goals}
                        clean={clean}
                        gamesPlayed={gamesPlayed}
                        wins={wins}
                        draws={draws}
                        loses={loses}
                        averageGoals={averageGoals}
                        failedToScore={failedToScore}
                        gamesScored={gamesScored}
                    />
                </Flex>
            </Grid>
        </Grid>
    );
};

export default TeamPageSingleColumn;

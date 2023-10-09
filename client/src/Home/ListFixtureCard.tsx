import * as React from 'react';
import { Card, CardContent } from '@mui/material';
import { Flex } from '../Util/Flex.jsx';
import { Link } from 'react-router-dom';
import ListFixtureTeam from './ListFixtureTeam';
import ListFixtureScore from './ListFixtureScore';
import ListFixtureTime from './ListFixtureTime';

interface ListFixtureCardProps {
    slug: string;
    status: string;
    elapsed: number;
    homeName: string;
    homeSlug: string;
    homeLogo: string;
    homeScore: number;
    awayName: string;
    awaySlug: string;
    awayLogo: string;
    awayScore: number;
    timezone: string;
}

const ListFixtureCard: React.FC<ListFixtureCardProps> = ({
    slug,
    status,
    elapsed,
    homeName,
    homeSlug,
    homeLogo,
    homeScore,
    awayName,
    awaySlug,
    awayLogo,
    awayScore,
    timezone,
}) => {
    return (
        <Link to={`/fixtures/${slug}`}>
            <Card>
                <CardContent sx={{ padding: '8px' }}>
                    <Flex style={{ flexDirection: 'row', gap: '4px', justifyContent: 'center', alignItems: 'center' }}>
                        <ListFixtureTime status={status} elapsed={elapsed} timezone={timezone} />
                        <ListFixtureTeam slug={homeSlug} name={homeName} logo={homeLogo} home={true} />
                        <ListFixtureScore homeScore={homeScore} awayScore={awayScore} status={status} />
                        <ListFixtureTeam slug={awaySlug} name={awayName} logo={awayLogo} home={false} />
                    </Flex>
                </CardContent>
            </Card>
        </Link>
    );
};

export default ListFixtureCard;

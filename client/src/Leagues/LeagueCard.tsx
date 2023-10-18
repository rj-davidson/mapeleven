import * as React from 'react';
import { Box, Card, CardContent, Typography } from '@mui/material';
import DisplayImage from '../Util/DisplayImage';
import { Flex } from '../Util/Flex.jsx';
import { Link } from 'react-router-dom';

interface LeagueCardProps {
    slug: string;
    leagueName: string;
    logo: string;
    countryCode: string;
    countryName: string;
    countryFlag: string;
}

const LeagueCard: React.FC<LeagueCardProps> = ({
                                               slug,
                                                   leagueName,
                                               logo,
                                               countryCode,
                                               countryName,
                                               countryFlag,
                                           }) => {
    return (
        <Link to={`/leagues/${slug}`}>
            <Card>
                <CardContent sx={{ padding: '8px', paddingRight: '16px' }}>
                    <Flex style={{ flexDirection: 'row', gap: '16px', justifyContent: 'space-between', alignItems: 'center' }}>
                        <Box sx={{ display: 'flex', justifyContent: 'left', alignItems: 'center' }}>
                            <DisplayImage src={logo} alt={leagueName} />
                        </Box>
                        <Typography variant='h6' component='h2' sx={{ fontSize: '16px', flex: '1' }}>
                            {leagueName}
                        </Typography>
                        <Box sx={{ display: 'flex', justifyContent: 'right', alignItems: 'center' }}>
                            <DisplayImage src={countryFlag} alt={countryName} />
                        </Box>
                    </Flex>
                </CardContent>
            </Card>
        </Link>
    );
};

export default LeagueCard;

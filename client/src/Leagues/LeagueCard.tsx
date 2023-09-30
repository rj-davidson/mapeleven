import * as React from 'react';
import { Card, CardContent, Grid, Typography } from '@mui/material';
import LeagueIcon from '@mui/icons-material/SportsSoccer';
import CupIcon from '@mui/icons-material/EmojiEvents';

const url = import.meta.env.VITE_API_URL;

interface League {
    id: number;
    logo: string;
    name: string;
    type: string;
}

interface DisplayLeaguesState {
    leagues: League[];
}

class LeagueCards extends React.Component<{}, DisplayLeaguesState> {
    constructor(props: {}) {
        super(props);
        this.state = {
            leagues: [],
        };
    }

    componentDidMount() {
        fetch(`${url}/leagues`)
            .then(response => response.json())
            .then(data => {
                this.setState({ leagues: data });
                console.log(data);
            });
    }

    render() {
        const leagueCards = this.state.leagues.map(league => {
            return (
                <Grid item xs={12} sm={12} md={12} lg={12} key={league.id}>
                    <Card sx={{ height: '100%' }}>
                        <CardContent
                            sx={{
                                flexGrow: 1,
                                display: 'flex',
                                alignItems: 'center',
                                justifyContent: 'space-between',
                                padding: '24px',
                            }}
                        >
                            <img src={`${url}/` + league.logo} alt={league.name + ' logo'} height={44} />
                            <Typography variant='h6' component='h2' sx={{ mr: 2 }}>
                                {league.name}
                            </Typography>
                            <Typography variant='body1' color='text.secondary' sx={{ mr: 2 }}>
                                {league.type.toLowerCase() === 'league' ? <LeagueIcon /> : <CupIcon />}
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>
            );
        });

        return (
            <div>
                <Typography variant='h4' component='h2' gutterBottom>
                    Leagues
                </Typography>
                <Grid container spacing={3} sx={{ justifyContent: 'space-between' }}>
                    {leagueCards}
                </Grid>
            </div>
        );
    }
}

export default LeagueCards;

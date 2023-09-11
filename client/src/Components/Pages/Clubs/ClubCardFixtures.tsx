import * as React from 'react';
import { Card, CardContent, CardMedia, Grid, Typography } from '@mui/material';

interface Club {
    id: number;
    logo: string;
    name: string;
    founded: string;
    code: string;
}

interface DisplayClubsState {
    clubs: Club[];
}

class ClubCard extends React.Component<{}, DisplayClubsState> {
    constructor(props: {}) {
        super(props);
        this.state = {
            clubs: [],
        };
    }

    componentDidMount() {
        fetch('http://localhost:8080/teams')
            .then(response => response.json())
            .then(data => {
                this.setState({ clubs: data });
                console.log(data);
            });
    }

    render() {
        const clubCards = this.state.clubs.map(club => {
            return (
                <Grid item xs={12} sm={12} md={12} lg={12} key={club.id}>
                    <Card sx={{ height: '100%' }}>
                        <CardContent
                            sx={{
                                flexGrow: 1,
                                display: 'flex',
                                alignItems: 'center',
                                justifyContent: 'center',
                                padding: '12px',
                            }}
                        >
                            <img
                                src={'http://localhost:8080/' + club.logo}
                                alt={club.name + ' logo'}
                                width={32}
                                height={32}
                                style={{ borderRadius: '8px', marginRight:'0px', border: '2px solid #D8DEE9' }}
                            />
                            <Typography variant='h6' component='h2' sx={{ mr: 0 }}>
                                {club.name}
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>
            );
        });

        return (
            <div>
                <Typography variant='h4' component='h2' gutterBottom>
                    Clubs
                </Typography>
                <Grid container spacing={2} sx={{ justifyContent: 'space-between' }}>
                    {clubCards}
                </Grid>
            </div>
        );
    }
}

export default ClubCard;

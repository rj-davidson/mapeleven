import * as React from 'react';
import { Card, CardContent, Grid, Typography } from '@mui/material';
import { APITeam } from '../../../Models/api_types';

interface ClubCardState {
  clubs: APITeam[];
}

class ClubCard extends React.Component<{}, ClubCardState> {
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
      });
  }

  render() {
    const clubCards = this.state.clubs.map((team) => (
      <Grid item xs={12} sm={12} md={12} lg={12} key={team.slug}>
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
            <img
              src={'http://localhost:8080/' + team.badge}
              alt={team.name.long + ' badge'}
              width={44}
              height={44}
            />
            <Typography variant='h6' component='h2' sx={{ mr: 2 }}>
              {team.name.long}
            </Typography>
            <Typography variant='body1' color='text.secondary' sx={{ mr: 2 }}>
              Founded: {team.founded}
            </Typography>
            <Typography variant='body2' color='text.secondary'>
              Country: {team.country?.name}
            </Typography>
          </CardContent>
        </Card>
      </Grid>
    ));

    return (
      <div>
        <Typography variant='h4' component='h2' gutterBottom>
          Clubs
        </Typography>
        <Grid container spacing={3} sx={{ justifyContent: 'space-between ' }}>
          {clubCards}
        </Grid>
      </div>
    );
  }
}

export default ClubCard;

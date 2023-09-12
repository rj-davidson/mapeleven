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
      .then(data => this.setState({ clubs: data }));
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
              justifyContent: 'center',
              padding: '12px',
            }}
          >
            <img
              src={'http://localhost:8080/' + team.badge}
              alt={team.name.long + ' badge'}
              width={32}
              height={32}
              style={{ borderRadius: '8px', marginRight:'0px', border: '2px solid #D8DEE9' }}
            />
            <Typography variant='h6' component='h2' sx={{ mr: 0 }}>
              {team.name.long}
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
        <Grid container spacing={2} sx={{ justifyContent: 'space-between' }}>
          {clubCards}
        </Grid>
      </div>
    );
  }
}

export default ClubCard;

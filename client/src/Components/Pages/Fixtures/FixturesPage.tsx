import * as React from 'react';
import '../../../App.css';
import Layout from '../../Layout/Layout';
import DaySwitcher from './DaySwitcher';
import Fixture from './Fixture';
import FixtureTeam from './FixtureTeam';
import { Container, Grid } from '@mui/material';

function FixturesPage(): JSX.Element {
    const fixtures = [];
    for (let i = 0; i < 6; i++) {
        fixtures.push(
            <Fixture
                key={i}
                homeTeam={
                    <FixtureTeam
                        logo='https://media-3.api-sports.io/football/teams/33.png'
                        name='Home team'
                        score={4}
                    />
                }
                awayTeam={
                    <FixtureTeam
                        logo='https://media-3.api-sports.io/football/teams/33.png'
                        name='Away Name'
                        score={3}
                    />
                }
                fixtureInfo='Premier League'
            />,
        );
    }

    return (
        <div>
            <Layout>
                <DaySwitcher />
                <Container maxWidth='lg'>
                    <Grid container spacing={2}>
                        {fixtures.map(fixture => (
                            <Grid item xs={12} md={12} lg={12} key={fixture.key}>
                                {fixture}
                            </Grid>
                        ))}
                    </Grid>
                </Container>
            </Layout>
        </div>
    );
}

export default FixturesPage;

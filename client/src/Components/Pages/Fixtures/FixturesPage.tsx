import * as React from 'react';
import '../../../App.css';
import Layout from '../../Layout/Layout';
import DaySwitcher from './DaySwitcher';
import Fixture from './Fixture';
import './FixturesPage.css';

function FixturesPage(): JSX.Element {
    return (
        <div>
            <Layout>
                <DaySwitcher />
                <Fixture
                    league='Premier League'
                    homeTeam='Manchester United'
                    homeLogo={'https://media-3.api-sports.io/football/teams/33.png'}
                    awayTeam='Liverpool'
                    awayLogo={'https://media-3.api-sports.io/football/teams/33.png'}
                    matchClock='50:30'
                    homeScore={1}
                    awayScore={0}
                />
                <Fixture
                    league='La Liga'
                    homeTeam='Real Madrid'
                    homeLogo={'https://media-3.api-sports.io/football/teams/33.png'}
                    awayTeam='Barcelona'
                    awayLogo={'https://media-3.api-sports.io/football/teams/33.png'}
                    matchClock='20:00'
                    homeScore={0}
                    awayScore={0}
                />
            </Layout>
        </div>
    );
}

export default FixturesPage;

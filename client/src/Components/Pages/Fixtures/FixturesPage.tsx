import { useState, useEffect } from 'react';
import FixtureInfo from './FixtureInfo';
import Layout from '../../Layout/Layout';
import DaySwitcher from './DaySwitcher';
import { Container, Grid } from '@mui/material';
import Fixture from './Fixture';
import FixtureTeam from './FixtureTeam';
import FixtureHeader from './FixtureHeader';

export default function FixturesPage(): JSX.Element {
    const [liveFixtures, setLiveFixtures] = useState([]);
    const [dateFixtures, setDateFixtures] = useState([]);
    const [loading, setLoading] = useState(true);
    const [selectedDate, setSelectedDate] = useState(new Date().toISOString().slice(0, 10));

    useEffect(() => {
        const fetchFixtures = () => {
            const countries = ['England', 'Germany', 'Italy', 'France', 'Spain'];

            // Fetch live fixtures
            fetch(
                `http://127.0.0.1:8080/fixtures?live=${
                    selectedDate === new Date().toISOString().slice(0, 10) ? 'all' : 'none'
                }`,
            )
                .then(response => response.json())
                .then(data => {
                    setLiveFixtures(data.response);
                });

            // Fetch date fixtures
            fetch(`http://127.0.0.1:8080/fixtures?date=${selectedDate}`)
                .then(response => response.json())
                .then(data => {
                    setDateFixtures(data.response.filter(fixture => countries.indexOf(fixture.league.country) !== -1));
                    setLoading(false);
                });
        };

        setLoading(true);
        fetchFixtures();
    }, [selectedDate]);

    const handleDateChange = (newDate: Date) => {
        setSelectedDate(newDate.toISOString().slice(0, 10));
    };

    return (
        <div style={{ width: '100%', padding: '0', margin: '0' }}>
            <Layout>
                <DaySwitcher onDateChange={handleDateChange} />
                <Grid container spacing={2}>
                    {liveFixtures.slice(0, 4).map((fixture, index) => (
                        <Grid item xs={12} md={12} lg={6} key={index}>
                            <Fixture
                                homeTeam={
                                    <FixtureTeam
                                        logo={fixture.teams.home.logo}
                                        name={fixture.teams.home.name}
                                        score={fixture.goals.home}
                                    />
                                }
                                awayTeam={
                                    <FixtureTeam
                                        logo={fixture.teams.away.logo}
                                        name={fixture.teams.away.name}
                                        score={fixture.goals.away}
                                    />
                                }
                                fixtureHeader={
                                    <FixtureHeader league={fixture.league.name} flag={fixture.league.flag} />
                                }
                                fixtureInfo={
                                    <FixtureInfo
                                        live={true}
                                        time={`${fixture.fixture.status.elapsed}'`}
                                        kickoff={fixture.fixture.date.slice(11, 16)}
                                        location={fixture.fixture.venue.city}
                                    />
                                }
                            />
                        </Grid>
                    ))}
                    {dateFixtures.slice(0, 10).map((fixture, index) => (
                        <Grid item xs={12} md={12} lg={6} key={index}>
                            <Fixture
                                homeTeam={<FixtureTeam logo={fixture.teams.home.logo} name={fixture.teams.home.name} />}
                                awayTeam={<FixtureTeam logo={fixture.teams.away.logo} name={fixture.teams.away.name} />}
                                fixtureHeader={
                                    <FixtureHeader league={fixture.league.name} flag={fixture.league.flag} />
                                }
                                fixtureInfo={
                                    <FixtureInfo
                                        live={false}
                                        time={fixture.fixture.date.slice(11, 16)}
                                        kickoff={new Date(fixture.fixture.timestamp * 1000).toLocaleTimeString()}
                                        location={fixture.fixture.venue.city}
                                    />
                                }
                            />
                        </Grid>
                    ))}
                </Grid>
            </Layout>
        </div>
    );
}

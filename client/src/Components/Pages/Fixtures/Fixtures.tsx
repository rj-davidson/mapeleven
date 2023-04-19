import * as React from "react";
import '../../../App.css';
import Layout from "../../Layout/Layout";
import DaySwitcher from "./DaySwitcher";
import FixtureLive from "./FixtureLive";
import "./Fixtures.css"

function Fixtures(): JSX.Element {
    return (
        <div>
            <Layout>
                <DaySwitcher />
                <FixtureLive
                    league="Premier League"
                    homeTeam="Manchester United"
                    homeLogo={"https://media-3.api-sports.io/football/teams/33.png"}
                    awayTeam="Liverpool"
                    awayLogo={"https://media-3.api-sports.io/football/teams/33.png"}
                    matchClock="50:30"
                    homeScore={1}
                    awayScore={0}
                />
                <FixtureLive
                    league="La Liga"
                    homeTeam="Real Madrid"
                    homeLogo={"https://media-3.api-sports.io/football/teams/33.png"}
                    awayTeam="Barcelona"
                    awayLogo={"https://media-3.api-sports.io/football/teams/33.png"}
                    matchClock="20:00"
                    homeScore={0}
                    awayScore={0}
                />
                <h1 className='fixtures'>FIXTURES</h1>;
            </Layout>
        </div>
    );
}

export default Fixtures;

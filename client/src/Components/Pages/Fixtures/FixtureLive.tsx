import * as React from 'react';
import './FixtureLive.css';

type FixtureLiveProps = {
    league: string;
    homeTeam: string;
    homeLogo: string;
    awayTeam: string;
    awayLogo: string;
    matchClock: string;
    homeScore: number;
    awayScore: number;
};

const FixtureLive: React.FC<FixtureLiveProps> = ({
                                                     league,
                                                     homeTeam,
                                                     homeLogo,
                                                     awayTeam,
                                                     awayLogo,
                                                     matchClock,
                                                     homeScore,
                                                     awayScore,
                                                 }) => {
    return (
        <div className="fixture-live">
            <div className="row">
                <div className="col">{league}</div>
                <div className="col">{matchClock}</div>
            </div>
            <div className="row">
                <img src={homeLogo} alt={homeTeam} className={"team-logo"} />
                <div className="col">{homeTeam}</div>
                <div className="col">{homeScore}</div>
            </div>
            <div className="row">
                <img src={awayLogo} alt={awayTeam} className={"team-logo"} />
                <div className="col">{awayTeam}</div>
                <div className="col">{awayScore}</div>
            </div>
        </div>
    );
};

export default FixtureLive;
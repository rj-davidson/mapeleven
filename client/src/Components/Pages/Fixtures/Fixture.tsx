import { MultiSlotComponentProps } from 'react';
import './Fixture.css';

type FixtureProps = {
    homeTeam: React.ReactNode;
    awayTeam: React.ReactNode;
    fixtureInfo: React.ReactNode;
};
const Fixture: React.FC<MultiSlotComponentProps> = ({ homeTeam, awayTeam, fixtureInfo }) => {
    return (
        <div className='fixture-wrapper'>
            <div className='teamsCol'>
                <div className='homeTeam'>{homeTeam}</div>
                <div className='awayTeam'>{awayTeam}</div>
            </div>
            <div className='infoCol'>{fixtureInfo}</div>
        </div>
    );
};

export default Fixture;

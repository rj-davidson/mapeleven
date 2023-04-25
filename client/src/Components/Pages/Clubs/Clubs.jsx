import React from 'react';
import '../../../App.css';
import Layout from '../../Layout/Layout.tsx';
import DisplayClubs from './DisplayClubs.jsx';
import ClubCard from './ClubCard.tsx';
import PassNetwork from './passnetwork.png';

function Clubs() {
    return (
        <div style={{ width: '100%', padding: '0', margin: '0' }}>
            <ClubCard />
            <div style={{ textAlign: 'center'}}>
                <img src={PassNetwork} alt="pass network" style={{ marginTop: '50px' }}/>
            </div>
        </div>
    );
}

export default Clubs;

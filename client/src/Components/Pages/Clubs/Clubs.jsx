import React from 'react';
import '../../../App.css';
import Layout from '../../Layout/Layout.tsx';
import DisplayClubs from './DisplayClubs.jsx';
import ClubCard from './ClubCard.tsx';

function Clubs() {
    return (
        <div style={{ width: '100%', padding: '0', margin: '0' }}>
            <ClubCard />
        </div>
    );
}

export default Clubs;

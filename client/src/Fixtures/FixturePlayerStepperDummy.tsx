import * as React from 'react';
import { useState, useEffect } from 'react';
import { APIFixtureEvent } from '../Models/api_types';
import FixturePlayerCard from './FixturePlayerCard';
import { Grid, Stepper } from '@mui/material';
import {Flex} from "../Util/Flex";
import {Tile} from "../Util/Tile";

interface FixturePlayerStepperDummyProps {
    events: APIFixtureEvent[];
}

function FixturePlayerStepperDummy() {
    const [currentPlayerIndex, setCurrentPlayerIndex] = useState(0);

    // Extract an array of players who scored from the filtered events
    const scoringPlayers = [
        { id: 1, name: 'Player 1', slug: 'l-messi-9t67nlxg' },
        { id: 2, name: 'Player 2', slug: 'e-haaland-mow0zrj9' },
        { id: 3, name: 'Player 3', slug: 'k-mbappe-4gw1b5rw' },
        // Add more players as needed
    ];

    useEffect(() => {
        const interval = setInterval(() => {
            // Increment the player index, and loop back to 0 if it exceeds the array length
            setCurrentPlayerIndex((prevIndex) => (prevIndex + 1) % scoringPlayers.length);
        }, 8000); // Change the interval time (in milliseconds) as needed

        // Clear the interval when the component is unmounted
        return () => clearInterval(interval);
    }, [scoringPlayers.length]);

    const settings = {
        infinite: true,
        speed: 1, // Adjust the speed as needed
        slidesToShow: 1,
        slidesToScroll: 1,
    };

    return (
        <Flex>
            <Tile>
                <h2>Scoring Players:</h2>
                <FixturePlayerCard slug={scoringPlayers[currentPlayerIndex]?.slug} />
            </Tile>
        </Flex>
    );
}

export default FixturePlayerStepperDummy;

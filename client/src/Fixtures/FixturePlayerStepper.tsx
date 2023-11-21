import * as React from 'react';
import { useState, useEffect } from 'react';
import {APIFixtureEvent} from "../Models/api_types";
import FixturePlayerCard from "./FixturePlayerCard";

interface FixturePlayerStepperProps {
    events: APIFixtureEvent[];
}

const FixturePlayerStepper: React.FC<FixturePlayerStepperProps> = ({ events }) => {
    const [currentPlayerIndex, setCurrentPlayerIndex] = useState(0);

    // Filter events to get only those with type 'Goal'
    const goalEvents = events.filter((events) => events.type === 'Goal');



    const scoringPlayers = goalEvents.map((event) => event.player);


    useEffect(() => {
        const interval = setInterval(() => {
            // Increment the player index, and loop back to 0 if it exceeds the array length
            setCurrentPlayerIndex((prevIndex) => (prevIndex + 1) % scoringPlayers.length);
        }, 3000); // Change the interval time (in milliseconds) as needed

        // Clear the interval when the component is unmounted
        return () => clearInterval(interval);
    }, [scoringPlayers.length]);

    const settings = {
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: 1,
        slidesToScroll: 1,
    };

    return (
        <div>
            <h2>Scoring Player Stepper:</h2>
            <p>Current Player: {scoringPlayers[currentPlayerIndex]?.name}</p>
            <FixturePlayerCard slug={scoringPlayers[currentPlayerIndex]?.slug} />
        </div>
    );
};

export default FixturePlayerStepper;

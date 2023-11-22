import * as React from 'react';
import { useState, useEffect } from 'react';
import {APIFixtureEvent} from "../Models/api_types";
import FixturePlayerCard from "./FixturePlayerCard";
import Typography from "@mui/material/Typography";
import { Grid } from '@mui/material';
import {Tile} from "../Util/TileTS";
import {Flex} from "../Util/Flex";

interface FixturePlayerStepperHomeProps {
    events: APIFixtureEvent[];
    homeTeamName: string;
}

const FixturePlayerStepperHome: React.FC<FixturePlayerStepperHomeProps> = ({ events, homeTeamName}) => {
    const [currentPlayerIndex, setCurrentPlayerIndex] = useState(0);

    // Filter events to get only those with type 'Goal'
    const goalEvents = events.filter((events) => events.type === 'Goal' && events.team.name === homeTeamName);


    const timeScored = goalEvents.map((event) => event.time);
    const scoringPlayers = goalEvents.map((event) => event.player);


    useEffect(() => {
        const interval = setInterval(() => {
            // Increment the player index, and loop back to 0 if it exceeds the array length
            setCurrentPlayerIndex((prevIndex) => (prevIndex + 1) % scoringPlayers.length);
        }, 6000); // Change the interval time (in milliseconds) as needed

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
        <Flex style={{flexDirection:'column', gap:'8px'}}>
            {timeScored[currentPlayerIndex] && scoringPlayers[currentPlayerIndex] &&
                <Flex style={{flexDirection:'column', gap:'8px'}}>
                    <Tile >
                        <Typography align="center" sx={{fontSize: '20px'}}>
                            Goal Scored in {timeScored[currentPlayerIndex].elapsed}th minute:
                        </Typography>
                    </Tile>
                    <FixturePlayerCard slug={scoringPlayers[currentPlayerIndex]?.slug} />
                </Flex>
            }
        </Flex>

    );
};

export default FixturePlayerStepperHome;

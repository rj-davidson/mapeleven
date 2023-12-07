import { APIFixtureEvent } from "../../Models/api_types";
import React from "react";
import Typography from "@mui/material/Typography";
import { Link } from "react-router-dom";
import SportsSoccer from "@mui/icons-material/SportsSoccer";
import Box from "@mui/material/Box";
import { EventTime } from "./EventTime";

type MissedPenaltyEventProps = {
    event: APIFixtureEvent;
    away: boolean;
};

export const MissedPenaltyEvent: React.FC<MissedPenaltyEventProps> = ({ event, away }) => {
    return (
        <div style={{ textAlign: away ? "right" : "left" }}>
            <EventTime event={event} />
            <Link to={`/players/${event.player.slug}`}>
                <Typography
                    sx={{
                        "&:hover": {
                            textDecoration: "underline",
                        },
                        fontSize: "20px",
                    }}
                >
                    <Box gap={1} display="inline-flex" alignItems="center">
                        {away ? (
                            <>
                                {event.player.name} <SportsSoccer style={{ color: 'red' }}/>
                            </>
                        ) : (
                            <>
                                <SportsSoccer style={{ color: 'red' }}/> {event.player.name}
                            </>
                        )}
                    </Box>
                </Typography>
            </Link>
            {event.detail && (
                <Box>
                    <Typography color={"text.secondary"} fontSize={"16px"}>
                        <Box gap={1} display="inline-flex" alignItems="center">
                            {event.detail}
                        </Box>
                    </Typography>
                </Box>
            )}
        </div>
    );
};

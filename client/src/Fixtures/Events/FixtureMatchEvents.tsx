import Grid from "@mui/material/Grid";
import Divider from "@mui/material/Divider";
import { type APIFixtureEvent } from "../../Models/api_types";
import React from "react";
import { GoalEvent } from "./GoalEvent";
import { CardEvent } from "./CardEvent";
import { SubstEvent } from "./SubstEvent";

type MatchEventProps = {
  event: APIFixtureEvent;
  awayTeamSlug: string;
};

const FixtureMatchEvent: React.FC<MatchEventProps> = ({
  event,
  awayTeamSlug,
}) => {
  return (
    <Grid container alignItems="center" spacing={2} paddingY={"4px"}>
      <Grid item xs={12} sx={{ paddingLeft: "16px", paddingRight: "16px" }}>
        {event.type === "subst" && (
          <SubstEvent event={event} away={awayTeamSlug === event.team.slug} />
        )}
        {event.type === "Goal" && (
          <GoalEvent event={event} away={event.team.slug === awayTeamSlug} />
        )}

        {event.type === "Card" && (
          <CardEvent event={event} away={awayTeamSlug === event.team.slug} />
        )}
      </Grid>
      <Grid item xs={12}>
        <Divider />
      </Grid>
    </Grid>
  );
};

export default FixtureMatchEvent;

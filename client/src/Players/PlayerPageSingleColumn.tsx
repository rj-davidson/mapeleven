import * as React from "react";
import Grid from "@mui/material/Grid";
import PlayerIDCard from "./PlayerIDCard";
import { APIPlayer } from "../Models/api_types";
import PlayerStats from "./PlayerStats";
import PlayerStatsShooting from "./PlayerStatsShooting";
import PlayerStatsTechnical from "./PlayerStatsTechnical";
import PlayerStatsDefense from "./PlayerStatsDefense";
import PlayerStatsGoalkeeper from "./PlayerStatsGoalkeeper";

interface PlayerPageSingleColumnProps {
  slug: string;
  playerData: APIPlayer | undefined;
}

const PlayerPageSingleColumn: React.FC<PlayerPageSingleColumnProps> = ({
  slug,
  playerData,
}) => {
  if (playerData?.statistics?.[0].games.position === "Goalkeeper") {
    return (
      <Grid container spacing={2}>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerIDCard slug={slug} playerData={playerData} />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerStats playerData={playerData} />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerStatsGoalkeeper playerData={playerData} />
        </Grid>
      </Grid>
    );
  }
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerIDCard slug={slug} playerData={playerData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerStats playerData={playerData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerStatsShooting playerData={playerData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerStatsTechnical playerData={playerData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerStatsDefense playerData={playerData} />
      </Grid>
    </Grid>
  );
};

export default PlayerPageSingleColumn;

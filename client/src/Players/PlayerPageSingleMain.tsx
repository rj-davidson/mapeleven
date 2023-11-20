import * as React from "react";
import Grid from "@mui/material/Grid";
import PlayerRadarChart from "./PlayerRadarChart";
import { APIPlayer } from "../Models/api_types";
import PlayerStatsShootingPieChart from "./PlayerStatsShootingPieChart";
import PlayerStatsTechnicalPieChart from "./PlayerStatsTechnicalPieChart";
import PlayerStatsPassingPieChart from "./PlayerStatsPassingPieChart";
import PlayerStatsDefensePieChart from "./PlayerStatsDefensePieChart";

interface PlayerPageSingleMainProps {
  playerData: APIPlayer | undefined;
}

const PlayerPageSingleMain: React.FC<PlayerPageSingleMainProps> = ({
  playerData,
}) => {
  const emptyData: APIPlayer = {
    age: 0,
    firstname: "",
    height: "",
    injured: false,
    lastname: "",
    name: "",
    photo: "",
    slug: "",
    weight: "",
  };
  if (playerData?.statistics?.[0].games.position === "Goalkeeper") {
    return (
      <Grid container spacing={2}>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <PlayerRadarChart playerData={playerData} />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
          <PlayerStatsDefensePieChart playerData={playerData} />
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
          <PlayerStatsPassingPieChart playerData={playerData} />
        </Grid>
      </Grid>
    );
  }
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerRadarChart playerData={playerData || emptyData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
        <PlayerStatsShootingPieChart playerData={playerData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
        <PlayerStatsPassingPieChart playerData={playerData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
        <PlayerStatsTechnicalPieChart playerData={playerData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
        <PlayerStatsDefensePieChart playerData={playerData} />
      </Grid>
    </Grid>
  );
};

export default PlayerPageSingleMain;

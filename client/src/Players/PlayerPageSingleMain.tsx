import * as React from "react";
import Grid from "@mui/material/Grid";
import PlayerRadarChart from "./PlayerRadarChart";
import { APIPlayer } from "../Models/api_types";

interface PlayerPageSingleMainProps {
  playerData: APIPlayer | undefined;
}

const PlayerPageSingleMain: React.FC<PlayerPageSingleMainProps> = ({
  playerData,
}) => {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerRadarChart playerData={playerData} />
      </Grid>
    </Grid>
  );
};

export default PlayerPageSingleMain;

import * as React from "react";
import Grid from "@mui/material/Grid";
import PlayerIDCard from "./PlayerIDCard";
import { APIPlayer } from "../Models/api_types";

interface PlayerPageSingleColumnProps {
  slug: string;
  playerData: APIPlayer | undefined;
}

const PlayerPageSingleColumn: React.FC<PlayerPageSingleColumnProps> = ({
  slug,
  playerData,
}) => {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <PlayerIDCard slug={slug} playerData={playerData} />
      </Grid>
    </Grid>
  );
};

export default PlayerPageSingleColumn;

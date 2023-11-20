import React from "react";
import "../App.css";
import Grid from "@mui/material/Grid";
import { Tile } from "../Util/TileTS";
import ListTeams from "./ListTeams.tsx";

function TeamPageAll() {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile>
          <ListTeams skeleton={5} limit={20} />
        </Tile>
      </Grid>
    </Grid>
  );
}

export default TeamPageAll;

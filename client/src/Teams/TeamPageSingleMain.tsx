import * as React from "react";
import Grid from "@mui/material/Grid";
import { APITSGoalMinuteSplit } from "../Models/api_types";
import TeamGoalSplitCharts from "./TeamGoalSplitCharts";
import TeamRadarChart from "./TeamRadarChart";
import { Tile } from "../Util/TileTS";
import Field from "./FieldSVG";

interface TeamPageSingleMainProps {
  goalMinuteSplit: APITSGoalMinuteSplit;
  name: string;
  goals: number;
  clean: number;
  fixtures: number;
  wins: number;
  averageGoals: number;
  gamesScored: number;
  failedToScore: number;
}

const TeamPageSingleMain: React.FC<TeamPageSingleMainProps> = ({
  goalMinuteSplit,
  name,
  goals,
  clean,
  fixtures,
  averageGoals,
  gamesScored,
  failedToScore,
  wins,
}) => {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <TeamRadarChart
          name={name}
          goals={goals}
          clean={clean}
          fixtures={fixtures}
          wins={wins}
          averageGoals={averageGoals}
          gamesScored={gamesScored}
          failedToScore={failedToScore}
        />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <TeamGoalSplitCharts goalMinuteSplit={goalMinuteSplit} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile sx={{ height: "385px" }}>
          <Field />
        </Tile>
      </Grid>
    </Grid>
  );
};

export default TeamPageSingleMain;

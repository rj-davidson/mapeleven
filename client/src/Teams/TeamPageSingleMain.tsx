import * as React from "react";
import Grid from "@mui/material/Grid";
import { APITeam, APITSGoalMinuteSplit } from "../Models/api_types";
import TeamGoalSplitCharts from "./TeamGoalSplitCharts";
import TeamRadarChart from "./TeamRadarChart";
import TeamStatsWinsPieChart from "./TeamStatsWinsPieChart";
import TeamStatsGoalsPieChart from "./TeamStatsGoalsPieChart";
import TeamStatsScoredPieChart from "./TeamStatsScoredPieChart";

interface TeamPageSingleMainProps {
  goalMinuteSplit: APITSGoalMinuteSplit | undefined;
  teamData: APITeam | undefined;
}

const TeamPageSingleMain: React.FC<TeamPageSingleMainProps> = ({
  goalMinuteSplit,
  teamData,
}) => {
  const emptyData: APITeam = {
    badge: "",
    leagueName: "",
    logo: "",
    name: { long: "", short: "" },
    slug: "",
  };
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <TeamRadarChart teamData={teamData || emptyData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <TeamGoalSplitCharts goalMinuteSplit={goalMinuteSplit} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
        <TeamStatsWinsPieChart teamData={teamData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
        <TeamStatsGoalsPieChart teamData={teamData} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={6} width="100%">
        <TeamStatsScoredPieChart teamData={teamData} />
      </Grid>
    </Grid>
  );
};

export default TeamPageSingleMain;

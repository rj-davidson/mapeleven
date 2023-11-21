import * as React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Grid from "@mui/material/Grid";
import Skeleton from "@mui/material/Skeleton";
import { APITeam } from "../Models/api_types";
import TeamPageSingleColumn from "./TeamPageSingleColumn";
import TeamPageSingleMain from "./TeamPageSingleMain";

const url = import.meta.env.VITE_API_URL;

function TeamPageSingle() {
  const { slug } = useParams();
  const [teamData, setTeamData] = useState<APITeam>();

  const [loading, setLoading] = useState<boolean>(true);
  const [leagueSlug, setLeagueSlug] = useState<string>("");
  const [teamRank, setTeamRank] = useState<string>("");

  useEffect(() => {
    fetch(`${url}/teams/${slug}`)
      .then((response) => response.json())
      .then((jsonData) => {
        setTeamData(jsonData as APITeam);
        setLeagueSlug(jsonData.competitions[0].league.slug);
      })
      .catch((error) => console.error(error));
    fetch(`${url}/leagues/${leagueSlug}`)
      .then((response) => response.json())
      .then((jsonData) => {
        const standings = jsonData.standings;

        if (standings) {
          // Find the team with the matching slug
          const team = standings.find((team) => team.team.slug === slug);

          // If the team is found, set its rank as the value
          if (team) {
            setTeamRank(getRank(team.rank));
          }
          setLoading(false);
        }
      })
      .catch((error) => console.error(error));
  }, [slug, leagueSlug]);

  const getRank = (rank: number): string => {
    if (rank == 1) {
      return `${rank}st`;
    }
    if (rank == 2) {
      return `${rank}nd`;
    }
    if (rank == 3) {
      return `${rank}rd`;
    }
    return `${rank}th`;
  };

  return loading ? (
    <Grid container spacing={2} direction="column">
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Skeleton
          variant="rectangular"
          height="150px"
          width="100%"
          sx={{ background: "gray", borderRadius: "12px" }}
        />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Skeleton
          variant="rectangular"
          height="600px"
          width="100%"
          sx={{ background: "gray", borderRadius: "12px" }}
        />
      </Grid>
    </Grid>
  ) : (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
        <TeamPageSingleColumn slug={slug} teamData={teamData} rank={teamRank} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
        <TeamPageSingleMain
          goalMinuteSplit={teamData?.competitions?.[0].stats?.goals.for.minute}
          teamData={teamData}
        />
      </Grid>
    </Grid>
  );
}

export default TeamPageSingle;

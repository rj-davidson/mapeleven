import * as React from "react";
import Grid from "@mui/material/Grid";
import { Flex } from "../Util/Flex.jsx";
import FormCard from "./FormCard.jsx";
import TeamIDCard from "./TeamIDCard";
import TeamStats from "./TeamStats";
import { APITeam } from "../Models/api_types";

interface TeamPageSingleColumnProps {
  slug: string | undefined;
  teamData: APITeam | undefined;
  rank: string;
}

const TeamPageSingleColumn: React.FC<TeamPageSingleColumnProps> = ({
  slug,
  teamData,
  rank,
}) => {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <TeamIDCard slug={slug} teamData={teamData} rank={rank} />
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Flex
          style={{
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "column",
          }}
        >
          <FormCard />
        </Flex>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Flex
          style={{
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "column",
          }}
        >
          <TeamStats teamData={teamData} />
        </Flex>
      </Grid>
    </Grid>
  );
};

export default TeamPageSingleColumn;

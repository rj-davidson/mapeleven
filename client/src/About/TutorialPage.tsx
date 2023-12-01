import * as React from "react";
import Typography from "@mui/material/Typography";
import "../App.css";
import Grid from "@mui/material/Grid";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";

function TutorialPage() {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Typography
            sx={{
              fontSize: { xs: "32px", sm: "32px" },
              marginLeft: "10px",
              fontWeight: "bold",
            }}
          >
            Tutorial
          </Typography>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Flex style={{ flexDirection: "column" }}>
            <Typography
              sx={{
                fontSize: { xs: "24px", sm: "24px" },
                marginLeft: "10px",
                fontWeight: "bold",
              }}
            >
              Home Page
            </Typography>
            <Typography
              sx={{ fontSize: { xs: "20px", sm: "20px" }, marginLeft: "10px" }}
            >
              Use the hamburger menu in the top left to find access to other
              pages on the site. Use the search bar in the top right to find
              players, teams, or leagues. The homepage contains popular players,
              teams, and a list of matches/fixtures. Click on a player, team, or
              match to go to its individual page. These matches are sorted by
              the data. Use the arrows to step forward or back one day. Click on
              the date to open a calendar.
            </Typography>
          </Flex>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Flex style={{ flexDirection: "column" }}>
            <Typography
              sx={{
                fontSize: { xs: "24px", sm: "24px" },
                marginLeft: "10px",
                fontWeight: "bold",
              }}
            >
              Matches/Fixtures Page
            </Typography>
            <Typography
              sx={{ fontSize: { xs: "20px", sm: "20px" }, marginLeft: "10px" }}
            >
              This page will show you the score, team lineups, goal scorers, and
              more for a specific match. Click on a player's image or name to go
              that player's page. Click on a team's name to go to that teams
              page. Scroll down to see the full list of match events.
            </Typography>
          </Flex>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Flex style={{ flexDirection: "column" }}>
            <Typography
              sx={{
                fontSize: { xs: "24px", sm: "24px" },
                marginLeft: "10px",
                fontWeight: "bold",
              }}
            >
              Player Page
            </Typography>
            <Typography
              sx={{ fontSize: { xs: "20px", sm: "20px" }, marginLeft: "10px" }}
            >
              On a player page, you can view and compare player stats. The
              player's ID card on the left displays basic player information.
              The radar comparison chart on the right will graph a radar chart
              for this current player. You can compare this player's stats to
              another player by using the search bar to the right of the radar
              chart and selecting another player. Use the selectors on the left
              of the radar chart to select different data points to graph. You
              can select up to 6, with a minimum of 3 data points. Scroll down
              and you will find more stats and visualizations.
            </Typography>
          </Flex>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Flex style={{ flexDirection: "column" }}>
            <Typography
              sx={{
                fontSize: { xs: "24px", sm: "24px" },
                marginLeft: "10px",
                fontWeight: "bold",
              }}
            >
              Team Page
            </Typography>
            <Typography
              sx={{ fontSize: { xs: "20px", sm: "20px" }, marginLeft: "10px" }}
            >
              On a team page, you can view and compare player stats. The team's
              ID card on the left displays basic team information. The radar
              comparison chart on the right will graph a radar chart for this
              current team. You can compare this team's stats to another team by
              using the search bar to the right of the radar chart and selecting
              another team. Use the selectors on the left of the radar chart to
              select different data points to graph. You can select up to 6,
              with a minimum of 3 data points. Scroll down and you will find
              more stats and visualizations plus a list players on the team.
            </Typography>
          </Flex>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Flex style={{ flexDirection: "column" }}>
            <Typography
              sx={{
                fontSize: { xs: "24px", sm: "24px" },
                marginLeft: "10px",
                fontWeight: "bold",
              }}
            >
              League Page
            </Typography>
            <Typography
              sx={{ fontSize: { xs: "20px", sm: "20px" }, marginLeft: "10px" }}
            >
              On a league page, you will find the current standings for that
              league. This will include each teams recent form, total points,
              games won, games lost, and more. Select a team's name to go to
              that team's page.
            </Typography>
          </Flex>
        </Tile>
      </Grid>
    </Grid>
  );
}

export default TutorialPage;

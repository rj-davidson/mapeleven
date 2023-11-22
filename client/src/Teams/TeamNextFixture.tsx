import * as React from "react";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import { APITeam } from "../Models/api_types";
import Typography from "@mui/material/Typography";
import DisplayImage from "../Util/DisplayImage";
import { Link } from "react-router-dom";

interface TeamNextFixtureProps {
  teamData: APITeam | undefined;
}

const TeamNextFixture: React.FC<TeamNextFixtureProps> = ({ teamData }) => {
  let fixtureSlug = "";
  let oppName = "";
  let oppPhoto = "";
  let oppSlug = "";
  if (teamData && teamData.competitions?.[0].fixtures) {
    const fixtures = teamData.competitions[0].fixtures;
    // Iterate through fixtures to find the first fixture with status "not started"
    for (let i = 0; i < fixtures.length; i++) {
      const fixture = fixtures[i];

      if (fixture.status === "Not Started") {
        fixtureSlug = fixture.slug;
        oppName = fixture.awayTeam.name.long;
        oppPhoto = fixture.awayTeam.badge;
        oppSlug = fixture.awayTeam.slug;
        break;
      }
    }
  }

  return (
    <Link to={`/fixtures/${fixtureSlug}`}>
      <Tile sx={{ alignItems: "center", justifyContent: "space-between" }}>
        <Flex style={{ flexDirection: "column" }}>
          <Typography
            sx={{
              fontSize: "20px",
              textAlign: "left",
              whiteSpace: "nowrap",
            }}
          >
            Next Match
          </Typography>
          <Link to={`/teams/${oppSlug}`}>
            <Typography
              sx={{
                fontSize: "24px",
                textAlign: "left",
                whiteSpace: "nowrap",
                fontWeight: "bold",
                "&:hover": {
                  textDecoration: "underline", // Add underline on hover
                },
              }}
            >
              {oppName}
            </Typography>
          </Link>
        </Flex>
        <Link to={`/teams/${oppSlug}`}>
          <DisplayImage
            src={oppPhoto}
            alt={oppName}
            width={"64px"}
            height={"64px"}
          />
        </Link>
      </Tile>
    </Link>
  );
};

export default TeamNextFixture;

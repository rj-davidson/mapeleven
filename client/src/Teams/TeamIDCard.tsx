import * as React from "react";
import { useState } from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import Tilt from "react-parallax-tilt";
import DisplayImage from "../Util/DisplayImage";
import TeamIDTitle from "./TeamIDTitle";
import TeamIDStatBox from "./TeamIDStatBox";
import { useMediaQuery } from "@mui/material";
import { APITeam } from "../Models/api_types";

interface TeamIDCardProps {
  slug: string | undefined;
  teamData: APITeam | undefined;
  rank: string;
}

const TeamIDCard: React.FC<TeamIDCardProps> = ({ slug, teamData, rank }) => {
  const [isHovered, setIsHovered] = useState<boolean>(false);
  const [inspect, setInspect] = useState<boolean>(false);

  const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");

  return (
    <Box>
      {!isSmallerThanLG ? (
        <Box
          onMouseEnter={() => setIsHovered(true)}
          onMouseLeave={() => setIsHovered(false)}
          onClick={() => setInspect(!inspect)}
        >
          <Tilt
            reset={inspect}
            glareEnable={true}
            tiltMaxAngleX={10}
            tiltMaxAngleY={10}
            tiltAngleXManual={inspect ? null : 0}
            tiltAngleYManual={inspect ? null : 0}
            perspective={1000}
            glareMaxOpacity={0.2}
            glareBorderRadius={"10px"}
            tiltReverse={true}
            style={{
              cursor: isHovered ? "pointer" : "default",
              borderRadius: "10px",
              boxShadow: isHovered ? "rgba(0,0,0,0.2) 0 0 10px" : "none",
            }}
          >
            <Tile
              sx={{
                flexDirection: "column",
                gap: "32px",
              }}
            >
              <TeamIDTitle
                slug={slug}
                name={teamData?.name.long}
                country={teamData?.country?.name}
                flag={teamData?.country?.flag}
                leagueSlug={teamData?.competitions?.[0].league.slug}
                leagueName={teamData?.competitions?.[0].league.name}
                leagueLogo={teamData?.competitions?.[0].league.logo}
              />

              <Flex
                style={{
                  justifyContent: "center",
                  alignItems: "center",
                  marginTop: "-12px",
                }}
              >
                <DisplayImage
                  src={teamData?.badge}
                  alt={teamData?.name.long}
                  width={"108px"}
                  height={"108px"}
                />
              </Flex>

              <Grid container spacing={1}>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox stat="Position" value={rank} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox
                    stat="Points"
                    value={
                      3 *
                        ((teamData?.competitions?.[0].stats?.fixtures.wins
                          .home || 0) +
                          (teamData?.competitions?.[0].stats?.fixtures.wins
                            .away || 0)) +
                      (teamData?.competitions?.[0].stats?.fixtures.draws.home ||
                        0) +
                      (teamData?.competitions?.[0].stats?.fixtures.draws.away ||
                        0)
                    }
                  />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox
                    stat="Games Played"
                    value={
                      (teamData?.competitions?.[0].stats?.fixtures.played
                        .home || 0) +
                      (teamData?.competitions?.[0].stats?.fixtures.played
                        .away || 0)
                    }
                  />
                </Grid>
              </Grid>

              <Typography
                sx={{
                  fontSize: "16px",
                  textAlign: "right",
                  fontStyle: "italic",
                }}
              >
                Founded {teamData?.founded}
              </Typography>
            </Tile>
          </Tilt>
        </Box>
      ) : (
        <Tile
          sx={{
            flexDirection: "column",
            gap: "32px",
          }}
        >
          <TeamIDTitle
            slug={slug}
            name={teamData?.name.long}
            country={teamData?.country?.name}
            flag={teamData?.country?.flag}
            leagueSlug={teamData?.competitions?.[0].league.slug}
            leagueName={teamData?.competitions?.[0].league.name}
            leagueLogo={teamData?.competitions?.[0].league.logo}
          />

          <Flex
            style={{
              justifyContent: "center",
              alignItems: "center",
              marginTop: "-12px",
            }}
          >
            <DisplayImage
              src={teamData?.badge}
              alt={teamData?.name.long}
              width={"108px"}
              height={"108px"}
            />
          </Flex>

          <Grid container spacing={1}>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox stat="Position" value={rank} />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox
                stat="Points"
                value={
                  3 *
                    ((teamData?.competitions?.[0].stats?.fixtures.wins.home ||
                      0) +
                      (teamData?.competitions?.[0].stats?.fixtures.wins.away ||
                        0)) +
                  (teamData?.competitions?.[0].stats?.fixtures.draws.home ||
                    0) +
                  (teamData?.competitions?.[0].stats?.fixtures.draws.away || 0)
                }
              />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox
                stat="Games Played"
                value={
                  (teamData?.competitions?.[0].stats?.fixtures.played.home ||
                    0) +
                  (teamData?.competitions?.[0].stats?.fixtures.played.away || 0)
                }
              />
            </Grid>
          </Grid>

          <Typography
            sx={{
              fontSize: "16px",
              textAlign: "right",
              fontStyle: "italic",
            }}
          >
            Founded {teamData?.founded}
          </Typography>
        </Tile>
      )}
    </Box>
  );
};

export default TeamIDCard;

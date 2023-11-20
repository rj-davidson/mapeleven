import * as React from "react";
import { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import Tilt from "react-parallax-tilt";
import DisplayImage from "../Util/DisplayImage";
import TeamIDTitle from "../Teams/TeamIDTitle";
import TeamIDStatBox from "../Teams/TeamIDStatBox";
import { useMediaQuery } from "@mui/material";

const url = import.meta.env.VITE_API_URL;

interface TeamIDCardProps {
  slug: string;
  name: string;
  badge: string;
}

const TeamIDCardFixturePage: React.FC<TeamIDCardProps> = ({
  slug,
  name,
  badge,
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false);
  const [inspect, setInspect] = useState<boolean>(false);

  const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");
  const [country, setCountry] = useState<string>("");
  const [flag, setFlag] = useState<string>("");
  const [founded, setFounded] = useState<string>("");
  const [gamesPlayed, setGamesPlayed] = useState<number>(0);
  const [wins, setWins] = useState<number>(0);
  const [draws, setDraws] = useState<number>(0);
  const [loses, setLoses] = useState<number>(0);

  useEffect(() => {
    fetch(`${url}/teams/${slug}`)
      .then((response) => response.json())
      .then((jsonData) => {
        setFlag(jsonData.country.flag);
        setCountry(jsonData.country.name);
        setFounded(jsonData.founded);
        setGamesPlayed(
          jsonData.competitions[0].stats.fixtures.played.home +
            jsonData.competitions[0].stats.fixtures.played.away
        );
        setWins(
          jsonData.competitions[0].stats.fixtures.wins.home +
            jsonData.competitions[0].stats.fixtures.wins.away
        );
        setDraws(
          jsonData.competitions[0].stats.fixtures.draws.home +
            jsonData.competitions[0].stats.fixtures.draws.away
        );
        setLoses(
          jsonData.competitions[0].stats.fixtures.loses.home +
            jsonData.competitions[0].stats.fixtures.loses.away
        );
      })
      .catch((error) => console.error(error));
  }, [slug]);

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
                gap: "24px",
              }}
            >
              <TeamIDTitle
                slug={slug}
                name={name}
                country={country}
                flag={flag}
              />

              <Flex
                style={{
                  justifyContent: "center",
                  alignItems: "center",
                  marginTop: "-12px",
                }}
              >
                <DisplayImage
                  src={badge}
                  alt={name}
                  width={"108px"}
                  height={"108"}
                />
              </Flex>

              <Grid container spacing={1}>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox stat="Points" value={3 * wins + draws} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox stat="Games Played" value={gamesPlayed} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox stat="Wins" value={wins} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox stat="Draws" value={draws} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <TeamIDStatBox stat="Losses" value={loses} />
                </Grid>
              </Grid>

              <Typography
                sx={{
                  fontSize: "16px",
                  textAlign: "right",
                  fontStyle: "italic",
                }}
              >
                Founded {founded}
              </Typography>
            </Tile>
          </Tilt>
        </Box>
      ) : (
        <Tile
          sx={{
            flexDirection: "column",
            gap: "24px",
          }}
        >
          <TeamIDTitle slug={slug} name={name} country={country} flag={flag} />

          <Flex
            style={{
              justifyContent: "center",
              alignItems: "center",
              marginTop: "-12px",
            }}
          >
            <DisplayImage
              src={badge}
              alt={name}
              width={"108px"}
              height={"108"}
            />
          </Flex>

          <Grid container spacing={1}>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox stat="Points" value={3 * wins + draws} />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox stat="Games Played" value={gamesPlayed} />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox stat="Wins" value={wins} />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox stat="Draws" value={draws} />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <TeamIDStatBox stat="Losses" value={loses} />
            </Grid>
          </Grid>

          <Typography
            sx={{
              fontSize: "16px",
              textAlign: "right",
              fontStyle: "italic",
            }}
          >
            Founded {founded}
          </Typography>
        </Tile>
      )}
    </Box>
  );
};

export default TeamIDCardFixturePage;

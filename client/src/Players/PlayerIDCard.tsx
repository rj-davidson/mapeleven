import * as React from "react";
import { useState } from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import { useMediaQuery } from "@mui/material";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import Tilt from "react-parallax-tilt";
import DisplayImage from "../Util/DisplayImage";
import PlayerIDTitle from "./PlayerIDTitle";
import PlayerIDStatBox from "./PlayerIDStatBox";
import { APIPlayer } from "../Models/api_types";

interface PlayerIDCardProps {
  slug: string;
  playerData: APIPlayer | undefined;
}

const PlayerIDCard: React.FC<PlayerIDCardProps> = ({ slug, playerData }) => {
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
              <PlayerIDTitle slug={slug} playerData={playerData} />

              <Flex
                style={{
                  justifyContent: "center",
                  alignItems: "center",
                  marginTop: "-20px",
                }}
              >
                <DisplayImage
                  src={playerData?.photo}
                  alt={playerData?.name}
                  width={"108px"}
                  height={"108px"}
                  sx={{
                    borderRadius: "10px",
                    border: "5px solid var(--light0)",
                  }}
                />
              </Flex>
              <Grid container spacing={1}>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <PlayerIDStatBox
                    stat="Position"
                    value={playerData?.statistics?.[0].games.position}
                  />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <PlayerIDStatBox stat="Age" value={playerData?.age} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <PlayerIDStatBox stat="Height" value={playerData?.height} />
                </Grid>
                <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
                  <PlayerIDStatBox stat="Weight" value={playerData?.weight} />
                </Grid>
              </Grid>
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
          <PlayerIDTitle slug={slug} playerData={playerData} />

          <Flex
            style={{
              justifyContent: "center",
              alignItems: "center",
              marginTop: "-20px",
            }}
          >
            <DisplayImage
              src={playerData?.photo}
              alt={playerData?.name}
              width={"108px"}
              height={"108px"}
              sx={{
                borderRadius: "10px",
                border: "5px solid var(--light0)",
              }}
            />
          </Flex>
          <Grid container spacing={1}>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <PlayerIDStatBox
                stat="Position"
                value={playerData?.statistics?.[0].games.position}
              />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <PlayerIDStatBox stat="Age" value={playerData?.age} />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <PlayerIDStatBox stat="Height" value={playerData?.height} />
            </Grid>
            <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
              <PlayerIDStatBox stat="Weight" value={playerData?.weight} />
            </Grid>
          </Grid>
        </Tile>
      )}
    </Box>
  );
};

export default PlayerIDCard;

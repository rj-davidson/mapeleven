import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Flex } from "../Util/Flex";
import DisplayImage from "../Util/DisplayImage";
import { useState } from "react";

interface LineupPlayerProps {
  player: any;
  index: any;
  photo: string;
  home: boolean;
  mobile: boolean;
  number: number;
}

const LineupPlayer: React.FC<LineupPlayerProps> = ({
  player,
  index,
  photo,
  home,
  mobile,
  number,
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false);

  function extractLastName(name: string): string {
    const namePattern = /.*\.\s?(.*)/;
    const match = name.match(namePattern);
    return match ? match[1] : name;
  }

  return (
    <Box
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      <Flex
        key={index}
        style={{
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <DisplayImage
          src={photo}
          alt={player.name}
          width={"50px"}
          height={"50px"}
          sx={{
            borderRadius: "50%",
            border: isHovered
              ? "thick solid var(--light0)"
              : home
              ? "thick solid var(--red)"
              : "thick solid var(--blue)",
          }}
        />
        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          paddingTop={1}
        >
          <Typography
            variant="subtitle2"
            style={{
              fontSize: mobile ? 10 : "initial",
              whiteSpace: player.name.includes(" ") ? "normal" : "nowrap",
            }}
            textAlign={"center"}
          >
            <span
              style={{ color: "#A8ADD9", marginRight: "3px" }}
            >{`${number}`}</span>
            {extractLastName(player.name)}
          </Typography>
        </Box>
      </Flex>
    </Box>
  );
};

export default LineupPlayer;

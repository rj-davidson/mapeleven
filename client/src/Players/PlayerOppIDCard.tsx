import * as React from "react";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import DisplayImage from "../Util/DisplayImage";
import PlayerIDTitle from "./PlayerIDTitle";
import { APIPlayer } from "../Models/api_types";

interface TeamOppIDCardProps {
  slug: string | undefined;
  playerData: APIPlayer | undefined;
}

const TeamOppIDCard: React.FC<TeamOppIDCardProps> = ({
  slug = "",
  playerData,
}) => {
  return (
    <Tile
      sx={{
        flexDirection: "column",
        justifyContent: "space-between",
        height: "100%",
        background: "var(--dark1)",
      }}
    >
      <PlayerIDTitle slug={slug} playerData={playerData} />

      <Flex
        style={{
          justifyContent: "center",
          alignItems: "center",
          height: "100%",
        }}
      >
        <DisplayImage
          src={playerData?.photo}
          alt={playerData?.name}
          width={"108px"}
          height={"108px"}
          sx={{
            borderRadius: "10px",
            boxShadow: "rgba(0,0,0,0.2) 0 0 10px",
          }}
        />
      </Flex>
    </Tile>
  );
};

export default TeamOppIDCard;

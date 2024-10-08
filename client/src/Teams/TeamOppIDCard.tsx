import * as React from "react";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import DisplayImage from "../Util/DisplayImage";
import TeamIDTitle from "./TeamIDTitle";

interface TeamOppIDCardProps {
  slug: string | undefined;
  name: string;
  badge: string;
  country: string;
  flag: string;
  leagueSlug: string;
  leagueName: string;
  leagueLogo: string;
}

const TeamOppIDCard: React.FC<TeamOppIDCardProps> = ({
  slug = "",
  name = "Team Name",
  badge = null,
  country = "Country",
  flag = null,
  leagueSlug,
  leagueName,
  leagueLogo,
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
      <TeamIDTitle
        slug={slug}
        name={name}
        country={country}
        flag={flag}
        leagueSlug={leagueSlug}
        leagueName={leagueName}
        leagueLogo={leagueLogo}
      />

      <Flex
        style={{
          justifyContent: "center",
          alignItems: "center",
          height: "100%",
        }}
      >
        <DisplayImage src={badge} alt={name} width={"108px"} height={"108"} />
      </Flex>
    </Tile>
  );
};

export default TeamOppIDCard;

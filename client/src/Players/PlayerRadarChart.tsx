import * as React from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import Typography from "@mui/material/Typography";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import PlayerRadarSearch from "./PlayerRadarSearch.jsx";
import {
  Legend,
  PolarAngleAxis,
  PolarGrid,
  Radar,
  RadarChart,
  ResponsiveContainer,
  Tooltip,
} from "recharts";
import { useEffect, useState } from "react";
import PlayerOppIDCard from "./PlayerOppIDCard";
import { APIPlayer } from "../Models/api_types";

interface PlayerRadarChartProps {
  playerData: APIPlayer;
}

interface ChartDatum {
  stat: string;
  A: number;
  Opponent: number;
}

const url = import.meta.env.VITE_API_URL;

const PlayerRadarChart: React.FC<PlayerRadarChartProps> = ({ playerData }) => {
  const [oppSlug, setOppSlug] = useState<string>();
  const [oppData, setOppData] = useState<APIPlayer>();

  // Array of all available stats
  const allStats = [
    "Goals",
    "Assists",
    "Shots",
    "Shots on Target",
    "Passes",
    "Dribble Success",
    "Duels Won",
    "Key Passes",
    "Pass Accuracy",
    "Saves",
  ];

  const [stat1, setStat1] = useState<string>("Goals");
  const [stat2, setStat2] = useState<string>("Passes");
  const [stat3, setStat3] = useState<string>("Shots");
  const [stat4, setStat4] = useState<string>("None");
  const [stat5, setStat5] = useState<string>("None");
  const [stat6, setStat6] = useState<string>("None");

  const selectedStats = [stat1, stat2, stat3, stat4, stat5, stat6];

  useEffect(() => {
    if (oppSlug) {
      fetch(`${url}/players/${oppSlug}`)
        .then((response) => response.json())
        .then((jsonData) => {
          setOppData(jsonData as APIPlayer);
        })
        .catch((error) => console.error(error));
    }
  }, [oppSlug]);

  const data: ChartDatum[] = [];

  const getStatValue = (stat: string, opp: boolean): number => {
    if (opp) {
      if (stat === "Goals") {
        return (oppData?.statistics?.[0].shooting.goals ?? 0) * 2;
      }
      if (stat === "Assists") {
        return (oppData?.statistics?.[0].shooting.assists ?? 0) * 2;
      }
      if (stat === "Shots") {
        return (oppData?.statistics?.[0].shooting.shots ?? 0) * 2;
      }
      if (stat === "Shots on Target") {
        return (oppData?.statistics?.[0].shooting.shots_on_target ?? 0) * 2;
      }
      if (stat === "Passes") {
        return (oppData?.statistics?.[0].technical.passes ?? 0) * 0.1;
      }
      if (stat === "Dribble Success") {
        return (oppData?.statistics?.[0].technical.dribbles_success ?? 0) * 2;
      }
      if (stat === "Duels Won") {
        return (oppData?.statistics?.[0].defense.duels_won ?? 0) * 2;
      }
      if (stat === "Key Passes") {
        return (oppData?.statistics?.[0].technical.key_passes ?? 0) * 2;
      }
      if (stat === "Pass Accuracy") {
        return (oppData?.statistics?.[0].technical.pass_accuracy ?? 0) * 2;
      }
      if (stat === "Saves") {
        return (oppData?.statistics?.[0].shooting.saves ?? 0) * 2;
      }
    }

    if (stat === "Goals") {
      return (playerData.statistics?.[0].shooting.goals ?? 0) * 2;
    }
    if (stat === "Assists") {
      return (playerData.statistics?.[0].shooting.assists ?? 0) * 2;
    }
    if (stat === "Shots") {
      return (playerData.statistics?.[0].shooting.shots ?? 0) * 2;
    }
    if (stat === "Shots on Target") {
      return (playerData.statistics?.[0].shooting.shots_on_target ?? 0) * 2;
    }
    if (stat === "Passes") {
      return (playerData.statistics?.[0].technical.passes ?? 0) * 0.1;
    }
    if (stat === "Dribble Success") {
      return (playerData.statistics?.[0].technical.dribbles_success ?? 0) * 2;
    }
    if (stat === "Duels Won") {
      return (playerData.statistics?.[0].defense.duels_won ?? 0) * 2;
    }
    if (stat === "Key Passes") {
      return (playerData.statistics?.[0].technical.key_passes ?? 0) * 2;
    }
    if (stat === "Pass Accuracy") {
      return (playerData.statistics?.[0].technical.pass_accuracy ?? 0) * 2;
    }
    if (stat === "Saves") {
      return (playerData.statistics?.[0].shooting.saves ?? 0) * 2;
    }

    return 0;
  };

  if (stat1 !== "None") {
    data.push({
      stat: stat1,
      A: getStatValue(stat1, false),
      Opponent: getStatValue(stat1, true),
    });
  }

  if (stat2 !== "None") {
    data.push({
      stat: stat2,
      A: getStatValue(stat2, false),
      Opponent: getStatValue(stat2, true),
    });
  }

  if (stat3 !== "None") {
    data.push({
      stat: stat3,
      A: getStatValue(stat3, false),
      Opponent: getStatValue(stat3, true),
    });
  }

  if (stat4 !== "None") {
    data.push({
      stat: stat4,
      A: getStatValue(stat4, false),
      Opponent: getStatValue(stat4, true),
    });
  }

  if (stat5 !== "None") {
    data.push({
      stat: stat5,
      A: getStatValue(stat5, false),
      Opponent: getStatValue(stat5, true),
    });
  }

  if (stat6 !== "None") {
    data.push({
      stat: stat6,
      A: getStatValue(stat6, false),
      Opponent: getStatValue(stat6, true),
    });
  }

  const handleSearch = (event, newValue) => {
    if (newValue) {
      setOppSlug(newValue.slug);
    }
  };

  // Function to handle data point selections
  const handleDataPointSelect = (event, index) => {
    if (index == 1) {
      setStat1(event.target.value);
    }
    if (index == 2) {
      setStat2(event.target.value);
    }
    if (index == 3) {
      setStat3(event.target.value);
    }
    if (index == 4) {
      setStat4(event.target.value);
    }
    if (index == 5) {
      setStat5(event.target.value);
    }
    if (index == 6) {
      setStat6(event.target.value);
    }
  };

  const generateMenuItems = (stat: string): JSX.Element[] => {
    const menuItems: JSX.Element[] = [];

    menuItems.push(
      <MenuItem key={stat} value={stat}>
        {stat}
      </MenuItem>
    );
    for (let i = 0; i < allStats.length; i++) {
      let shouldAdd: boolean = true;
      for (let j = 0; j < selectedStats.length; j++) {
        if (allStats[i] === selectedStats[j]) {
          shouldAdd = false;
          break;
        }
      }
      if (shouldAdd) {
        menuItems.push(
          <MenuItem key={allStats[i]} value={allStats[i]}>
            {allStats[i]}
          </MenuItem>
        );
      }
    }

    let count = 0;
    for (let j = 0; j < selectedStats.length; j++) {
      if (selectedStats[j] !== "None") {
        count += 1;
      }
    }

    if (count > 3 && stat !== "None") {
      menuItems.push(
        <MenuItem key={"None"} value="None">
          {"None"}
        </MenuItem>
      );
    }

    return menuItems;
  };

  const getOriginalValue = (label: string): string => {
    let orgStat: string = "";
    switch (label) {
      case "Goals":
        orgStat = `${playerData.statistics?.[0].shooting.goals ?? 0}`;

        break;
      case "Shots on Target":
        orgStat = `${playerData.statistics?.[0].shooting.shots_on_target ?? 0}`;
        break;
      case "Passes":
        orgStat = `${playerData.statistics?.[0].technical.passes ?? 0}`;
        break;
      case "Assists":
        orgStat = `${playerData.statistics?.[0].shooting.assists ?? 0}`;
        break;
      case "Shots":
        orgStat = `${playerData.statistics?.[0].shooting.shots ?? 0}`;
        break;
      case "Dribble Success":
        orgStat = `${
          playerData.statistics?.[0].technical.dribbles_success ?? 0
        }`;
        break;
      case "Duels Won":
        orgStat = `${playerData.statistics?.[0].defense.duels_won ?? 0}`;
        break;
      case "Key Passes":
        orgStat = `${playerData.statistics?.[0].technical.key_passes ?? 0}`;
        break;
      case "Pass Accuracy":
        orgStat = `${playerData.statistics?.[0].technical.pass_accuracy ?? 0}%`;
        break;
      case "Saves":
        orgStat = `${playerData.statistics?.[0].shooting.saves ?? 0}`;
        break;
      default:
        return orgStat;
    }
    return orgStat;
  };

  const getOppOriginalValue = (label: string): string => {
    let orgStat: string = "";
    switch (label) {
      case "Goals":
        orgStat = `${oppData?.statistics?.[0].shooting.goals ?? 0}`;

        break;
      case "Shots on Target":
        orgStat = `${oppData?.statistics?.[0].shooting.shots_on_target ?? 0}`;
        break;
      case "Passes":
        orgStat = `${oppData?.statistics?.[0].technical.passes ?? 0}`;
        break;
      case "Assists":
        orgStat = `${oppData?.statistics?.[0].shooting.assists ?? 0}`;
        break;
      case "Shots":
        orgStat = `${oppData?.statistics?.[0].shooting.shots ?? 0}`;
        break;
      case "Dribble Success":
        orgStat = `${oppData?.statistics?.[0].technical.dribbles_success ?? 0}`;
        break;
      case "Duels Won":
        orgStat = `${oppData?.statistics?.[0].defense.duels_won ?? 0}`;
        break;
      case "Key Passes":
        orgStat = `${oppData?.statistics?.[0].technical.key_passes ?? 0}`;
        break;
      case "Pass Accuracy":
        orgStat = `${oppData?.statistics?.[0].technical.pass_accuracy ?? 0}%`;
        break;
      case "Saves":
        orgStat = `${oppData?.statistics?.[0].shooting.saves ?? 0}`;
        break;
      default:
        return orgStat;
    }
    return orgStat;
  };

  /* eslint-disable react/prop-types */
  const CustomTooltip = ({ active, payload, label }) => {
    if (active && payload) {
      const colorA = payload[0].color;
      const colorB = payload[1] ? payload[1].color : null;

      return (
        <Box
          sx={{
            background: "var(--dark0)",
            borderRadius: "10px",
            padding: "10px",
            color: "white",
          }}
        >
          <Typography>{label}</Typography>
          <Flex style={{ gap: "6px", color: colorA }}>
            {payload[0].name}:
            <Typography sx={{ fontWeight: 800, wrap: "inline" }}>
              {getOriginalValue(label)}
            </Typography>
          </Flex>
          {payload[1] && (
            <Flex style={{ gap: "6px", color: colorB }}>
              {payload[1].name}:
              <Typography sx={{ fontWeight: 800, wrap: "inline" }}>
                {getOppOriginalValue(label)}
              </Typography>
            </Flex>
          )}
        </Box>
      );
    }

    return null;
  };

  return (
    <Tile sx={{ flexDirection: "column" }}>
      <Grid container spacing={{ xs: 6, sm: 1, md: 1, lg: 1 }}>
        <Grid item xs={12} sm={12} md={12} lg={2.6} width="100%">
          <Flex style={{ flexDirection: "column", height: "100%" }}>
            <Typography
              sx={{
                fontSize: "20px",
                textAlign: "left",
                marginBottom: "24px",
                whiteSpace: "nowrap",
              }}
            >
              Compare Player Stats
            </Typography>
            <Flex
              style={{
                flexDirection: "column",
                justifyContent: "space-between",
                height: "100%",
                gap: "8px",
              }}
            >
              <Select
                size="small"
                value={stat1}
                onChange={(e) => handleDataPointSelect(e, 1)}
                sx={{ color: stat1 === "None" ? "grey" : "inherit" }}
              >
                {generateMenuItems(stat1)}
              </Select>
              <Select
                size="small"
                value={stat2}
                onChange={(e) => handleDataPointSelect(e, 2)}
                sx={{ color: stat2 === "None" ? "grey" : "inherit" }}
              >
                {generateMenuItems(stat2)}
              </Select>
              <Select
                size="small"
                value={stat3}
                onChange={(e) => handleDataPointSelect(e, 3)}
                sx={{ color: stat3 === "None" ? "grey" : "inherit" }}
              >
                {generateMenuItems(stat3)}
              </Select>
              <Select
                size="small"
                value={stat4}
                onChange={(e) => handleDataPointSelect(e, 4)}
                sx={{ color: stat4 === "None" ? "grey" : "inherit" }}
              >
                {generateMenuItems(stat4)}
              </Select>
              <Select
                size="small"
                value={stat5}
                onChange={(e) => handleDataPointSelect(e, 5)}
                sx={{ color: stat5 === "None" ? "grey" : "inherit" }}
              >
                {generateMenuItems(stat5)}
              </Select>
              <Select
                size="small"
                value={stat6}
                onChange={(e) => handleDataPointSelect(e, 6)}
                sx={{ color: stat6 === "None" ? "grey" : "inherit" }}
              >
                {generateMenuItems(stat6)}
              </Select>
            </Flex>
          </Flex>
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={6.2} width="100%">
          <Box
            sx={{
              width: "100%",
              height: { xs: "180px", sm: "350px", md: "350px", lg: "350px" },
              fontSize: { xs: "12px", sm: "16px", md: "16px", lg: "16px" },
            }}
          >
            <ResponsiveContainer>
              <RadarChart data={data} innerRadius={"10%"}>
                <PolarGrid gridType={"circle"} />
                <PolarAngleAxis
                  dataKey="stat"
                  stroke="white"
                  axisLine={false}
                  tickLine={false}
                />
                <Radar
                  name={playerData.name}
                  dataKey="A"
                  stroke="var(--red)"
                  fill="var(--red)"
                  fillOpacity={0.4}
                  strokeWidth="4px"
                />
                {oppSlug ? (
                  <Radar
                    name={oppData?.name}
                    dataKey="Opponent"
                    stroke="var(--blue)"
                    fill="var(--blue)"
                    fillOpacity={0.2}
                    strokeWidth="4px"
                  />
                ) : null}
                <Tooltip
                  contentStyle={{
                    background: "var(--dark0)",
                    border: "none",
                    borderRadius: "10px",
                  }}
                  content={
                    <CustomTooltip active={null} payload={null} label={null} />
                  }
                />
                <Legend />
              </RadarChart>
            </ResponsiveContainer>
          </Box>
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={3.2} width="100%">
          <Flex
            style={{
              flexDirection: "column",
              justifyContent: "space-around",
              height: "100%",
              gap: "8px",
            }}
          >
            <PlayerRadarSearch handleSearch={handleSearch} />

            {oppData?.photo ? (
              <PlayerOppIDCard slug={oppSlug} playerData={oppData} />
            ) : (
              <Tile
                sx={{
                  flexDirection: "column",
                  gap: "28px",
                  height: "100%",
                  background: "var(--dark1)",
                  border: "none",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Typography>None</Typography>
              </Tile>
            )}
          </Flex>
        </Grid>
      </Grid>
    </Tile>
  );
};

export default PlayerRadarChart;

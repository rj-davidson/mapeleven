import * as React from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import Typography from "@mui/material/Typography";
import { Tile } from "../Util/TileTS";
import { Flex } from "../Util/Flex.jsx";
import TeamRadarSearch from "./TeamRadarSearch.jsx";
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
import TeamOppIDCard from "./TeamOppIDCard";
import { APITeam } from "../Models/api_types";

interface TeamRadarChartProps {
  teamData: APITeam;
}

interface ChartDatum {
  stat: string;
  A: number;
  Opponent: number;
}

const url = import.meta.env.VITE_API_URL;

const TeamRadarChart: React.FC<TeamRadarChartProps> = ({ teamData }) => {
  const [oppSlug, setOppSlug] = useState<string>();
  const [oppName, setOppName] = useState<string>("");
  const [oppBadge, setOppBadge] = useState<string>("");
  const [oppGoals, setOppGoals] = useState<number>(0);
  const [oppClean, setOppClean] = useState<number>(0);
  const [oppWins, setOppWins] = useState<number>(0);
  const [oppFailed, setOppFailed] = useState<number>(0);
  const [oppScored, setOppScored] = useState<number>(0);
  const [oppAvgGoals, setOppAvgGoals] = useState<number>(0);
  const [oppCountry, setOppCountry] = useState<string>("");
  const [oppFlag, setOppFlag] = useState<string>("");
  const [oppFixtures, setOppFixtures] = useState<number>(0);
  const [oppLeagueSlug, setOppLeagueSlug] = useState<string>("");
  const [oppLeagueName, setOppLeagueName] = useState<string>("");
  const [oppLeagueLogo, setOppLeagueLogo] = useState<string>("");

  // Array of all available stats
  const allStats = [
    "Goals",
    "Clean Sheets",
    "Fixtures",
    "Wins",
    "Games Scored",
    "Average Goals",
    "Failed To Score",
  ];

  const [stat1, setStat1] = useState<string>("Wins");
  const [stat2, setStat2] = useState<string>("Games Scored");
  const [stat3, setStat3] = useState<string>("Clean Sheets");
  const [stat4, setStat4] = useState<string>("None");
  const [stat5, setStat5] = useState<string>("None");
  const [stat6, setStat6] = useState<string>("None");

  const selectedStats = [stat1, stat2, stat3, stat4, stat5, stat6];

  useEffect(() => {
    if (oppSlug) {
      fetch(`${url}/teams/${oppSlug}`)
        .then((response) => response.json())
        .then((jsonData) => {
          setOppName(jsonData.name.long);
          setOppBadge(jsonData.badge);
          setOppGoals(jsonData.competitions[0].stats.goals.for.total.total);
          setOppClean(jsonData.competitions[0].stats.clean_sheet.total);
          setOppCountry(jsonData.country.name);
          setOppFlag(jsonData.country.flag);
          setOppAvgGoals(
            jsonData.competitions[0].stats.goals.for.average.total
          );
          setOppWins(
            jsonData.competitions[0].stats.fixtures.wins.home +
              jsonData.competitions[0].stats.fixtures.wins.away
          );
          setOppFailed(jsonData.competitions[0].stats.failed_to_score.total);
          setOppFixtures(
            jsonData.competitions[0].stats.fixtures.played.home +
              jsonData.competitions[0].stats.fixtures.played.away
          );
          setOppScored(
            jsonData.competitions[0].stats.fixtures.played.home +
              jsonData.competitions[0].stats.fixtures.played.away -
              jsonData.competitions[0].stats.failed_to_score.total
          );
          setOppLeagueSlug(jsonData.competitions[0].league.slug);
          setOppLeagueName(jsonData.competitions[0].league.name);
          setOppLeagueLogo(jsonData.competitions[0].league.logo);
        })
        .catch((error) => console.error(error));
    }
  }, [oppSlug]);

  const data: ChartDatum[] = [];

  const goalNormal = 75;
  const fixNormal = 40;
  const avgNormal = 4;

  const fixtures =
    (teamData.competitions?.[0].stats?.fixtures.played.home || 0) +
    (teamData.competitions?.[0].stats?.fixtures.played.away || 0);

  const goals = teamData.competitions?.[0].stats?.goals.for.total.total || 0;

  const clean = teamData.competitions?.[0].stats?.clean_sheet.total || 0;

  const gamesScored =
    (teamData.competitions?.[0].stats?.fixtures.played.home || 0) +
    (teamData.competitions?.[0].stats?.fixtures.played.away || 0) -
    (teamData.competitions?.[0].stats?.failed_to_score.total || 0);

  const wins =
    (teamData.competitions?.[0].stats?.fixtures.wins.home || 0) +
    (teamData.competitions?.[0].stats?.fixtures.wins.away || 0);

  const averageGoals = parseFloat(
    teamData.competitions?.[0].stats?.goals.for.average.total ?? "0.0"
  );

  const failedToScore =
    teamData.competitions?.[0].stats?.failed_to_score.total || 0;

  const getStatValue = (stat: string, opp: boolean): number => {
    if (opp) {
      if (stat === "Goals") {
        return oppGoals / goalNormal;
      }
      if (stat === "Clean Sheets") {
        return oppClean / oppFixtures;
      }
      if (stat === "Fixtures") {
        return oppFixtures / fixNormal;
      }
      if (stat === "Wins") {
        return oppWins / oppFixtures;
      }
      if (stat === "Games Scored") {
        return oppScored / oppFixtures;
      }
      if (stat === "Average Goals") {
        return oppAvgGoals / avgNormal;
      }
      if (stat === "Failed To Score") {
        return oppFailed / oppFixtures;
      }
    }

    if (stat === "Goals") {
      return goals / goalNormal;
    }
    if (stat === "Clean Sheets") {
      return clean / fixtures;
    }
    if (stat === "Fixtures") {
      return fixtures / fixNormal;
    }
    if (stat === "Wins") {
      return wins / fixtures;
    }
    if (stat === "Games Scored") {
      return gamesScored / fixtures;
    }
    if (stat === "Average Goals") {
      return averageGoals / avgNormal;
    }
    if (stat === "Failed To Score") {
      return failedToScore / fixtures;
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
        orgStat = `${goals}`;
        break;
      case "Wins":
        orgStat = `${((wins / fixtures) * 100).toFixed(1)}% (${wins})`;
        break;
      case "Games Scored":
        orgStat = `${((gamesScored / fixtures) * 100).toFixed(
          1
        )}% (${gamesScored})`;
        break;
      case "Clean Sheets":
        orgStat = `${((clean / fixtures) * 100).toFixed(1)}% (${clean})`;
        break;
      case "Fixtures":
        orgStat = `${fixtures}`;
        break;
      case "Average Goals":
        orgStat = `${averageGoals}`;
        break;
      case "Failed To Score":
        orgStat = `${((failedToScore / fixtures) * 100).toFixed(
          1
        )}% (${failedToScore})`;
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
        orgStat = `${oppGoals}`;
        break;
      case "Wins":
        orgStat = `${((oppWins / oppFixtures) * 100).toFixed(1)}% (${oppWins})`;
        break;
      case "Games Scored":
        orgStat = `${((oppScored / oppFixtures) * 100).toFixed(
          1
        )}% (${oppScored})`;
        break;
      case "Clean Sheets":
        orgStat = `${((oppClean / oppFixtures) * 100).toFixed(
          1
        )}% (${oppClean})`;
        break;
      case "Fixtures":
        orgStat = `${oppFixtures}`;
        break;
      case "Average Goals":
        orgStat = `${oppAvgGoals}`;
        break;
      case "Failed To Score":
        orgStat = `${((oppFailed / oppFixtures) * 100).toFixed(
          1
        )}% (${oppFailed})`;
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
              sx={{ fontSize: "20px", textAlign: "left", marginBottom: "24px" }}
            >
              Compare Team Stats
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
                <PolarGrid gridType={"circle"} style={{ strokeWidth: 2 }} />
                <PolarAngleAxis
                  dataKey="stat"
                  stroke="white"
                  axisLine={false}
                  tickLine={false}
                />
                <Radar
                  name={teamData.name.long}
                  dataKey="A"
                  stroke="var(--red)"
                  fill="var(--red)"
                  fillOpacity={0.4}
                  strokeWidth="4px"
                />
                {oppSlug ? (
                  <Radar
                    name={oppName}
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
            <TeamRadarSearch handleSearch={handleSearch} />

            {oppBadge && oppFlag ? (
              <TeamOppIDCard
                slug={oppSlug}
                name={oppName}
                badge={oppBadge}
                country={oppCountry}
                flag={oppFlag}
                leagueSlug={oppLeagueSlug}
                leagueName={oppLeagueName}
                leagueLogo={oppLeagueLogo}
              />
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

export default TeamRadarChart;

import * as React from "react";
import { ResponsiveContainer, Pie, PieChart, Cell, Tooltip } from "recharts";
import { APIPlayer } from "../Models/api_types.js";
import { Tile } from "../Util/TileTS";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import { useMediaQuery } from "@mui/material";

interface PlayerStatsShootingPieChartProps {
  playerData: APIPlayer | undefined;
}

interface PlayerStatsShootingPieChartData {
  name: string;
  shots: number;
}

const PlayerStatsShootingPieChart: React.FC<
  PlayerStatsShootingPieChartProps
> = ({ playerData }) => {
  const isSmallerThanLG = useMediaQuery("(max-width: 1260px)");
  const shots: PlayerStatsShootingPieChartData[] = [
    {
      name: "Dribbles Failed",
      shots:
        (playerData?.statistics?.[0].technical.dribble_attempts || 0) -
        (playerData?.statistics?.[0].technical.dribbles_success || 0),
    },
    {
      name: "Dribbles Succeeded",
      shots: playerData?.statistics?.[0].technical.dribbles_success || 0,
    },
  ];

  const COLORS = [
    "#BF616A", // Red
    "#A3BE8C", // Green
  ];

  interface CustomTooltipProps {
    active?: boolean;
    payload?: any[];
  }

  const CustomTooltip: React.FC<CustomTooltipProps> = ({ active, payload }) => {
    if (active && payload && payload.length) {
      return (
        <div
          style={{
            background: "var(--dark0)",
            padding: "5px",
            borderRadius: "3px",
            color: "white",
          }}
        >
          {payload[0].name}: {payload[0].value}
        </div>
      );
    }
    return null;
  };

  return (
    <Tile sx={{ height: "250px" }}>
      <Grid container spacing={2}>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <Typography
            sx={{ fontSize: { xs: "20px", sm: "20px" }, textAlign: "left" }}
          >
            Technical Chart
          </Typography>
        </Grid>
        <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
          <Tile sx={{ height: "100%", background: "var(--dark1)" }}>
            <ResponsiveContainer>
              <PieChart>
                <Pie
                  data={shots}
                  innerRadius="30%"
                  outerRadius="45%"
                  dataKey="shots"
                  label={({
                    cx,
                    cy,
                    midAngle,
                    innerRadius,
                    outerRadius,
                    index,
                  }) => {
                    const RADIAN = Math.PI / 180;
                    const radius =
                      25 + innerRadius + (outerRadius - innerRadius);
                    const x = cx + radius * Math.cos(-midAngle * RADIAN);
                    const y = cy + radius * Math.sin(-midAngle * RADIAN);
                    if (shots[index].shots === 0) {
                      return null;
                    }
                    return (
                      <text
                        x={x}
                        y={y}
                        fill="white"
                        textAnchor={x > cx ? "start" : "end"}
                        dominantBaseline="central"
                        style={{ fontSize: isSmallerThanLG ? "12px" : "18px" }}
                      >
                        {shots[index].name}
                      </text>
                    );
                  }}
                >
                  {shots.map((entry, index) => (
                    <Cell
                      key={`cell-${index}`}
                      fill={
                        shots[index].shots === 0
                          ? "transparent"
                          : COLORS[index % COLORS.length]
                      }
                    />
                  ))}
                </Pie>
                <Tooltip content={<CustomTooltip />} />
              </PieChart>
            </ResponsiveContainer>
          </Tile>
        </Grid>
      </Grid>
    </Tile>
  );
};

export default PlayerStatsShootingPieChart;

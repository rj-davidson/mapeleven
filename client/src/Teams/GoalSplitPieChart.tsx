import * as React from "react";
import { ResponsiveContainer, Pie, PieChart, Cell, Tooltip } from "recharts";
import { APITSGoalMinuteSplit } from "../Models/api_types.js";

interface GoalSplitPieChartProps {
  data: APITSGoalMinuteSplit | undefined;
}

interface GoalSplitPieChartData {
  name: string;
  goals: number;
}

const GoalSplitPieChart: React.FC<GoalSplitPieChartProps> = ({ data }) => {
  const goals: GoalSplitPieChartData[] = [
    { name: "0'-15'", goals: data?.["0-15"].total || 0 },
    { name: "16'-30'", goals: data?.["16-30"].total || 0 },
    { name: "31'-45'", goals: data?.["31-45"].total || 0 },
    { name: "46'-60'", goals: data?.["46-60"].total || 0 },
    { name: "61'-75'", goals: data?.["61-75"].total || 0 },
    { name: "76'-90'", goals: data?.["76-90"].total || 0 },
    { name: "91'-105'", goals: data?.["91-105"].total || 0 },
    { name: "106'-120'", goals: data?.["106-120"].total || 0 },
  ];

  const COLORS = [
    "#BF616A", // Red
    "#D08770", // Orange
    "#EBCB8B", // Yellow
    "#A3BE8C", // Green
    "#5E81AC", // Blue
    "#B48EAD", // Indigo
    "#FFC0CB", // Pink
    "#7FB3D5", // Light Blue
  ];

  interface CustomTooltipProps {
    active?: boolean;
    payload?: any[];
  }

  const CustomTooltip: React.FC<CustomTooltipProps> = ({ active, payload }) => {
    if (active && payload && payload.length) {
      const totalGoals = payload[0].value;
      const tooltipStyle = {
        background: "var(--dark0)",
        padding: "5px",
        borderRadius: "3px",
        color: "white",
      };
      return <div style={tooltipStyle}>Goals: {totalGoals}</div>;
    }
    return null;
  };

  return (
    <ResponsiveContainer>
      <PieChart>
        <Pie
          data={goals}
          innerRadius="30%"
          outerRadius="45%"
          dataKey="goals"
          label={({ cx, cy, midAngle, innerRadius, outerRadius, index }) => {
            const RADIAN = Math.PI / 180;
            const radius = 25 + innerRadius + (outerRadius - innerRadius);
            const x = cx + radius * Math.cos(-midAngle * RADIAN);
            const y = cy + radius * Math.sin(-midAngle * RADIAN);
            if (goals[index].goals === 0) {
              return null;
            }
            return (
              <text
                x={x}
                y={y}
                fill="white"
                textAnchor={x > cx ? "start" : "end"}
                dominantBaseline="central"
              >
                {goals[index].name}
              </text>
            );
          }}
        >
          {goals.map((entry, index) => (
            <Cell
              key={`cell-${index}`}
              fill={
                goals[index].goals === 0
                  ? "transparent"
                  : COLORS[index % COLORS.length]
              }
            />
          ))}
        </Pie>
        <Tooltip content={<CustomTooltip />} />
      </PieChart>
    </ResponsiveContainer>
  );
};

export default GoalSplitPieChart;

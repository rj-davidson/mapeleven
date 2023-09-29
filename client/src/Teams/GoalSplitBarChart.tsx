import * as React from 'react';
import { ResponsiveContainer, BarChart, Cell } from 'recharts';
import { APITSGoalMinuteSplit } from '../Models/api_types.js';
import { Bar, XAxis, YAxis } from 'recharts';

interface GoalSplitBarChartProps {
    data: APITSGoalMinuteSplit;
}

interface GoalSplitBarChartData {
    name: string;
    goals: number;
}

const GoalSplitBarChart: React.FC<GoalSplitBarChartProps> = ({ data }) => {
    const goals: GoalSplitBarChartData[] = [
        { name: "0'-15'", goals: data['0-15'].total },
        { name: "16'-30'", goals: data['16-30'].total },
        { name: "31'-45'", goals: data['31-45'].total },
        { name: "46'-60'", goals: data['46-60'].total },
        { name: "61'-75'", goals: data['61-75'].total },
        { name: "76'-90'", goals: data['76-90'].total },
        { name: "91'-105'", goals: data['91-105'].total },
        { name: "106'-120'", goals: data['106-120'].total },
    ];

    const COLORS = [
        '#BF616A', // Red
        '#D08770', // Orange
        '#EBCB8B', // Yellow
        '#A3BE8C', // Green
        '#5E81AC', // Blue
        '#B48EAD', // Indigo
        '#FFC0CB', // Pink
        '#7FB3D5', // Light Blue
    ];

    return (
        <ResponsiveContainer>
            <BarChart
                data={goals}
                margin={{
                    top: 20,
                    right: 30,
                    left: 20,
                    bottom: 5,
                }}
            >
                <XAxis dataKey='name' tick={{ fill: 'white' }} />
                <Bar dataKey='goals' label={{ position: 'top' }}>
                    {goals.map((entry, index) => (
                        <Cell key={`cell-${index}`} fill={COLORS[index % 20]} />
                    ))}
                </Bar>
            </BarChart>
        </ResponsiveContainer>
    );
};

export default GoalSplitBarChart;

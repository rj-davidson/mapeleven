import React, { useState, useRef, useEffect } from "react";
import * as d3 from 'd3';
import {RadarChart} from './radarChart.js';

function SpiderChart() {
    // Declare state variables using the `useState` hook.
    const [data, setData] = useState([[]]);

    // Create a reference for the SVG element.
    const svgRef = useRef();

    // Define a function to handle the button click.
    const handleClick = () => {
        // Send a GET request to the API.
        fetch('http://localhost:8080/stats/topscorers')
            .then(response => response.json())
            .then(jsonData => {
                console.log(jsonData.response);
                // Extract the data from the response.
                const response = jsonData.response;
                const newData = []

                const firstshots = response[0].statistics[0].shots.total;
                const firston = response[0].statistics[0].shots.on;
                const firstpasses = response[0].statistics[0].passes.key;
                const firstduels = response[0].statistics[0].duels.won;
                const firstdribbles = response[0].statistics[0].dribbles.success;

                const first = [//iPhone
                    {axis:"Shots",value:firstshots},
                    {axis:"Shots on Target",value:firston},
                    {axis:"Key Passes",value:firstpasses},
                    {axis:"Duels Won",value:firstduels},
                    {axis:"Successful Dribbles",value:firstdribbles},
                ]
                
                const secondshots = response[1].statistics[0].shots.total;
                const secondon = response[1].statistics[0].shots.on;
                const secondpasses = response[1].statistics[0].passes.key;
                const secondduels = response[1].statistics[0].duels.won;
                const seconddribbles = response[1].statistics[0].dribbles.success;

                const second = [//iPhone
                    {axis:"Shots",value:secondshots},
                    {axis:"Shots on Target",value:secondon},
                    {axis:"Key Passes",value:secondpasses},
                    {axis:"Duels Won",value:secondduels},
                    {axis:"Successful Dribbles",value:seconddribbles},
                ]

                const thirdshots = response[19].statistics[0].shots.total;
                const thirdon = response[19].statistics[0].shots.on;
                const thirdpasses = response[19].statistics[0].passes.key;
                const thirdduels = response[19].statistics[0].duels.won;
                const thirddribbles = response[19].statistics[0].dribbles.success;

                const third = [//iPhone
                    {axis:"Shots",value:thirdshots},
                    {axis:"Shots on Target",value:thirdon},
                    {axis:"Key Passes",value:thirdpasses},
                    {axis:"Duels Won",value:thirdduels},
                    {axis:"Successful Dribbles",value:thirddribbles},
                ]
                newData.push(first);
                newData.push(second);
                newData.push(third);
                // Update the state variables with the new data.
                setData(newData);
            })
            .catch(error => console.error(error));
    };

    // Define the D3 visualization using the `useEffect` hook.
    useEffect(() => {
        // Set up the SVG container.
        const w = 600;
        const h = 400;
        const svg = d3.select(svgRef.current)
            .attr('width', w)
            .attr('height', h)
            .style('overflow', 'visible')
            .style('margin-top', '75px')
            .style('margin-bottom', '280px');

        var margin = {top: 100, right: 100, bottom: 100, left: 100},
            width = Math.min(700, window.innerWidth - 10) - margin.left - margin.right,
            height = Math.min(width, window.innerHeight - margin.top - margin.bottom - 20);

        const color = d3.scaleOrdinal()
            .range(["#3854FC","#CC333F", "green"]);

        var radarChartOptions = {
            w: width,
            h: height,
            margin: margin,
            maxValue: 0.5,
            levels: 5,
            roundStrokes: true,
            color: color
        };
        //Call function to draw the Radar chart
        RadarChart(svgRef.current, data, radarChartOptions);

    }, [data]);

    return (
        <div className="BarChart">
            {/* Render a button that triggers the `handleClick` function when clicked. */}
            <button className={"button"} onClick={handleClick}>Get Top Scorers</button>
            {/* Render an SVG element using the `svgRef` reference. */}
            <svg ref={svgRef}></svg>
        </div>
    );
}

export default SpiderChart;


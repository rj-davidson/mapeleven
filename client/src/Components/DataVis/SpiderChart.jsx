import React, { useState, useRef, useEffect } from "react";
import * as d3 from 'd3';
import {RadarChart} from './radarChart.js';
import './SpiderChart.css'

function SpiderChart() {
    // Declare state variables using the `useState` hook.
    const [data, setData] = useState([[]]);
    const [playerList, setList] = useState([]);
    const [player1, setPlayer1] = useState(0);
    const [player2, setPlayer2] = useState(1);

    // Create a reference for the SVG element.
    const svgRef = useRef();
    // Fetch the player list when the page loads.
    useEffect(() => {
        fetch('http://localhost:8080/stats/topscorers')
            .then(response => response.json())
            .then(jsonData => {
                console.log(jsonData.response[0].player.name);
                // Extract the data from the response.
                const response = jsonData.response;
                const newNames = []
                for (let i = 0; i < response.length; i++) {
                    const name = response[i].player.name;
                    newNames.push(name);
                }

                // Update the state variables with the new data.
                setList(newNames);
            })
            .catch(error => console.error(error));
    }, []);

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

                function addPlayers(i){
                    const goals = response[i].statistics[0].goals.total;
                    const assists = response[i].statistics[0].goals.assists;
                    const shotson = response[i].statistics[0].shots.on;
                    const keypasses = response[i].statistics[0].passes.key;
                    const dribbles = response[i].statistics[0].dribbles.success;

                    const player = [
                        {axis:"Goals",value:goals},
                        {axis:"Assists",value:assists},
                        {axis:"Shots on Target",value:shotson},
                        {axis:"Key Passes",value:keypasses},
                        {axis:"Successful Dribbles",value:dribbles},
                    ]

                    newData.push(player);
                }

                addPlayers(player1);
                addPlayers(player2);

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
        d3.select(svgRef.current)
            .attr('width', w)
            .attr('height', h)
            .style('overflow', 'visible');

        var margin = {top: 100, right: 100, bottom: 100, left: 100},
            width = Math.min(700, window.innerWidth - 10) - margin.left - margin.right,
            height = Math.min(width, window.innerHeight - margin.top - margin.bottom - 20);

        const color = d3.scaleOrdinal()
            .range(["#3854FC", "#CC333F"]);

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
        <div className="SpiderChart">
            <div className="spider-ui">
                <button className={"button"} onClick={handleClick}>generate chart</button>
                <div className={"drop-downs"}>
                    <div className="pl1">
                        <div className="player1title">player 1</div>
                        <select className="player1" value={player1} onChange={e => setPlayer1(parseInt(e.target.value))}>
                            <option value="">select player 1</option>
                            {playerList.map((playerName, index) => (
                                <option key={index} value={index}>{playerName}</option>
                            ))}
                        </select>
                    </div>
                    <div className="pl1">
                        <div>player 2</div>
                        <select className="player2" value={player2} onChange={e => setPlayer2(parseInt(e.target.value))}>
                            <option value="">select player 2</option>
                            {playerList.map((playerName, index) => (
                                <option key={index} value={index}>{playerName}</option>
                            ))}
                        </select>
                    </div>
                </div>
            </div>
            <div className="chart">
                <svg ref={svgRef}></svg>
            </div>
        </div>
    );
}

export default SpiderChart;


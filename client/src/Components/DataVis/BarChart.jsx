import React, { useState, useRef, useEffect } from "react";
import * as d3 from 'd3';
import "./BarChart.css";

function BarChart() {
    // Declare state variables using the `useState` hook.
    const [data, setData] = useState([]);
    const [names, setNames] = useState([]);

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
                const newNames = []
                for (let i = 0; i < response.length; i++) {
                    const goals = response[i].statistics[0].goals.total;
                    const name = response[i].player.name;
                    newData.push(goals);
                    newNames.push(name);
                }
                // Update the state variables with the new data.
                setData(newData);
                setNames(newNames);
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
            .style('margin-top', '75px');

        // Set up the scaling.
        const xScale = d3.scaleBand()
            .domain(data.map((val, i) => i))
            .range([0, w])
            .padding(0.5);

        const yScale = d3.scaleLinear()
            .domain([0, d3.max(data)])
            .range([h, 0]);

        // Set up the axis.
        const xAxis = d3.axisBottom(xScale)
            .ticks(data.length)
            .tickFormat((d,i) => names[i]);

        const yAxis = d3.axisLeft(yScale)
            .ticks(5);
        svg.append('g')
            .attr('class', 'axis')
            .call(xAxis)
            .attr('transform', `translate(0, ${h})`)
            .selectAll("text")
            .style("text-anchor", "end")
            .attr("dx", "-.8em")
            .attr("dy", ".15em")
            .attr("transform", "rotate(-65)");
        svg.append('g')
            .attr('class', 'axis')
            .call(yAxis);

        // Labels
        svg.append('text')
            .attr('x', w/2)
            .attr('y', h + 80)
            .text('players')
            .attr('class', 'labels');
        svg.append('text')
            .attr('x', -80)
            .attr('y', h/2)
            .text('goals')
            .attr('class', 'labels');

        // Set up the bars.
        svg.selectAll('.bar')
            .data(data)
            .join('rect')
            .transition()
            .duration(3000)
            .attr('x', (v, i) => xScale(i))
            .attr('y', yScale)
            .attr('width', xScale.bandwidth())
            .attr('height', val => h - yScale(val))
            .attr('class', 'bar')

    }, [data]);

    return (
        <div className="BarChart">
            {/* Render a button that triggers the `handleClick` function when clicked. */}
            <button className={"button"} onClick={handleClick}>get top scorers</button>
            {/* Render an SVG element using the `svgRef` reference. */}
            <svg ref={svgRef}></svg>
        </div>
    );
}

export default BarChart;


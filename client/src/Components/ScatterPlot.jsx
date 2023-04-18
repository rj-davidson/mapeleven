import React, { useState, useRef, useEffect } from "react";
import * as d3 from 'd3';
import "./ScatterPlot.css";

function ScatterPlot() {
    // Declare state variables using the `useState` hook.
    const [data, setData] = useState([
        { x: 40, y: 30 },
        { x: 20, y: 15 },
        { x: 8, y: 2 },
        { x: 12, y: 5 }
    ]);
    //const [names, setNames] = useState([]);

    // Create a reference for the SVG element.
    const svgRef = useRef();

    // Define a function to handle the button click.
    const handleClick = () => {
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
        const xScale = d3.scaleLinear()
            .domain([0, d3.max(data, d => d.x)])
            .range([0, w]);

        const yScale = d3.scaleLinear()
            .domain([0, d3.max(data, d => d.y)])
            .range([h, 0]);

        // Set up the axis.
        const xAxis = d3.axisBottom(xScale)
            .ticks(8);

        const yAxis = d3.axisLeft(yScale)
            .ticks(8);

        svg.append('g')
            .attr('class', 'axis')
            .call(xAxis)
            .attr('transform', `translate(0, ${h})`);
        svg.append('g')
            .attr('class', 'axis')
            .call(yAxis);

        svg.append('text')
            .attr('x', w/2)
            .attr('y', h + 50)
            .text('shots')
            .attr('class', 'labels');
        svg.append('text')
            .attr('x', -80)
            .attr('y', h/2)
            .text('goals')
            .attr('class', 'labels');

        // Set up the bars.
        svg.selectAll()
            .data(data)
            .enter()
            .append('circle')
                .attr('cx', d => xScale(d.x))
                .attr('cy', d => yScale(d.y))
                .attr('r', 4)
                .attr('class', 'point')
                .append('title') // add a `title` element to the `circle` elements
                    .text(d => `x: ${d.x}, y: ${d.y}`);

    }, [data]);

    return (
        <div className="ScatterPlot">
            <div className="container">
                {/* Render a button that triggers the `handleClick` function when clicked. */}
                <button className={"button"} onClick={handleClick}>Get Top Scorers</button>
                {/* Render an SVG element using the `svgRef` reference. */}
                <svg ref={svgRef}></svg>
            </div>
        </div>
    );
}

export default ScatterPlot;


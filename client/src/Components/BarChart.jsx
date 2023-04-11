import React, {useState, useRef, useEffect} from "react";
import * as d3 from 'd3';

function BarChart() {
    const [data] = useState([100, 80, 180, 300, 120, 220]);
    const svgRef = useRef();

    useEffect(() => {
        // set up container
        const w = 400;
        const h = 300;
        const svg = d3.select(svgRef.current)
            .attr('width', w)
            .attr('height', h)
            .style('overflow', 'visible')
            .style('margin-top', '75px');

        // set up scaling
        const xScale = d3.scaleBand()
            .domain(data.map((val, i) => i))
            .range([0, w])
            .padding(0.5);

        const yScale = d3.scaleLinear()
            .domain([0, h])
            .range([h, 0]);

        // set up the axis
        const xAxis = d3.axisBottom(xScale)
            .ticks(data.length);
        const yAxis = d3.axisLeft(yScale)
            .ticks(5);
        svg.append('g')
            .attr('class', 'axis')
            .call(xAxis)
            .attr('transform', `translate(0, ${h})`);
        svg.append('g')
            .attr('class', 'axis')
            .call(yAxis);

        // setting the svg data
        svg.selectAll('.bar')
            .data(data)
            .join('rect')
            .attr('x', (v, i) => xScale(i))
            .attr('y', yScale)
            .attr('width', xScale.bandwidth())
            .attr('height', val => h - yScale(val))
            .attr('class', 'bar');

    }, [data]);

    return (
        <div className="BarChart">
            <svg ref={svgRef}></svg>
        </div>
    );
}

export default BarChart;

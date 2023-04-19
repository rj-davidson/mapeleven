import React, { useState, useRef, useEffect } from "react";
import * as d3 from 'd3';
import "./PictureButton.css";

function BarChart() {
    // Declare state variables using the `useState` hook.
    const [data, setData] = useState([]);
    const [names, setNames] = useState([]);

    // Create a reference for the SVG element.
    const svgRef = useRef();

    // Define a function to handle the button click.
    const handleClick = () => {
        // Send a GET request to the API.
        fetch('http://localhost:8080/league/39')
            .then(response => response.json())
            .then(jsonData => {
                setJasonData(jsonData);
                // Extract the data from the response.
                /*const response = jsonData.getElementsByTagName("league");

                const img = document.createElement('img');
                img.src = imageUrl;
                const container = document.querySelector('.image-container');
                container.appendChild(img);
            */
            })

            .catch(error => console.error(error));
    };

    // Define the D3 visualization using the `useEffect` hook.

    return (
        <div>
            <button onClick={handleClick}>Get Data</button>
        </div>
    );
}

export default BarChart;


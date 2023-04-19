import React from "react";
import '../../App.css';
import NavigationBar from "../NavigationBar.jsx";
import BarChart from "../BarChart.jsx";

function Players() {
    return (
        <div>
            <NavigationBar />
            <BarChart />
            <h1 className='players'>PLAYERS</h1>;
        </div>
    );
}

export default Players

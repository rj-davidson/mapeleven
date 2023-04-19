import React from "react";
import '../../App.css';
import BarChart from "../BarChart";
import NavigationBar from "../NavigationBar";

function Home() {
    return (
        <div>
            <NavigationBar />
            <BarChart />
        </div>
    );
}

export default Home;
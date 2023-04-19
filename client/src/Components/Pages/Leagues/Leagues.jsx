import React from "react";
import '../../../App.css';
import NavigationBar from "../../Layout/NavigationBar.tsx";
import DisplayLeagues from "./DisplayLeagues.jsx";

function Leagues() {
    return (
        <div>
            <NavigationBar />
            <DisplayLeagues />
            <h1 className='leagues'>LEAGUES</h1>;
        </div>
    );
}

export default Leagues;
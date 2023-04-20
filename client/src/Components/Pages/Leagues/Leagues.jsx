import React from "react";
import '../../../App.css';
import DisplayLeagues from "./DisplayLeagues.jsx";
import Layout from "../../Layout/Layout.tsx";

function Leagues() {
    return (
        <div>
            <Layout>
                <DisplayLeagues />
                <h1 className='leagues'>LEAGUES</h1>;
            </Layout>
        </div>
    );
}

export default Leagues;
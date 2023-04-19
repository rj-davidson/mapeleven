import React from "react";
import '../../../App.css';
import Layout from "../../Layout/Layout.tsx";
import DaySwitcher from "./DaySwitcher.tsx";

function Fixtures() {
    return (
        <div>
            <Layout>
                <DaySwitcher />
                <h1 className='fixtures'>FIXTURES</h1>;
            </Layout>
        </div>
    );
}

export default Fixtures;

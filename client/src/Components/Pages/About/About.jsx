import React from "react";
import '../../../App.css';
import NavigationBar from "../../Layout/NavigationBar.tsx";

function About() {
    return(
        <div>
            <NavigationBar />
        <h1>
            Team Member Names:

            Ethan Quinlan
            James Gibb
            RJ Davidson
            Ronnie Koe



            Project Summary:
            A web app that focuses on soccer data visualization. Our aim is to use match event data to help fans, scouts and clubs better understand player and team performances. Our scope is not limited to data visualization but also data analytics and data journalism.


            Published Cloud Site:
            We are planning to use AWS.


            Complete Project Description:
            Map Eleven aims to provide soccer fans, clubs, and scouts a tool to visualize match event data. The web application’s name is a combination of the idea of mapping data and the number of players a team can field in a soccer match, Map Eleven. The match event data is complex and hard to understand without tools to visualize the data. With Map Eleven, we want to provide users with a data dashboard that clearly communicates this data through visualization techniques. The charts will give valuable insights into the data of players, matches, clubs and leagues. The data for this project is already collected by several companies. We plan to import this data and use scalable technologies to create a service that will allow users to customize their experience using the visualization tools that we create. The purpose of this custom experience is to allow users to best understand the data by selecting their own data points and chart type.
        </h1>
        </div>
    );
}

export default About;

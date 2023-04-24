import * as React from 'react';
import { Box, List, ListItem, ListItemText, Typography } from '@mui/material';
import '../../../App.css';

function About() {
    return (
        <Box sx={{ p: 2 }}>
            <Typography variant='body1' mb={2}>
                Team Member Names: Ethan Quinlan, James Gibb, RJ Davidson, Ronnie Koe
            </Typography>
            <Typography variant='body1' mb={2}>
                Project Summary: A web app that focuses on soccer data visualization. Our aim is to use match event data
                to help fans, scouts and clubs better understand player and team performances. Our scope is not limited
                to data visualization but also data analytics and data journalism.
            </Typography>
            <Typography variant='body1' mb={2}>
                Published Cloud Site: We are planning to use AWS.
            </Typography>
            <Typography variant='body1' mb={2}>
                Complete Project Description: mapeleven aims to provide soccer fans, clubs, and scouts a tool to
                visualize match event data. The web application’s name is a combination of the idea of mapping data and
                the number of players a team can field in a soccer match, mapeleven. The match event data is complex and
                hard to understand without tools to visualize the data. With mapeleven, we want to provide users with a
                data dashboard that clearly communicates this data through visualization techniques. The charts will
                give valuable insights into the data of players, matches, clubs and leagues. The data for this project
                is already collected by several companies. We plan to import this data and use scalable technologies to
                create a service that will allow users to customize their experience using the visualization tools that
                we create. The purpose of this custom experience is to allow users to best understand the data by
                selecting their own data points and chart type.
            </Typography>
            <Typography variant='body1'>
                Our Goals:
                <List>
                    <ListItem>
                        <ListItemText primary='Provide a user-friendly platform for soccer data analysis' />
                    </ListItem>
                    <ListItem>
                        <ListItemText primary='Help clubs, scouts and fans make data-driven decisions' />
                    </ListItem>
                    <ListItem>
                        <ListItemText primary='Use creative data visualization techniques to communicate insights' />
                    </ListItem>
                    <ListItem>
                        <ListItemText primary='Bring transparency to soccer data and enable informed debate' />
                    </ListItem>
                </List>
            </Typography>
        </Box>
    );
}

export default About;

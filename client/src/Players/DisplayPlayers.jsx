import React, {useState} from "react";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import {console} from "yarn/lib/cli.js";

class DisplayPlayers extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            players: []
        };
    }

    componentDidMount() {
        fetch('http://localhost:8080/stats/topscorers')
            .then(response => response.json())
            .then(jsonData => {
                console.log(jsonData.response[0].player.name);
                // Extract the data from the response.
                const response = jsonData.response;
                const newPlayers = []
                for (let i = 0; i < response.length; i++) {
                    const player = response[i].player;
                    newPlayers.push(player);
                }

                // Update the state variables with the new data.
                this.state.players = newPlayers;
            })
            .catch(error => console.error(error));
    }

    render() {
        let table_data = this.state.players.map(player => {
            return (
                <TableRow key={player.id}>
                    <TableCell>{player.name}</TableCell>
                    <TableCell>{player.position}</TableCell>
                    <TableCell>{player.height}</TableCell>
                    <TableCell>{player.weight}</TableCell>
                </TableRow>
            );
        });

        return (
            <div>
                <TableContainer>
                    <Table>
                        <caption>Players</caption>
                        <TableHead>
                            <TableRow>
                                <TableCell>Player Name</TableCell>
                                <TableCell>Player Position</TableCell>
                                <TableCell>Player Height</TableCell>
                                <TableCell>Player Weight</TableCell>
                            </TableRow>
                        </TableHead>

                        <TableBody>{table_data}</TableBody>
                    </Table>
                </TableContainer>
            </div>
        );
    }
}

export default DisplayPlayers;


//<TableCell>{player.nationality}</TableCell>
//<TableCell>{player.age}</TableCell>
//export default DisplayPlayers;
/*
* constructor(props) {
        super(props);
        this.state = {
            players: []
        };
    }
* */

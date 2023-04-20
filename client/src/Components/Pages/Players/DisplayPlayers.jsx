import React from "react";
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
        fetch("http://localhost:8080/stats/topscorers")
            .then(response => response.json())
            .then(data => {
            this.setState({ players: data });
            console.log(data);
            console.log(data.players);
            console.log(data.players[0].player.name);
        })
            .catch(error => console.log(error));
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

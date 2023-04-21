import React from "react";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import {json} from "react-router-dom";

class TestDisplayPlayers extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            players: []
        }

        this.componentDidMount = this.componentDidMount.bind(this)
        this.componentDidMount();
    }

    componentDidMount() {
        fetch("http://localhost:8080/stats/topscorers").then(
            (response) => response.json()
        ).then(data => {
            const response = data.response;
            console.log(data)
            for(let i = 0; i < response.length; i++) {
                const player = response[i].player;
                this.state.players.push(player);
            }
            console.log(this.state.players)
        })
    }

    render() {
        let table_data = this.state.players.map(player => {
            return (
                <TableRow key={player.id}>
                    <TableCell>{player.name}</TableCell>
                    <TableCell>{player.nationality}</TableCell>
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
                                <TableCell>Player Nationality</TableCell>
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

export default TestDisplayPlayers;

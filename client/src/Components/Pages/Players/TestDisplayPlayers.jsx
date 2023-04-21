import React from "react";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";

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
            const tempPlayers = [];
            console.log(data)
            for(let i = 0; i < response.length; i++) {
                const player = response[i].player;
                tempPlayers.push(player);
            }
            this.setState({ players: tempPlayers })
            console.log(this.state.players)
        })
    }

    render() {
        let table_data = this.state.players.map(player => {
            return (
                <TableRow key={player.id}>
                    <TableCell>{player.photo}</TableCell>
                    <TableCell>{player.name}</TableCell>
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
                                <TableCell>Player Photo</TableCell>
                                <TableCell>Player Name</TableCell>
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

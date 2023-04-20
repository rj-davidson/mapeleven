import React from "react";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";

class DisplayLeagues extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            leagues: []
        }

        this.componentDidMount = this.componentDidMount.bind(this)
        this.componentDidMount();
    }

    componentDidMount() {
        fetch("http://localhost:8080/leagues").then(
            (response) => response.json()
        ).then(data => {
            this.setState({ leagues: data })
            console.log(data)
        })
    }

    render() {
        let table_data = this.state.leagues.map((league) => {
            return (
                <TableRow key={league.id}>
                    <TableCell>{league.id}</TableCell>
                    <TableCell>{league.name}</TableCell>
                    <TableCell>{league.type}</TableCell>
                </TableRow>
            )
        })

        return(
            <div>
                <TableContainer>
                    <Table>
                        <caption>Leagues</caption>
                        <TableHead>
                            <TableRow>
                                <TableCell>League Id</TableCell>
                                <TableCell>League Name</TableCell>
                                <TableCell>League Type</TableCell>
                            </TableRow>
                        </TableHead>

                        <TableBody>
                            {table_data}
                        </TableBody>
                    </Table>
                </TableContainer>
            </div>
        )
    }

}

export default DisplayLeagues;

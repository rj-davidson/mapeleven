import React from "react";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";

class DisplayClubs extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clubs: []
        }

        this.componentDidMount = this.componentDidMount.bind(this)
        this.componentDidMount();
    }

    componentDidMount() {
        fetch("http://localhost:8080/teams").then(
            (response) => response.json()
        ).then(data => {
            this.setState({ clubs: data })
            console.log(data)
        })
    }

    render() {
        let table_data = this.state.clubs.map((club) => {
            return (
                <TableRow key={club.id}>
                    <TableCell>{<img src={'http://localhost:8080/' + club.logo} alt="League Photo" width="50" height="50" />}</TableCell>
                    <TableCell>{club.name}</TableCell>
                    <TableCell>{club.founded}</TableCell>
                    <TableCell>{club.code}</TableCell>
                </TableRow>
            )
        })

        return(
            <div>
                <TableContainer>
                    <Table>
                        <caption>Clubs</caption>
                        <TableHead>
                            <TableRow>
                                <TableCell></TableCell>
                                <TableCell>Club Name</TableCell>
                                <TableCell>Club Founded</TableCell>
                                <TableCell>Club Code</TableCell>
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

export default DisplayClubs;

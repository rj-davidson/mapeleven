import React from "react";
import {Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography} from "@mui/material";
import {Link} from "react-router-dom";

const url = import.meta.env.VITE_API_URL;

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
        fetch(`{${url}/leagues`).then(
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
                    <TableCell>{<img src={`${url}/` + league.logo} alt="League Photo" width="50" height="50" />}</TableCell>
                    <TableCell>{league.id}</TableCell>
                    <Link to={`/teams/${league.slug}`}>
                        <Typography
                            sx={{
                                '&:hover': {
                                    textDecoration: 'underline', // Add underline on hover
                                },
                            }}
                        >
                            {league.name}
                        </Typography>
                    </Link>
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
                                <TableCell></TableCell>
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

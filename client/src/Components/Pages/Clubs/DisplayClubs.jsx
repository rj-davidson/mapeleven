import React from "react";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";

class DisplayClubs extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clubs: []
        }
        this.componentDidMount = this.componentDidMount.bind(this);
        this.componentDidMount();
    }

    componentDidMount() {
        fetch("http://localhost:8080/teams")
          .then(response => response.json())
          .then(data => this.setState({ clubs: data }));
    }

    render() {
        let table_data = this.state.clubs.map((team) => (
          <TableRow key={team.slug}>
              <TableCell>
                  <img
                    src={'http://localhost:8080/' + team.badge}
                    alt={`${team.name.long} Badge`}
                    width="50"
                    height="50"
                  />
              </TableCell>
              <TableCell>{team.name.long}</TableCell>
              <TableCell>{team.founded}</TableCell>
              <TableCell>{team.country ? team.country.code : "-"}</TableCell>
          </TableRow>
        ));

        return(
          <TableContainer>
              <Table>
                  <caption>Clubs</caption>
                  <TableHead>
                      <TableRow>
                          <TableCell></TableCell>
                          <TableCell>Club Name</TableCell>
                          <TableCell>Founded</TableCell>
                          <TableCell>Country Code</TableCell>
                      </TableRow>
                  </TableHead>

                  <TableBody>
                      {table_data}
                  </TableBody>
              </Table>
          </TableContainer>
        )
    }
}

export default DisplayClubs;

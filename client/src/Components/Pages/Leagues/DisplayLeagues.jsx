import React from "react";

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


    //.then(leagues => this.setState({ leagues: leagues }));
    render() {
       let table_data = this.state.leagues.map((league) => {
            return (
                <tr key={league.id}>
                    <td>{league.id}</td>
                    <td>{league.name}</td>
                    <td>{league.type}</td>
                </tr>
            )
       })


         return(
            <div>
                <h1>Leagues</h1>
                <table>
                    <tbody>
                    {table_data}
                    </tbody>
                </table>
            </div>
        )
    }

}

export default DisplayLeagues;
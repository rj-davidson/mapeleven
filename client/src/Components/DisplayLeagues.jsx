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
        fetch("http://localhost:3000/leagues").then(
            (response) => response.json()
            ).then(data => {
                this.setState({ leagues: data });
                console.log(data);
            })
    }


    //.then(leagues => this.setState({ leagues: leagues }));
    render() {
        return (
            <div>
                <h1>Leagues</h1>
                <ul>
                    {this.state.leagues.map(league => (
                        <li key={league.id}></li>
                    ))}
                </ul>
            </div>
        );
    }
}

export default DisplayLeagues;
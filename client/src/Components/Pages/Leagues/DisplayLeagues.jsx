import React from "react";
//import table from "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/js/bootstrap.bundle.min.js";



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
                <table className='table table-striped'>
                    <caption className='caption-top'>Leagues</caption>
                    <thead>
                        <tr>
                            <th>League Id</th>
                            <th>League Name</th>
                            <th>League Type</th>
                        </tr>
                    </thead>

                    <tbody>
                    {table_data}
                    </tbody>
                </table>
            </div>
        )
    }

}
//className='table table-striped'
export default DisplayLeagues;
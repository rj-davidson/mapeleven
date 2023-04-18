import React from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import BarChart from "./Components/BarChart.jsx";
import NavigationBar from "./Components/NavigationBar.jsx";
//import {BrowserRouter, Router, Routes} from "react-router-dom";
import Home from "./Components/tabpages/Home.jsx";
import Fixtures from "./Components/tabpages/Fixtures.jsx";
import Players from "./Components/tabpages/Players.jsx";
import Clubs from "./Components/tabpages/Clubs.jsx";
import Leagues from "./Components/tabpages/Leagues.jsx";
import About from "./Components/tabpages/About.jsx";
//import NavigationBar from "./NavigationBar.jsx";
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';

function App() {
  //const [count, setCount] = useState(0);



  return (
      <div className='App'>
          <Router>
              <NavigationBar />
              <Routes>
                    <Route path='/' exact component={Home} />
                    <Route path='/fixtures' component={Fixtures} />
                    <Route path='/players' component={Players} />
                    <Route path='/clubs' component={Clubs} />
                    <Route path='/leagues' component={Leagues} />
                    <Route path='/about' component={About} />
              </Routes>
          </Router>
          <BarChart />
      </div>
  );
}
export default App;



        /*return (
            <div className="App">
                <div>
                    <a href="https://vitejs.dev" target="_blank">
                        <img src={viteLogo} className="logo" alt="Vite logo" />
                    </a>
                    <a href="https://reactjs.org" target="_blank">
                        <img src={reactLogo} className="logo react" alt="React logo" />
                    </a>
                </div>
                <h1>Vite + React</h1>
                <div className="card">
                    <button onClick={() => setCount((count) => count + 1)}>
                        count is {count}
                    </button>
                    <p>
                        Edit <code>src/App.jsx</code> and save to test HMR
                    </p>
                </div>
                <p className="read-the-docs">
                    Click on the Vite and React logos to learn more
                </p>
                <BarChart />


            </div>
            */





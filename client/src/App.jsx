import React from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import BarChart from "./BarChart.jsx";
import NavigationBar from "./Components/NavigationBar.jsx";
import {BrowserRouter, Router, Routes} from "react-router-dom";
//import NavigationBar from "./NavigationBar.jsx";
//import {BrowserRouter as Router, Switch, Route} from 'react-router-dom';

function App() {
  //const [count, setCount] = useState(0);
//this is a comment
  return (
      <div className='App'>
        <h1>Hello world</h1>
          <BrowserRouter>
              <NavigationBar />
          </BrowserRouter>
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





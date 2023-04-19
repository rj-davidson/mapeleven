import React from 'react'
import './App.css'
import BarChart from "./Components/BarChart.jsx";
import NavigationBar from "./Components/NavigationBar.jsx";
import PictureButton from "./Components/PictureButton.jsx";
import Home from "./Components/tabpages/Home";
import Fixtures from "./Components/tabpages/Fixtures";
import Players from "./Components/tabpages/Players";
import Clubs from "./Components/tabpages/Clubs";
import Leagues from "./Components/tabpages/Leagues";
import About from "./Components/tabpages/About";
//import NavigationBar from "./NavigationBar.jsx";
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';

function App() {
  //const [count, setCount] = useState(0);



  return (
      <div className='App'>
          <Router>
              <NavigationBar />
              <BarChart />
              <Routes>
                  <Route path='/fixtures' component={Fixtures} />
                  <Route path='/players' component={Players} />
                  <Route path='/clubs' component={Clubs} />
                  <Route path='/leagues' component={Leagues} />
                  <Route path='/about' component={About} />
                  <Route path='/' exact component = {Home}/>
              </Routes>
          </Router>
      </div>
  );
}
export default App;



        /*return (

              <NavigationBar />
              <BarChart />
              <h1>Hello World</h1>
          <!--<BarChart /> -->
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





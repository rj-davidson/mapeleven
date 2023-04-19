import React from 'react'
import './App.css'
import Fixtures from "./Components/Pages/Fixtures/Fixtures.tsx";
import Players from "./Components/Pages/Players/Players.jsx";
import Clubs from "./Components/Pages/Clubs/Clubs.jsx";
import Leagues from "./Components/Pages/Leagues/Leagues.jsx";
import About from "./Components/Pages/About/About";
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';

function App() {
  //const [count, setCount] = useState(0);



  return (
      <div className='App'>
          <Router>
              <Routes>
                  <Route path='/fixtures' element={<Fixtures/>} />
                  <Route path='/players' element={<Players/>} />
                  <Route path='/clubs' element={<Clubs/>} />
                  <Route path='/leagues' element={<Leagues/>} />
                  <Route path='/about' element={<About/>} />
                  <Route path='/' exact element = {<Fixtures/>}/>
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





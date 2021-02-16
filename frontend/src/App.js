import React from 'react';
import './App.css';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Navbar from './component/Navbar';
import Dashboard from './component/Guess';

function App() {
    return (
      <div className="wrapper">
        <Navbar></Navbar>
        <BrowserRouter>
          <Switch>
            <Route path="/Guess">
              <Dashboard />
            </Route>
          </Switch>
        </BrowserRouter>
      </div>
    );
  }
  
export default App;
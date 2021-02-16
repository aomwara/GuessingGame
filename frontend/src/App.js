import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import { BrowserRouter, Route, Switch } from 'react-router-dom';
import React, { useState } from 'react';
import useToken from './component/useToken';
import authCheck from './component/authCheck'

import Guess from './component/Guess';
import Login from './component/Login'

function App() {
  const { token, setToken } = useToken();
  var statusAuth = 0
  console.log("be-> "+statusAuth)

  authCheck(token).then(status => {
    statusAuth = status
    console.log("inf =>"+ statusAuth)
  });

  console.log("af-> "+statusAuth)
  if(!token || statusAuth) {
    return <Login setToken={setToken} />
  }

  return (
    <div className="wrapper">
      <BrowserRouter>
        <Switch>
          <Route path="/">
            <Guess />
          </Route>
          <Route path="/guess">
            <Guess />
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;

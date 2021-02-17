import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Redirect } from 'react-router'
import React, { useState, useEffect } from 'react';
import useToken from './component/useToken';
import authCheck from './component/authCheck'

import Guess from './component/Guess';
import Login from './component/Login'

function App() {
  const { token, setToken } = useToken();
  const [isAuth, setAuth] = useState(0);

  useEffect(() => {
    fetch('http://localhost:8888/authCheck', {
      method: 'GET',
      headers: {
        'Token': token
      },
    }).then(response => response.json())
    .then((responseData) => {
        setAuth(responseData.status)
    })
  }, []);

  if(!token || !isAuth){
    return <Login setToken={setToken} setAuth={setAuth}/>
  }
  
    return (
      <div className="wrapper">
        <BrowserRouter>
          <Switch>
            <Route exact path="/">
                <Redirect to="/guess" />
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

import './App.css';

import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Dashboard from './component/Guess';
import NavbarComponent from './component/Navbar';
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  return (
    <div className="wrapper">
      <NavbarComponent></NavbarComponent>
      <BrowserRouter>
        <Switch>
          <Route path="/guess">
            <Dashboard />
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;

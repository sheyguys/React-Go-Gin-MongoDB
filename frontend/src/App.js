import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Register from './Register/Register';
import Result from './Result/Result';
import './App.css';

function App() {
  return (
    <Router>
    <Switch>
      <Route path='/' exact={true} component={Register} />
      <Route path='/result' exact={true} component={Result} />
    </Switch>
  </Router>
  );
}

export default App;

import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Register from './Register/Register';
import './App.css';

function App() {
  return (
    <Router>
    <Switch>
      <Route path='/' exact={true} component={Register} />
    </Switch>
  </Router>
  );
}

export default App;

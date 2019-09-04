import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Register from './Register/Register';
import Result from './Result/Result';
import Edit from './Edit/Edit';
import Home from './Home/Home';
import './App.css';

function App() {
  return (
    <Router>
    <Switch>
      <Route path='/' exact={true} component={Home} />
      <Route path='/register' exact={true} component={Register} />
      <Route path='/result' exact={true} component={Result} />
      <Route path="/employee/:studentId" component={Edit} />
    </Switch>
  </Router>
  );
}

export default App;

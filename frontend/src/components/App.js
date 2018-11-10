import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import logo from '../assets/logo.svg';
import Login from './Login';
import Reset from './Reset';
import Signup from './Signup';

// eslint-disable-next-line react/prefer-stateless-function
class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <Switch>
          <Route exact path="/">
            <div className="app">
              <header className="app-header">
                <img src={logo} className="app-logo" alt="logo" />
                <h1 className="app-title">Welcome to React</h1>
              </header>
              <p className="app-intro">
                To get started, edit
                <code>
                  src/App.js
                </code>
                and save to reload.
              </p>
            </div>
          </Route>
          <Route exact path="/about">
            <p>
              Pinpoint is a versatile club application managment application
            </p>
          </Route>
          <Route exact path="/login" component={Login} />
          <Route exact path="/reset" component={Reset} />
          <Route exact path="/signup" component={Signup} />
        </Switch>
      </BrowserRouter>
    );
  }
}

export default App;

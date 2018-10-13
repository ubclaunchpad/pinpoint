import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import logo from './logo.svg';
import './App.scss';

// eslint-disable-next-line react/prefer-stateless-function
class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <Switch>
          <Route exact path="/">
            <div className="App">
              <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <h1 className="App-title">Welcome to React</h1>
              </header>
              <p className="App-intro">
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
        </Switch>
      </BrowserRouter>
    );
  }
}

export default App;

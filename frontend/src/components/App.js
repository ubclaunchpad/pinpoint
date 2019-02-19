import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Pinpoint from 'pinpoint-client';
import Cookies from 'universal-cookie';
import logo from '../assets/logo.svg';
import ClubsSelection from './ClubsSelection';
import Login from './Login';
import Reset from './Reset';
import Signup from './Signup';
import Navbar from './Navbar';
import CreateApplicationPeriod from '../scenes/ApplicationPeriod/create';


// eslint-disable-next-line react/prefer-stateless-function
class App extends Component {
  constructor(props, context) {
    super(props, context);
    this.state = {
      loggedIn: false,
    };
    this.cookies = new Cookies();
    this.attemptLogOut = this.attemptLogOut.bind(this);
    this.attemptLogIn = this.attemptLogIn.bind(this);
  }

  // Check cookie session here to keep logon whenever user reloads page
  componentDidMount() {
    if (this.cookies.get('userSession')) {
      this.setState({ loggedIn: true });
    }
  }

  attemptLogOut() {
    this.setState({ loggedIn: false });
    this.cookies.remove('userSession');
  }

  attemptLogIn(token) {
    this.setState({ loggedIn: true });
    this.cookies.set('userSession', token, { path: '/' });
  }

  render() {
    const { client } = this.props;
    const { loggedIn } = this.state;
    return (
      <BrowserRouter>
        <div>
          <Navbar loggedIn={loggedIn} attemptLogOut={this.attemptLogOut} />
          <Switch>
            <Route exact path="/">
              <div className="app">
                <header className="app-header">
                  <img src={logo} className="app-logo" alt="logo" />
                  <h1 className="app-title">Welcome to React</h1>
                </header>
                <p className="app-intro">
                  To get started, edit
                  <code>src/App.js</code>
                  and save to reload.
                </p>
              </div>
            </Route>
            <Route exact path="/about">
              <p>
                Pinpoint is a versatile club application managment application
              </p>
            </Route>
            <Route exact path="/me/clubs" component={ClubsSelection} />
            <Route
              exact
              path="/login"
              component={() => <Login client={client} attemptLogIn={this.attemptLogIn} />}
            />
            <Route exact path="/reset" component={Reset} />
            <Route exact path="/signup" component={() => <Signup client={client} />} />
            <Route exact path="/scenes/applicationperiod" component={CreateApplicationPeriod} />
          </Switch>
        </div>
      </BrowserRouter>
    );
  }
}

App.defaultProps = {
  client: new Pinpoint(),
};

App.propTypes = {
  client: PropTypes.instanceOf(Pinpoint.API),
};

export default App;

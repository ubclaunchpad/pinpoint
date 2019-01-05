import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Pinpoint from 'pinpoint-client';
// import logo from '../assets/logo.svg';
import Login from './Login';
import Reset from './Reset';
import Signup from './Signup';
import Navbar from './Navbar';

// eslint-disable-next-line react/prefer-stateless-function
class App extends Component {
  render() {
    const { client } = this.props;
    return (
      <BrowserRouter>
        <div>
          <div>
            <Navbar />
          </div>
          <Switch>
            <Route exact path="/">
              <p>
                Pinpoint is a versatile club application managment application
              </p>
            </Route>
            <Route exact path="/login" component={() => <Login client={client} />} />
            <Route exact path="/reset" component={Reset} />
            <Route exact path="/signup" component={Signup} />
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

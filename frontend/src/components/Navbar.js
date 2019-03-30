import React, { Component } from 'react';
import PropTypes from 'prop-types';
import logo from '../assets/pinpointlogo.png';

class Navbar extends Component {
  constructor(props, context) {
    super(props, context);
    this.state = {
      loggedIn: false,
    };
    this.loginStateDiv = this.loginStateDiv.bind(this);
  }

  componentDidMount() {
    const { loggedIn } = this.props;
    this.setState({ loggedIn: loggedIn });
  }

  // Check for log in state changed from prop
  componentDidUpdate() {
    let { loggedIn } = this.state;
    const prevLogState = loggedIn;
    ({ loggedIn } = loggedIn) = this.props;

    if (loggedIn !== prevLogState) {
      this.updateLogIn(loggedIn);
    }
  }

  updateLogIn(logState) {
    this.setState({ loggedIn: logState });
  }

  loginStateDiv() {
    const { attemptLogOut } = this.props;
    const { loggedIn } = this.state;
    if (loggedIn) {
      return (
        <span>
          <li><a className="margin-sides-s" href="/me/clubs">My Clubs</a></li>
          <li><button className="click-button button-small animate-button" type="submit" onClick={attemptLogOut}>Log Out</button></li>
        </span>
      );
    }
    return (
      <span>
        <li><a className="animate-link" href="/signup">Sign Up</a></li>
        <li><a className="animate-link" href="/login">Log In</a></li>
      </span>
    );
  }

  render() {
    return (
      <div className="pad-nav border-m">
        <nav className="navbar">
          <div className="pinpointlogo">
            <a href="/">pinpoint</a>
          </div>
          <ul className="margin-right-s">
            {this.loginStateDiv()}
          </ul>
        </nav>
      </div>
    );
  }
}

Navbar.propTypes = {
  loggedIn: PropTypes.bool,
  attemptLogOut: PropTypes.func.isRequired,
};

Navbar.defaultProps = {
  loggedIn: false,
};

export default Navbar;

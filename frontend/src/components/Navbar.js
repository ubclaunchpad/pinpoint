import React, { Component } from 'react';
import logo from '../assets/pinpointlogo.png';

class Navbar extends Component {
  constructor(props, context) {
    super(props, context);
    this.state = {
      loggedIn: false,
    };
    this.checkLogin = this.checkLogin.bind(this);
    this.attemptLogOut = this.attemptLogOut.bind(this);
  }

  // Return navbar layout based on whether user is logged in or not
  // For now default is logged in state for dev purpose
  checkLogin() {
    const { loggedIn } = this.state;
    if (loggedIn) {
      return (
        <span>
          <li><a className="margin-nav" href="/me/clubs">My Clubs</a></li>
          <li><button className="margin-nav click-button button-medium animate-button" type="submit" onClick={this.attemptLogOut}>Log Out</button></li>
        </span>
      );
    }
    return (
      <span>
        <li><a className="margin-nav" href="/signup">Sign Up</a></li>
        <li><a className="margin-nav" href="/login">Log In</a></li>
      </span>
    );
  }

  attemptLogOut() {
    this.setState({ loggedIn: false });
  }

  render() {
    return (
      <div className="pad-nav border-m">
        <nav className="animate-menu navbar">
          <div className="logo">
            <a href="/">
              <img src={logo} className="pinpointlogo" alt="logo" />
            </a>
          </div>
          <ul className="margin-nav">
            <li><a href="/">Home</a></li>
            {this.checkLogin()}
          </ul>
        </nav>
      </div>
    );
  }
}

export default Navbar;

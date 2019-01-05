import React, { Component } from 'react';
// import logo from './../assets'
import logo from '../assets/pinpointlogo.png';

class Navbar extends Component {
  constructor(props, context) {
    super(props, context);
    this.state = {
      loggedIn: false,
    };
    this.checkLogin = this.checkLogin.bind(this);
  }

  checkLogin() {
    const { loggedIn } = this.state;
    if (loggedIn) {
      return (
        <li><a href="/me/clubs">My Clubs</a></li>
      );
    }
    return (
      <span>
        <li><a href="/signup">Sign Up</a></li>
        <li><a href="/login">Log In</a></li>
      </span>
    );
  }

  render() {
    return (
      <div className="navbardiv">
        <nav id="navbar" className="animate-menu">
          <div className="nav-wrapper">
            <div className="logo">
              <a href="/">
                <img src={logo} className="pinpointlogo" alt="logo" />
              </a>
            </div>
            <ul id="menu">
              <li><a href="/">Home</a></li>
              {this.checkLogin()}
            </ul>

          </div>
        </nav>
      </div>
    );
  }
}

export default Navbar;

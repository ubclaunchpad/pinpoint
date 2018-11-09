import React, { Component } from 'react';

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: '',
    };
    this.updateTextFields = this.updateTextFields.bind(this);
    this.attemptLogin = this.attemptLogin.bind(this);
  }

  updateTextFields(e) {
    const { email, password } = this.state;
    console.log(email);
    console.log(password);
    const loginField = e.target.getAttribute('type');
    this.setState({ [loginField]: e.target.value });
  }

  // TODO once endpoint is set up, currently does nothing
  attemptLogin(e) {
    const { email, password } = this.state;
    console.log(email);
    console.log(password);
    console.log(e);
  }

  render() {
    return (
      <div className="flex-al-center">
        <div className="title magin-title">Sign In</div>
        <div className="flex-inlinegrid magin-form">
          <input className="input-box input-small" type="email" placeholder="Email" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="password" placeholder="Password" onChange={this.updateTextFields} />
        </div>
        <div>
          <input type="checkbox" />
          <span>Remember me</span>
        </div>
        <button className="click-button button-small animate-button magin-form" type="submit" onClick={this.attemptLogin}>Sign in</button>
        <div className="loginhelp">
          <a href="/login">Forgot Password?</a>
        </div>
        <div className="loginhelp">
          <span>Don&#x2019;t have an account? &nbsp;</span>
          <a href="/login">Sign up</a>
        </div>
      </div>
    );
  }
}

export default Login;

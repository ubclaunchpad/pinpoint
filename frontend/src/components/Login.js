import React, { Component } from 'react';

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: '',
    };
    this.updatetextfields = this.updatetextfields.bind(this);
  }

  updatetextfields(e) {
    const { email, password } = this.state;
    console.log(email);
    console.log(password);
    const loginfield = e.target.getAttribute('type');
    this.setState({ [loginfield]: e.target.value });
  }

  render() {
    return (
      <div className="login">
        <div className="title">Sign In</div>
        <div className="fields">
          <input type="email" placeholder="Email" onChange={this.updatetextfields} />
          <input type="password" placeholder="Password" onChange={this.updatetextfields} />
        </div>
        <div>
          <input type="checkbox" />
          <span>Remember me</span>
        </div>
        <button className="submit" type="submit">Sign in</button>
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

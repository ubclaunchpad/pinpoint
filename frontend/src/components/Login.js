import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Pinpoint from 'pinpoint-client';
import Notification from './Notification';

class Login extends Component {
  static contextTypes = {
    router: PropTypes.func.isRequired,
  }

  constructor(props, context) {
    super(props, context);
    this.state = {
      email: '',
      password: '',
      notification: null,
      userToken: null,
    };
    this.updateTextField = this.updateTextField.bind(this);
    this.attemptLogin = this.attemptLogin.bind(this);
  }

  updateTextField(e) {
    const loginField = e.target.getAttribute('name');
    this.setState({
      notification: null,
      [loginField]: e.target.value,
    });
  }

  // Checks user log in credentials
  async attemptLogin(token) {
    const { email, password, userToken } = this.state;
    const { client, attemptLogIn } = this.props;
    if (!email || !password) {
      this.setState({
        notification: {
          type: 'error',
          message: 'Please fill in all fields.',
        },
        userToken: token,
      });
      console.log(userToken); // temporarily bypass unused state for eslint
    } else {
      try {
        const resp = await client.login({ email, password });
        attemptLogIn(resp);
        const { router: { history } } = this.context;
        history.push('/');
      } catch {
        this.setState({
          notification: {
            type: 'error',
            message: 'Incorrect Credentials.',
          },
        });
      }
    }
  }

  render() {
    const { notification } = this.state;
    return (
      <div className="flex-al-center">
        <div className="title margin-title">Sign In</div>
        <Notification {...notification} />
        <div className="flex-inlinegrid margin-ends-xs">
          <input className="input-box input-small" type="email" name="email" placeholder="Email" onChange={this.updateTextField} />
          <input className="input-box input-small" type="password" name="password" placeholder="Password" onChange={this.updateTextField} />
        </div>

        <div>
          <input type="checkbox" />
          <span>Remember me</span>
        </div>
        <button className="click-button button-small animate-button margin-ends-xs" type="submit" onClick={this.attemptLogin}>Sign in</button>
        <div className="loginhelp">
          <a href="/reset">Forgot Password?</a>
        </div>
        <div className="loginhelp">
          <span>Don&#x2019;t have an account? &nbsp;</span>
          <a href="/signup">Sign up</a>
        </div>
      </div>
    );
  }
}

Login.propTypes = {
  client: PropTypes.instanceOf(Pinpoint.API).isRequired,
};

export default Login;

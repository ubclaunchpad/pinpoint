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
  async attemptLogin() {
    const { email, password } = this.state;
    const { client, setLoginState } = this.props;
    if (!email || !password) {
      this.setState({
        notification: {
          type: 'error',
          message: 'Please fill in all fields.',
        },
      });
    } else {
      try {
        setLoginState(await client.login({ email, password }));
        const { router: { history } } = this.context;
        history.push('/');
      } catch {
        this.setState({
          notification: {
            type: 'error',
            message: 'Incorrect Credentials',
          },
        });
      }
    }
  }

  render() {
    const { notification } = this.state;
    return (
      <div className="flex-al-center card margin-top-100 margin-sides-auto w-400">
        <div className="title card-title">Login</div>
        <Notification {...notification} />
        <div className="flex-inlinegrid margin-ends-xs margin-top-50">
          <input className="input-box input-small" type="email" name="email" placeholder="Email" onChange={this.updateTextField} />
          <input className="input-box input-small" type="password" name="password" placeholder="Password" onChange={this.updateTextField} />
        </div>

        <div>
          <input type="checkbox" />
          <span className="card-text">Remember me</span>
        </div>
        <button className="button-click button-small animate-button margin-ends-xs" type="submit" onClick={this.attemptLogin}>Sign in</button>
        <div className="loginhelp">
          <a href="/reset">Forgot Password?</a>
        </div>
        <div className="loginhelp pad-bot-25">
          <span className="card-text">Don&#x2019;t have an account? &nbsp;</span>
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

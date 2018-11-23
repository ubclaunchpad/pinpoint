import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Pinpoint from 'pinpoint-client';

class Login extends Component {
  static contextTypes = {
    router: PropTypes.func.isRequired,
  }

  constructor(props, context) {
    super(props, context);
    this.state = {
      email: '',
      password: '',
      failed: false,
      showmessage: false,
      message: { messageType: '', content: '' },
    };
    this.updateTextFields = this.updateTextFields.bind(this);
    this.attemptLogin = this.attemptLogin.bind(this);
  }

  updateTextFields(e) {
    const loginField = e.target.getAttribute('type');
    this.setState({ showmessage: false });
    this.setState({ [loginField]: e.target.value });
  }

  // TODO once endpoint is set up, currently does nothing
  async attemptLogin() {
    this.setState({ failed: false });
    const { email, password } = this.state;
    const { client } = this.props;

    if (!email || !password) {
      this.setState({ showmessage: true, message: { messageType: 'error', content: ' Please fill in all fields.' } });
      console.log('Please fill in all fields');
    }

    const resp = await client.login({ email, password });
    if (resp.status === 200) {
      const { router: { history } } = this.context;
      history.push('/');
    } else {
      this.setState({ failed: true });
    }
  }

  // TODO - Use this for displaying messages based on success
  // Function can be reused elsewhere, potentially put into a lib
  // content: string input
  // messageType: "info", "success", "warning", "error"
  generateMessage() {
    const { message, showmessage } = this.state;
    const colors = {
      info: 'blue',
      success: 'green',
      warning: 'orange',
      error: 'red',
    };

    if (showmessage) {
      return (
        <div className={`pad-ends-xs highlight-${colors[message.messageType]}`}>
          <i className="fa fa-times-circle" />
          {message.content}
        </div>
      );
    }
  }


  render() {
    const { failed } = this.state;
    return (
      <div className="flex-al-center">
        <div className="title margin-title">Sign In</div>
        { this.generateMessage() }
        <div className="flex-inlinegrid margin-ends-xs">
          <input className="input-box input-small" type="email" placeholder="Email" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="password" placeholder="Password" onChange={this.updateTextFields} />
        </div>

        { failed ? <div>Invalid credentials</div> : null }

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
          <a href="/login">Sign up</a>
        </div>
      </div>
    );
  }
}


Login.propTypes = {
  client: PropTypes.instanceOf(Pinpoint.API).isRequired,
};

export default Login;

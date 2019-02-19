import React, { Component } from 'react';
import Notification from './Notification';

class Signup extends Component {
  constructor(props, context) {
    super(props, context);
    this.state = {
      name: '',
      email: '',
      password: '',
      passwordConfirm: '',
      message: null,
      notification: {
        type: 'error',
        showNotification: false,
        message: "",
        transient: false,
      },
    };
    this.updateTextField = this.updateTextField.bind(this);
    this.attemptSignup = this.attemptSignup.bind(this);
    this.clickHandle = this.clickHandle.bind(this);
  }

  updateTextField(e) {
    const field = e.target.getAttribute('type');
    if (e.target.getAttribute('id') === 'confirm-password') {
      this.setState({ message: null, passwordConfirm: e.target.value });
    } else {
      this.setState({ message: null, [field]: e.target.value });
    }
  }

  // TODO once endpoint is set up, currently does nothing
  async attemptSignup() {
    const {
      email,
      password,
      name,
      passwordConfirm,
    } = this.state;

    const { client } = this.props;

    if (!email || !password || !name || !passwordConfirm) {
      this.setState({ message: { messageType: 'error', content: ' Please fill in all fields.' } });
    } else if (passwordConfirm !== password) {
      this.setState({ message: { messageType: 'error', content: ' Please make sure your passwords match.' } });
    } else {
      try {
        await client.createAccount({ email, name, password });
      } catch (e) {
        this.setState({ message: { messageType: 'error', content: ' Incorrect Credentials.' } });
      }
    }
  }

  // content: string input
  // messageType: "info", "success", "warning", "error"
  generateMessage() {
    const { message } = this.state;
    const colors = {
      info: 'blue',
      success: 'green',
      warning: 'orange',
      error: 'red',
    };

    const shape = {
      info: 'fa-info-circle',
      success: 'fa-check',
      warning: 'fa-warning',
      error: 'fa-times-circle',
    };

    if (message) {
      return (
        <div className={`pad-ends-xs highlight-${colors[message.messageType]}`}>
          <i className={`fa ${shape[message.messageType]}`} />
          {message.content}
        </div>
      );
    }
  }

  clickHandle() {
    this.setState({ notification: {
      type: "error",
      message: "this is an error msg",
      showNotification: true,
      transient: false,
    }});
  }

  render() {
    return (
      <div className="flex-al-center">
        <button onClick={this.clickHandle.bind(this)}>button</button>
        <Notification {...this.state.notification}></Notification>
        <div className="title margin-title">Sign-up</div>
        { this.generateMessage() }
        <div className="flex-inlinegrid margin-ends-xs">
          <input className="input-box input-small" type="name" placeholder="Name" onChange={this.updateTextField} />
          <input className="input-box input-small" type="email" placeholder="Email" onChange={this.updateTextField} />
          <input className="input-box input-small" type="password" placeholder="Password" onChange={this.updateTextField} />
          <input id="confirm-password" className="input-box input-small" type="password" placeholder="Confirm Password" onChange={this.updateTextField} />
        </div>
        <div className="margin-top-xs">
          <input type="checkbox" />
          <span>Send me e-mail updates</span>
        </div>
        <button className="click-button button-small animate-button margin-ends-xs" type="submit" onClick={this.attemptSignup}><a href="/login">Sign up</a></button>
        <div className="margin-top-xs">
          <span>Already have a pinpoint account? &nbsp;</span>
          <a href="/login">Sign In</a>
        </div>
      </div>
    );
  }
}

export default Signup;

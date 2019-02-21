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
      notification: {
        type: 'error',
        showNotification: false,
        message: '',
        transient: false,
      },
    };
    this.updateTextField = this.updateTextField.bind(this);
    this.attemptSignup = this.attemptSignup.bind(this);
  }

  updateTextField(e) {
    const field = e.target.getAttribute('name');
    this.setState({
      notification: {
        ...this.notification,
        showNotification: false,
      },
      [field]: e.target.value,
    });
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
      this.setState({
        notification: {
          type: 'error',
          message: 'Please fill in all fields.',
          showNotification: true,
          transient: false,
        },
      });
    } else if (passwordConfirm !== password) {
      this.setState({
        notification: {
          type: 'error',
          message: 'Please make sure your passwords match.',
          showNotification: true,
          transient: false,
        },
      });
    } else {
      try {
        await client.createAccount({ email, name, password });
      } catch (e) {
        this.setState({
          notification: {
            type: 'error',
            message: 'Failed to create a new account.',
            showNotification: true,
            transient: false,
          },
        });
      }
    }
  }

  render() {
    const { notification } = this.state;
    return (
      <div className="flex-al-center">
        <div className="title margin-title">Sign-up</div>
        <Notification {...notification} />
        <div className="flex-inlinegrid margin-ends-xs">
          <input className="input-box input-small" type="name" name="name" placeholder="Name" onChange={this.updateTextField} />
          <input className="input-box input-small" type="email" name="email" placeholder="Email" onChange={this.updateTextField} />
          <input className="input-box input-small" type="password" name="password" placeholder="Password" onChange={this.updateTextField} />
          <input className="input-box input-small" type="password" name="confirmPassword" placeholder="Confirm Password" onChange={this.updateTextField} />
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

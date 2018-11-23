import React, { Component } from 'react';

class Signup extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      email: '',
      password: '',
      confirmpassword: '',
      showmessage: false,
      message: { messageType: '', content: '' },
    };
    this.updateTextFields = this.updateTextFields.bind(this);
    this.attemptSignup = this.attemptSignup.bind(this);
  }

  updateTextFields(e) {
    const { name, confirmpassword } = this.state;
    const { email, password } = this.state;
    console.log(name, email, password, confirmpassword);
    const field = e.target.getAttribute('type');
    this.setState({ showmessage: false });
    if (field === 'password' && e.target.getAttribute('placeholder') !== 'Password') {
      this.setState({ confirmpassword: e.target.value });
    } else {
      this.setState({ [field]: e.target.value });
    }
  }

  // TODO once endpoint is set up, currently does nothing
  attemptSignup(e) {
    const { email, password } = this.state;
    const { name, confirmpassword } = this.state;
    console.log(email, password, name, confirmpassword, e);

    if (!email || !password || !name || !confirmpassword) {
      this.setState({ showmessage: true, message: { messageType: 'error', content: ' Please fill in all fields.' } });
    } else if (confirmpassword !== password) {
      this.setState({ showmessage: true, message: { messageType: 'error', content: ' Please make sure your passwords match.' } });
    } else {
      // TODO Send signup information to backend here
      this.setState({ showmessage: true, message: { messageType: 'success', content: ' Success!' } });
    }
  }

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
    return (
      <div className="flex-al-center">
        <div className="title margin-title">Signup</div>
        { this.generateMessage() }
        <div className="flex-inlinegrid margin-ends-xs">
          <input className="input-box input-small" type="name" placeholder="Name" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="email" placeholder="Email" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="password" placeholder="Password" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="password" placeholder="Confirm Password" onChange={this.updateTextFields} />
        </div>
        <div className="margin-top-xs">
          <input type="checkbox" />
          <span>Send me e-mail updates</span>
        </div>
        <button className="click-button button-small animate-button margin-ends-xs" type="submit" onClick={this.attemptSignup}>Sign up</button>
        <div className="margin-top-xs">
          <span>Already have a pinpoint account? &nbsp;</span>
          <a href="/login">Sign In</a>
        </div>
      </div>
    );
  }
}

export default Signup;

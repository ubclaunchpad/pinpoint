import React, { Component } from 'react';

class Signup extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      email: '',
      password: '',
      confirmpassword: '',
    };
    this.updateTextFields = this.updateTextFields.bind(this);
    this.attemptSignup = this.attemptSignup.bind(this);
  }

  updateTextFields(e) {
    const { name, confirmpassword } = this.state;
    const { email, password } = this.state;
    console.log(name, email, password, confirmpassword);
    const field = e.target.getAttribute('type');

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

    if (confirmpassword !== password) {
      console.log('Please make sure password is same as your confirmation');
    } else {
      // TODO Send signup information to backend here
      console.log('Success');
    }
  }

  generateError(error) {
    const { email, password } = this.state;
    const { name, confirmpassword } = this.state;
    console.log(email, password, name, confirmpassword, e);

    return (
      <div>{error}</div>
    );
  }

  render() {
    return (
      <div className="flex-al-center">
        <div className="title margin-title">Signup</div>
        <div className="flex-inlinegrid margin-top-xs margin-bottom-xs">
          <input className="input-box input-small" type="name" placeholder="Name" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="email" placeholder="Email" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="password" placeholder="Password" onChange={this.updateTextFields} />
          <input className="input-box input-small" type="password" placeholder="Confirm Password" onChange={this.updateTextFields} />
        </div>
        <div className="margin-top-xs">
          <input type="checkbox" />
          <span>Send me e-mail updates</span>
        </div>
        <button className="click-button button-small animate-button margin-top-xs margin-bottom-xs" type="submit" onClick={this.attemptSignup}>Sign up</button>
        <div className="margin-top-xs">
          <span>Already have a pinpoint account? &nbsp;</span>
          <a href="/login">Sign In</a>
        </div>
      </div>
    );
  }
}

export default Signup;

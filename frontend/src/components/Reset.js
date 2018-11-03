import React, { Component } from 'react';

class Reset extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
    };
    this.updatetextfields = this.updatetextfields.bind(this);
  }

  updatetextfields(e) {
    const { email } = this.state;
    console.log(email);
    const loginfield = e.target.getAttribute('type');
    this.setState({ [loginfield]: e.target.value });
  }

  render() {
    return (
      <div className="reset">
        <div className="title">Reset Password</div>
        <p>Enter the e-mail linked to your account</p>
        <p>We&#x2019;ll send you an e-mail with a link to reset your password.</p>
        <div className="fields">
          <input type="email" placeholder="E-mail address" onChange={this.updatetextfields} />
        </div>
        <div>
          <button className="submit" type="submit">Send reset link</button>
        </div>
        <div className="loginhelp">
          <a href="/login">Back to login</a>
        </div>
      </div>
    );
  }
}

export default Reset;

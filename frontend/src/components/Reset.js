import React, { Component } from 'react';

class Reset extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
    };
    this.updatetextfields = this.updatetextfields.bind(this);
    this.attemptSendReset = this.attemptSendReset.bind(this);
  }

  updatetextfields(e) {
    const { email } = this.state;
    console.log(email);
    const infoField = e.target.getAttribute('type');
    this.setState({ [infoField]: e.target.value });
  }

  // TODO, Currently does nothing
  attemptSendReset(e) {
    const { email } = this.state;
    console.log(email);
    console.log(e);
  }

  render() {
    return (
      <div className="flex-al-center">
        <div className="title margin-title">Reset Password</div>
        <p>Enter the e-mail linked to your account</p>
        <p>We&#x2019;ll send you an e-mail with a link to reset your password.</p>
        <div className="flex-inlinegrid margin-top margin-bottom">
          <input className="input-box input-small" type="email" placeholder="E-mail address" onChange={this.updatetextfields} />
        </div>
        <div>
          <button className="click-button button-medium animate-button margin-top margin-bottom" type="submit" onClick={this.attemptSendReset}>Send reset link</button>
        </div>
        <div>
          <a href="/login">Back to login</a>
        </div>
      </div>
    );
  }
}

export default Reset;

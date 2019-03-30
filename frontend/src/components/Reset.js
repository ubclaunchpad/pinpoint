import React, { Component } from 'react';

class Reset extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
    };
    this.updateTextFields = this.updateTextFields.bind(this);
    this.attemptSendReset = this.attemptSendReset.bind(this);
  }

  updateTextFields(e) {
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
      <div className="flex-al-center card margin-top-100 margin-sides-auto w-500">
        <div className="title card-title">Reset Password</div>
        <p className="card-text">Enter the e-mail linked to your account</p>
        <p className="card-text">We&#x2019;ll send you an e-mail with a link to reset your password.</p>
        <div className="flex-inlinegrid margin-top-xs margin-bottom-xs">
          <input className="input-box input-small" type="email" placeholder="E-mail address" onChange={this.updateTextFields} />
        </div>
        <div className="pad-ends-25">
          <button className="button-click button-medium animate-button margin-top-xs margin-bottom-xs" type="submit" onClick={this.attemptSendReset}>Send reset link</button>
        </div>
        <div>
          <a href="/login">Back to login</a>
        </div>
      </div>
    );
  }
}

export default Reset;

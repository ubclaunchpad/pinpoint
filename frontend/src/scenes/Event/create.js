import React, { Component } from 'react';
import Notification from '../../components/Notification';
class CreateEvent extends Component {
  constructor(props, context) {
    super(props, context);
    this.state = {
      eventname: '',
      notification: null,
    };
    this.updateTextFields = this.updateTextFields.bind(this);
  }

  updateTextFields(e) {
    const infoField = e.target.getAttribute('name');
    this.setState({ [infoField]: e.target.value });
  }

  async checkEventName() {
    const { eventname } = this.state;

    if (!eventname) {
      this.setState({
        notification: {
          type: 'error',
          message: 'Event name cannot be empty',
        },
      });
    }
  }

  render() {
    const { notification, eventname } = this.state;
    return (
      <div>
        <div className="card card-smallpad card-box margin-top-100 margin-sides-auto w-800">
          <div>
            <a href="/login">Go Back</a>
            <div className="card-align-right"><a href="/login">X</a></div>
          </div>
          <div className="pad-left-100 title card-title">Create Event</div>
          <Notification {...notification} />
          <h2 className="pad-left-100 fw-normal card-description pad-bot-m">Add a point of data entry to your application period</h2>
          <h2 className="pad-left-100 flex-ai-start pad-top- fw-normal card-text">Event Name</h2>
          <div className="pad-left-100 flex-inlinegrid margin-top-xs margin-bottom-xs pad-bot-m">
            <input className="input-box input-large" name="eventname" type="eventname" placeholder="Eg. Launch Pad Interview Notes" onChange={this.updateTextFields} />
          </div>
          <div className="pad-ends-25 flex-al-center">
            <button className="button-click button-small animate-button margin-ends-xs margin-right-s" type="submit" eventname={eventname} onClick={this.checkEventName}><a href="/event/type">Next</a></button>
            <a href="/login" className="card-underline">Cancel</a>
          </div>
        </div>
      </div>
    );
  }
}

export default CreateEvent;

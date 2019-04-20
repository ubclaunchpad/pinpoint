import React, { Component } from 'react';
import Notification from '../../components/Notification';
class CreateEvent extends Component {
  constructor(props, context) {
    super(props, context);
    this.state = {
      eventtype: '',
      notification: null,
    };
    this.updateTextFields = this.updateTextFields.bind(this);
  }

  updateTextFields(e) {
    const infoField = e.target.getAttribute('name');
    this.setState({ [infoField]: e.target.value });
    this.attemptCreateEvent();
  }

  // TODO once endpoint is set up, currently does nothing
  async attemptCreateEvent() {
    const {
      eventtype,
    } = this.state;
    const { client, eventname } = this.props;

    if (!eventtype) {
      this.setState({
        notification: {
          type: 'error',
          message: 'Event type cannot be empty',
        },
      });
    } else {
      try {
        await client.createEvent({ eventname, eventtype });
      } catch (e) {
        this.setState({
          notification: {
            type: 'error',
            message: 'Failed to create a new event.',
          },
        });
      }
    }
  }

  render() {
    const { notification } = this.state;
    return (
      <div>
        <div className="card card-smallpad card-box margin-top-100 margin-sides-auto w-600">
          <div>
            <a href="/event">Go Back</a>
            <div className="card-align-right"><a href="/event">Skip</a></div>
          </div>
          <div className="pad-left-50 title card-title">Event type</div>
          <Notification {...notification} />
          <h2 className="pad-left-50 fw-normal card-description pad-bot-xs">Pick an event type</h2>
          <div className="pad-ends-25 flex-al-center">
            <button className="button-click button-small animate-button margin-ends-xs margin-right-s" type="submit" name="email" onClick={this.updateTextFields}><a href="/event/type">Email</a></button>
            <button className="button-click button-small animate-button margin-ends-xs margin-right-s" type="submit" name="form" onClick={this.updateTextFields}><a href="/event/type">Form</a></button>
            <button className="button-click button-small animate-button margin-ends-xs margin-right-s" type="submit" name="interview" onClick={this.updateTextFields}><a href="/event/type">Interview</a></button>
            <button className="button-click button-small animate-button margin-ends-xs margin-right-s" type="submit" name="notes" onClick={this.updateTextFields}><a href="/event/type">Notes</a></button>
          </div>
        </div>
      </div>
    );
  }
}

export default CreateEvent;

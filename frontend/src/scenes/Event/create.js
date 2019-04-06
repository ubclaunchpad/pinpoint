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
    const infoField = e.target.getAttribute('id');
    this.setState({ [infoField]: e.target.value });
  }

  // TODO once endpoint is set up, currently does nothing
  async attemptCreateEvent() {
    const {
      eventname,
    } = this.state;

    const { client } = this.props;

    if (!eventname) {
      this.setState({
        notification: {
          type: 'error',
          message: 'Event name cannot be empty',
        },
      });
    } else {
      try {
        await client.createEvent({ eventname });
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
        <div className="flex-al-center card margin-top-100 margin-sides-auto w-800">
          <div className="margin-left-s"><a href="/login">Go Back</a></div>
          <div className="title card-title">Create Event</div>
          <Notification {...notification} />
          <h2 className="fw-normal card-description">Add a point of data entry to your application period</h2>
          <h2 className="flex-ai-start pad-top-xxl fw-normal card-text">Event Name</h2>
          <div className="flex-inlinegrid margin-top-xs margin-bottom-xs">
            <input className="input-box input-large" name="eventname" type="eventname" placeholder="Eg. Launch Pad Interview Notes" onChange={this.updatetextfields} />
          </div>
          <button className="button-click button-small animate-button margin-ends-xs margin-right-s" type="submit" onClick={this.attemptCreateEvent}><a href="/event/type">Next</a></button>
          <a href="/login">Cancel</a>
        </div>
      </div>
    );
  }
}

export default CreateEvent;

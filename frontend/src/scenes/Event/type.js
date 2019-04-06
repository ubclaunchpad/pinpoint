import React, { Component } from 'react';

class CreateEvent extends Component {
  constructor(props) {
    super(props);
    this.updateTextFields = this.updateTextFields.bind(this);
  }

  updateTextFields(e) {
    const infoField = e.target.getAttribute('id');
    this.setState({ [infoField]: e.target.value });
  }

  render() {
    return (
      <div>
        <div className="flex-al-center card margin-top-100 margin-sides-auto w-600">
          <div className="title card-title">Create Event</div>
          <h2 className="fw-normal card-description">Add a point of data entry to your application period</h2>
          <h2 className="flex-ai-start pad-top-xxl fw-normal card-text">Event Name</h2>
          <div className="flex-inlinegrid margin-top-xs margin-bottom-xs">
            <input className="input-box input-large" name="eventname" type="eventname" placeholder="Eg. Launch Pad Interview Notes" onChange={this.updatetextfields} />
          </div>
          <button className="button-click button-small animate-button margin-ends-xs" type="submit" onClick={this.attemptCreateEvent}><a href="/event/type">Next</a></button>
          <h2 className="fw-normal card-description">Cancel</h2>
        </div>
      </div>
    );
  }
}

export default CreateEvent;

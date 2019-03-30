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
        <div className="flex dir-col pad-left-300">
          <div className="title margin-title">Create Event</div>
          <h2 className="fw-normal">Add a point of data entry to your application period</h2>
          <h2 className="flex-ai-start pad-top-xxl fw-normal">Event Name </h2>
          <div className="flex-inlinegrid margin-top-xs margin-bottom-xs">
            <input className="input-box input-large" id="applications" placeholder="Eg. Launch Pad Interview Notes" onChange={this.updatetextfields} />
          </div>
        </div>
      </div>
    );
  }
}

export default CreateEvent;

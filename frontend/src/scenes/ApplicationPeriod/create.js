import React, { Component } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

class CreateApplicationPeriod extends Component {
  constructor(props) {
    super(props);
    this.state = {
      startDate: new Date(),
      endDate: new Date(),
      hasError: false,
    };
    this.updateTextFields = this.updateTextFields.bind(this);
    this.handleChangeStart = this.handleChangeStart.bind(this);
    this.handleChangeEnd = this.handleChangeEnd.bind(this);
  }

  handleChangeStart(date) {
    const { endDate } = this.state;
    if (endDate < date) {
      this.setState({ hasError: true });
      const { hasError } = this.state;
      console.log('Start date cannot be greater than end date!', hasError);
    } else {
      this.setState({
        startDate: date,
      });
    }
  }

  handleChangeEnd(date) {
    const { startDate } = this.state;
    if (startDate > date) {
      this.setState({ hasError: true });
      const { hasError } = this.state;
      console.log('Start date cannot be greater than end date!', hasError);
    } else {
      this.setState({ endDate: date });
    }
  }

  updateTextFields(e) {
    const infoField = e.target.getAttribute('id');
    this.setState({ [infoField]: e.target.value });
  }

  render() {
    const { startDate } = this.state;
    const { endDate } = this.state;
    return (
      <div>
        <div className="flex dir-col pad-left-300">
          <div className="title margin-title">Create Application Period</div>
          <h2 className="flex-ai-start pad-top-xxl fw-normal">Applications</h2>
          <div className="flex-inlinegrid margin-top-xs margin-bottom-xs">
            <input className="input-box input-large" id="applications" placeholder="Application" onChange={this.updatetextfields} />
          </div>
          <h2 className="pad-top-xxl fw-normal">Allow applicants to apply</h2>
        </div>
        <div className="flex dir-row pad-left-300">
          <div>
            <div>From</div>
            <DatePicker
              className="input-box input-small"
              selected={startDate}
              onChange={this.handleChangeStart}
            />
          </div>
          <div className="pad-left-m">
            <div>To</div>
            <DatePicker
              className="input-box input-small"
              selected={endDate}
              onChange={this.handleChangeEnd}
            />
          </div>
        </div>
      </div>
    );
  }
}

export default CreateApplicationPeriod;

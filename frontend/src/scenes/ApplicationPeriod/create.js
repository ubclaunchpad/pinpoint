import React, { Component } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

class ApplicationPeriod extends Component {
  constructor(props) {
    super(props);
    this.state = {
      startDate: new Date(),
      endDate: new Date(),
      hasError: false,
    };
    this.updatetextfields = this.updatetextfields.bind(this);
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

  updatetextfields(e) {
    const infoField = e.target.getAttribute('type');
    this.setState({ [infoField]: e.target.value });
  }

  render() {
    const { startDate } = this.state;
    const { endDate } = this.state;
    return (
      <div>
        <div className="flex dir-col pad-left-xxxl">
          <div className="title margin-title">Create Application Period</div>
          <div className="heading1 flex-ai-start pad-top-xxl">Applications</div>
          <div className="flex-inlinegrid margin-top-xs margin-bottom-xs">
            <input className="input-box input-large" type="applications" placeholder="Application" onChange={this.updatetextfields} />
          </div>
          <div className="heading1 pad-top-xxl">Allow applicants to apply</div>
        </div>
        <div className="flex dir-row pad-left-xxxl">
          <div>
            <div>From</div>
            <DatePicker
              selected={startDate}
              onChange={this.handleChangeStart}
            />
          </div>
          <div className="pad-left-m">
            <div>To</div>
            <DatePicker
              selected={endDate}
              onChange={this.handleChangeEnd}
            />
          </div>
        </div>
      </div>
    );
  }
}

export default ApplicationPeriod;

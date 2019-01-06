import React, { Component } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

class ApplicationPeriod extends Component {
  constructor(props) {
    super(props);
    this.state = {
      application: '',
      startDate: new Date(),
      endDate: new Date(),
    };
    this.updatetextfields = this.updatetextfields.bind(this);
    this.handleChangeStart = this.handleChangeStart.bind(this);
    this.handleChangeEnd = this.handleChangeEnd.bind(this);
  }

  handleChangeStart(date) {
    this.setState({
      startDate: date,
    });
  }

  handleChangeEnd(date) {
    this.setState({
      endDate: date,
    });
  }

  updatetextfields(e) {
    const { application } = this.state;
    console.log(application);
    const infoField = e.target.getAttribute('type');
    this.setState({ [infoField]: e.target.value });
  }

  render() {
    //                      { startDate, ... }
    // this.state.date.start
    // const { date: { start } = {} } = this.state
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

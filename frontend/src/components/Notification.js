import React, { Component } from 'react';
import PropTypes from 'prop-types';

class Notification extends Component {
  getColorClass() {
    const { type } = this.props;
    const colors = {
      info: 'blue',
      success: 'green',
      warning: 'orange',
      error: 'red',
    };
    return colors[type];
  }

  getIconClass() {
    const { type } = this.props;
    const icons = {
      info: 'fa-info-circle',
      success: 'fa-check',
      warning: 'fa-warning',
      error: 'fa-times-circle',
    };
    return icons[type];
  }

  render() {
    const { showNotification, message } = this.props;
    if (showNotification) {
      return (
        <div className={`pad-ends-xs highlight-${this.getColorClass()}`}>
          <i className={`fa ${this.getIconClass()}`} />
&nbsp;
          {message}
        </div>
      );
    }
    return <span />;
  }
}

Notification.propTypes = {
  type: PropTypes.oneOf(['info', 'success', 'warning', 'error']),
  showNotification: PropTypes.bool,
  message: PropTypes.string,
  transient: PropTypes.bool,
};

Notification.defaultProps = {
  type: 'info', // info, success, warning, error
  showNotification: false,
  message: '',
  transient: false,
};

export default Notification;

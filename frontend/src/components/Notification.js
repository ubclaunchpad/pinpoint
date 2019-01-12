import React, { Component } from 'react';

class Notification extends Component {
	render() {
		if (this.props.showNotification) {
			return (
				<div className="notification">
					<p>this.props.message</p>
				</div>
			);
		}
		return <span></span>;
	}
}

Notification.defaultProps = {
	type: 'info', // info, success, warning, error
	showNotification: false,
	message: "",
	transient: false,
}

export default Notification;

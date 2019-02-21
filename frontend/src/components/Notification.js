import React, { Component } from 'react';

class Notification extends Component {
	getColorClass() {
		let colors = {
			info: 'blue',
			success: 'green',
			warning: 'orange',
			error: 'red',
		};
		return colors[this.props.type];
	}

	getIconClass() {
		let icons = {
			info: 'fa-info-circle',
			success: 'fa-check',
			warning: 'fa-warning',
			error: 'fa-times-circle',
		};
		return icons[this.props.type];
	}

	render() {
		if (this.props.showNotification) {
			return (
				<div className={`pad-ends-xs highlight-${this.getColorClass()}`}>
					<i className={`fa ${this.getIconClass()}`} />&nbsp;
					{this.props.message}
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

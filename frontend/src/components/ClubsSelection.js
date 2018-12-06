import React, { Component } from 'react';
require('../assets/ubclaunchpad-logo.png');

class ClubsSelection extends Component {
  constructor(props) {
    super(props);
    this.state = { clubs: [] };
    this.generateclubs = this.generateclubs.bind(this);
  }

  // Mock loading of database object into component
  componentDidMount() {
    this.setState({ clubs: [{ name: 'UBC Launchpad 1', imagePath: '../assets/ubclaunchpad-logo.png' }, { name: 'UBC Launchpad 2', imagePath: '../assets/ubclaunchpad-logo.png' }] });
  }

  // TODO Replace unique key with database id of club
  generateclubs() {
<<<<<<< HEAD
=======
    console.log(this.state);
>>>>>>> 815f46f034c074b70b77bef94c55ea0163bdffb5
    const { clubs } = this.state;
    return clubs.map((club) => (
      <li className="margin-left-s margin-right-s" key={Math.random() * 10000}>
        <img className="club-img-l" src={require('../assets/ubclaunchpad-logo.png')} alt={club.name} />
        <p className="textwrap">
          {club.name}
        </p>
      </li>
    ));
  }

  render() {
    return (
      <div className="flex-al-center">
        <div className="title margin-title"> Your Clubs </div>
        <ul className="flex-inline margin-top-xs margin-bottom-xs">
          {this.generateclubs()}
          <li className="margin-left-s margin-right-s" key={Math.random() * 10000}>
            <img className="club-img-l" src={require('../assets/newclub.png')} alt="new club" />
            <p className="textwrap">
              Create new club
            </p>
          </li>
        </ul>
      </div>
    );
  }
}

export default ClubsSelection;

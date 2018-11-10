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
    console.log(this.state);
    console.log('test');
    console.log();

    const { clubs } = this.state;
    const clublist = clubs.map((club) => (
      <li className="margin-left-s margin-right-s" key={Math.random() * 10000}>
        <img className="club-img-l" src={club.imagePath} alt={club.name} />
        <p className="textwrap">
          {club.name}
        </p>
      </li>
    ));

    return <ul className="flex-inline margin-top-xs margin-bottom-xs">{ clublist }</ul>;
  }

  render() {
    return (
      <div className="flex-al-center">
        <div className="title margin-title"> Your Clubs </div>
        <ul className="flex-inline margin-top-xs margin-bottom-xs">
          {this.generateclubs()}
          <li>
            <img className="club-img-l pad-left-m" src={require('../assets/newclub.png')} alt="new club" />
            <p> Create a new club </p>
          </li>
        </ul>
      </div>
    );
  }
}

export default ClubsSelection;

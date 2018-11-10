import React from 'react';
import '../styles/components/clubselection.scss';

const clubs = [{ name: 'UBC Launchpad', imagePath: './ubclaunchpad-logo.png' }];
export default function ClubsSelection() {
  return (
    <div className="flex ai-center jc-center">
      <h1> Your Clubs </h1>
      <ul className="flex dir-column">
        {clubs.map((club) => (
          <li>
            <img src={require('../assets/ubclaunchpad-logo.png')} alt="clubImage" />
            <p className="clubname">
              {club.name}
            </p>
          </li>
        ))}
        <li>
          <img className="image pad-left-m" src={require('../assets/newclub.png')} alt="new club" />
          <button type="button">
            Create new club
          </button>
        </li>
      </ul>
    </div>
  );
}

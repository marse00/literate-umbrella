import React from "react";
import "../../App.css";
import {Link} from "react-router-dom"


function Nav() {
  const linksColor = {
    color: 'white',
    textDecoration: 'none'
  };

  return (
    <nav>
        <h3> Logo </h3>
        <ul className="nav-dirs">
        <Link style={linksColor} to="/about">
          <li>About</li>
        </Link>
        <Link style={linksColor} to="/profile">
          <li>Profile</li>
        </Link>
        <Link style={linksColor} to="/login">
          <li>Login</li>
        </Link>
        </ul>
    </nav>
  );
}

export default Nav;

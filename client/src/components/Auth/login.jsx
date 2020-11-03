import React, { Component } from "react";
import axios from "axios";

let config = {
  headers: {
    withCredentials: true,
    "Content-Type": "text/html",
  },
};

class LoginForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
    };
  }
  render() {
    return (
      <React.Fragment>
        <h1>Login</h1>
        <form onSubmit={this.handleSubmit}>
          <input
            type="username"
            name="username"
            value={this.state.username}
            onChange={this.handleChange}
            required
          />
          <input
            type="password"
            name="password"
            value={this.state.password}
            onChange={this.handleChange}
            required
          />
          <button type="submit">Register</button>
        </form>
      </React.Fragment>
    );
  }
  handleChange = (event) => {
    console.log(event.target);
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  handleSubmit = () => {
    const username = this.state.username;
    const password = this.state.password;
    //const thought = this.state.thought;

    axios
      .post(
        "http://localhost:8080/",

        {
          User: {
            Username: username,
            Password: password,
            //Thought: thought,
          },
        },

        config
      )
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
  };
}

export default LoginForm;

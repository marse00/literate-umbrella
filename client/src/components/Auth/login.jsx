import React, { Component } from "react";
import axios from "axios";



class LoginForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
      usernameError: "",
      passwordError: ""
    };
  }
  o
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
           <div>{this.state.usernameError} </div>
          <input
            type="password"
            name="password"
            value={this.state.password}
            onChange={this.handleChange}
            required
          />
          <div>{this.state.passwordError} </div>
          <button type="submit">Login</button>
        </form>
      </React.Fragment>
    );
  }
  handleChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };



  handleSubmit = (event) => {
    event.preventDefault()
    const username = this.state.username;
    const password = this.state.password;

    

    let config = {
      headers: {
        withCredentials: true,
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
    };
    axios
      .post(
        "/login",
        {
          User: {
            Username: username,
            Password: password,
          },
        },
        config
      )
      
      .then(response => {
        if (response.status === 200) {
          console.log("response status: " + response.status);
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  };
}

export default LoginForm;

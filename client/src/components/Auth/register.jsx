import React, { Component } from "react";
import axios from "axios";


class SignUpForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
      password_confirmation: "",
      email: "",
      usernameError: "",
      passwordError: "",
      password_confirmationError: "",
      emailError: ""
    };
  }
  render() {
    return (
      <React.Fragment>
        <h1>Register</h1>
        <form onSubmit={this.handleSubmit}>
        <h6>Username:</h6>
          <input
            type="username"
            name="username"
            value={this.state.username}
            onChange={this.handleChange}
            required
          />
          <div>{this.state.usernameError} </div>
          <h6>E-mail:</h6>
          <input
            type="email"
            name="email"
            value={this.state.email}
            onChange={this.handleChange}
            required
          />
          <div>{this.state.emailError} </div>
          <h6>Password:</h6>
          <input
            type="password"
            name="password"
            value={this.state.password}
            onChange={this.handleChange}
            required
          />
          <div>{this.state.passwordError} </div>
          <h6>Password confirmation:</h6>
          <input
            type="password"
            name="password_confirmation"
            value={this.state.password_confirmation}
            onChange={this.handleChange}
            required
          />
          <div>{this.state.password_confirmationError} </div>
          <button type="submit">Register</button>
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
    event.preventDefault();

    const username = this.state.username;
    const password = this.state.password;
    const password_confirmation = this.state.password_confirmation;
    const email = this.state.email;

    if(password !== password_confirmation)
    {
      this.setState({password_confirmationError: "The passwords need to match"})
      
    }
    else
    {
      let config = {
        headers: {
          withCredentials: true,
        },
      };
  
  
      axios
        .post(
          "/register",
  
          {
            User: {
              Username: username,
              Password: password,
              Email: email
            },
          },
  
          config
        )
        .then(function (response) {
          console.log( response);
        })
        .catch(function (error) {
          console.log("Axios error: "+error);
        });
    }
    
   
  };
}

export default SignUpForm;

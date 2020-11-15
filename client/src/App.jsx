import React from "react";
import "./App.css";
import LoginForm from "./components/Auth/Login";
import Nav from "./components/Links/Nav"
import Home from "./components/Links/Home"
import About from "./components/Links/About"
import Watch from "./components/Links/Watch"
import Profile from "./components/Links/Profile"
import {Route, Switch, BrowserRouter as Router} from "react-router-dom"

function App() {
  return (
    <Router>
      <div className="App">
          <Nav />
          <Switch>
            <Route path="/" exact component={Home}/>
            <Route path="/about" component={About}/>
            <Route path="/watch" component={Watch}/>
            <Route path="/profile" component={Profile}/>
            <Route path="/login" component={LoginForm}/>
          </Switch>
      </div>
    </Router>
  );
}

export default App;

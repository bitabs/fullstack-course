import React, { Component } from "react"
import "./App.css"
import {
  BrowserRouter as Router,
  Route
} from "react-router-dom"

import Home from "components/pages/Home"
import Signup from "components/pages/Signup"

class App extends Component {
  render() {
    return (
      <Router>
        <Route exact path="/" component={Home}/>
        <Route path="/signup" component={Signup}/>
      </Router>
    );
  }
}

export default App;

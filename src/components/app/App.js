import React, { Component } from "react";
import { connect } from 'react-redux'
import { fetchPosts } from "../../actions/postActions";
import "./App.css";
import Home from 'src/components/pages/Home/Home'

class App extends Component {

  componentDidMount() {
    this.props.fetchPosts()
  }

  render() {
    return (
      <Home posts={this.props.posts} />
    );
  }
}

const mapStateToProps = state => ({
  ...state
});

const mapDispatchToProps = dispatch => ({
  fetchPosts: () => dispatch(fetchPosts())
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import Footer from './Footer.js';
import  TopMenu from './TopMenu.js';

class Home extends Component {
  render() {
    return (
<div>
      <img src="burger_pic.png" />
</div>
    );
  }
}

export default withRouter(Home);


import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';

class Signup extends Component {
  render() {
    return (  
        <div>
            <span>Signup</span>
        </div>
    );
  }
}

export default withRouter(Signup);


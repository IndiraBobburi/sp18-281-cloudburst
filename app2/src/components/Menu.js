import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';

class Menu extends Component {
  render() {
    return (  
        <div>
            <span>Menu</span>
        </div>
    );
  }
}

export default withRouter(Menu);


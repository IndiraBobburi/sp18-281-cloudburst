import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import {BrowserRouter } from 'react-router-dom';
import MainComponent from './components/MainComponent';

class App extends Component {
  render() {
    return (
      <div className="App">
          <BrowserRouter>
              <MainComponent/>
          </BrowserRouter>
      </div>
    );
  }
}

export default App;

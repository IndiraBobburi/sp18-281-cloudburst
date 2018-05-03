import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import * as API from '../api/API.js';
class RestaurantMenu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            menuItems: []
        }
    }
componentWillMount(){
//var id =  localStorage.getItem("ResId");
    var self = this.state;
           API.getMenu()
               .then((res) => {
                   console.log(res);
                   self.menuItems = res.menu;
                   this.setState(self);
               });
}
  render() {
      var ResMenu = [];
      var data = this.state.menuItems;
      data.map(function (temp, index) {
          ResMenu.push(
              <div className="row border-1-black margin-top-20">
                  <div className="col-md-8 div-res">
                      <div>

                          Name: {temp.name}
                      </div>
                      <div>
                          Price: {temp.price}
                      </div>
                      <div>
                          {temp.description}
                      </div>
                  </div>
                  <div className="col-md-4">
                      <button onClick={ () =>{this.addToCart(temp.id)}}  className="login-button" id="btnLogin" type="button">BUY</button>
                  </div>
              </div>
          );
      });
    return (
        <div className="row">
            <div className="col-md-2">
            </div>
            <div className="col-md-8">
                <div className="login-block">
                    <div className="padding-top-20">
                        {ResMenu}
                    </div>
                </div>
            </div>
            <div className="col-md-2">
            </div>
        </div>
    );
  }
}

export default withRouter(RestaurantMenu);


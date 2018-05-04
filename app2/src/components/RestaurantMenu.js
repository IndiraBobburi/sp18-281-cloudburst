import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import * as API from '../api/API.js';
class RestaurantMenu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            menuItems: [],
            cart:{}
        }
    }
componentWillMount(){
    var self = this.state;
           API.getMenu()
               .then((res) => {
                   console.log(res);
                   self.menuItems = res.menu;
                   var userId = localStorage.getItem("userId");
                   API.getCart(userId)
                       .then((res) => {
                           console.log(res);
                           self.cart = res;
                           this.setState(self);
                       });
                   //this.setState(self);
               });


}
    addToCart = (Itemid)=>{
        var cart = this.state.cart;
        var new_cart=null;
        var resId = parseInt(localStorage.getItem("ResId"));
        var self = this.state;
        if(resId == cart.restaurantId){
var items = cart.items;
var notFound = true;
for(var i = 0;i<items.length;i++){
    if(parseInt(items[i].id) === parseInt(Itemid)){
        items[i].quantity = items[i].quantity + 1;
        notFound = false;
    }
}if(notFound == true){
                cart.items.push({
                    "id": Itemid,
                    "quantity":1
                });
            }
        }

        else{
            cart = {
                "id": localStorage.getItem("userId"),
                "restaurantId": resId,
                "items": [
                    {
                        "id": Itemid,
                        "quantity":1
                    }
                    ]

            }
        }
       // localStorage.setItem("cart",cart);
        API.addToCart(cart)
            .then((res) => {
                console.log(res);
                self.cart = cart;
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
      },this);
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


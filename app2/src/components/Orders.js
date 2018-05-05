import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import * as API from '../api/API.js';

class Orders extends Component {
    constructor(props) {
        super(props);
        this.state = {
            orders:[]
        }
    }
    componentWillMount() {
        var self = this.state;
        var userId = localStorage.getItem("userId");
        API.getAllOrders(userId)
            .then((res) => {
                console.log(res);
                self.orders = res;
                this.setState(self);
            });
    }
  render() {
      var ordersList = [];
      var data = this.state.orders;
      var items = [];
      if(data && data.length!=0){
          data.map(function (temp, index) {
              items = temp.items;
              var itemList = [];
              items.map(function (tempItem, index) {
                  itemList.push(
                          <div className="col-md-8">
                              <div>
                                  {tempItem.Name}
                              </div>
                              <div>
                                  Qty: {tempItem.quantity}
                              </div>
                          </div>
                  );

              });

              ordersList.push(
                  <div className="row border-1-black margin-top-20">
                      <div className="col-md-8 div-res">
                          <div>
                              Restaurant: {temp.restaurantName}
                          </div>
                          <div>
                              Items:
                          </div>
                          <div>
                              {itemList}
                          </div>
                      </div>
                      <div className="col-md-4">

                      </div>
                  </div>
              );
          });
      }
    return (  
        <div className="row">
            <div className="col-md-2">
            </div>
            <div className="col-md-8">

                <div className="login-block">
                    <div className="logo">
                    <span>
                        ORDERS
                    </span>
                    </div>
            {ordersList}
                </div>
            </div>
            <div className="col-md-2">
            </div>
        </div>
    );
  }
}

export default withRouter(Orders);


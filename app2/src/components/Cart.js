import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import * as API from '../api/API.js';
class Cart extends Component {

    constructor(props) {
        super(props);
        this.state = {
            cart:{}
        }
    }
    componentWillMount() {
        var self = this.state;

        API.getCart('1234')
            .then((res) => {
                console.log(res);
                self.cart = res;
                this.setState(self);
            });
    }
    order=()=>{
        var self = this.state;
        API.getCart('1234')
            .then((res) => {
                console.log(res);
                self.cart = res;
                this.setState(self);
            });
    }

        render() {
            var cartList = [];
            var data = this.state.cart.items;
if(data && data.length!=0){
            data.map(function (temp, index) {
                cartList.push(
                    <div className="row border-1-black margin-top-20">
                        <div className="col-md-8 div-res">
                            <div>

                                Name: {temp.id}
                            </div>
                            <div>
                                Price: {temp.quantity}
                            </div>
                        </div>
                        <div className="col-md-4">

                        </div>
                    </div>
                );
            }, this);
        }
    return (
        <div className="row">
            <div className="col-md-2">
            </div>
            <div className="col-md-8">
                <div className="login-block">
                    <div className="row border-1-black margin-top-20">
                        <div className="col-md-8 div-res">
                            <div>
                                 {this.state.cart.id}
                            </div>

                        </div>
                        <div className="col-md-4">

                        </div>
                    </div>
                    <div className="padding-top-20">
                        {cartList}
                    </div>
                    <button onClick={() => {
                        this.order()
                    }} className="login-button" id="btnLogin" type="button">CHECKOUT
                    </button>
                </div>
            </div>
            <div className="col-md-2">
            </div>
        </div>
    );
  }
}

export default withRouter(Cart);


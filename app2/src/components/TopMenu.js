import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import Footer from './Footer.js';

class TopMenu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            userid:localStorage.getItem("userId")
        }
    }
    gotoHome = () =>{
        this.props.history.push("/");
    }

    gotoMenu = () =>{
        this.props.history.push("/menu");
    }

    gotoSignin = () =>{
        this.props.history.push("/signin");
    }

    gotoSignup = () =>{
        this.props.history.push("/signup");
    }
    gotoCart = () =>{
        this.props.history.push("/cart");
    }
    gotoOrders = ()=>{
        this.props.history.push("/orders");
    }
    logout = () =>{
        this.props.history.push("/");
        localStorage.clear();
    }
  render() {
        if(localStorage.getItem("userId") != null){
            return (
                <div className="row">
                    <div className="col-md-6">
                        <img src="./burger_logo.png"  className="logo"/>
                    </div>
                    <div className="col-md-6 menu-heading">
                        <div className="row">
                            <span>Welcome {this.state.userid}</span>
                        </div>
                        <div className="row">
                            <div className="col-md-2">
                                <div>
                                    <span onClick={ () =>{this.gotoHome()}}>HOME</span>
                                </div>
                            </div>
                            <div className="col-md-2">
                                <div>
                                    <span onClick={ () =>{this.gotoMenu()}}>MENU</span>
                                </div>
                            </div>
                            
                            <div className="col-md-2">
                                <div>
                                    <span onClick={ () =>{this.gotoCart()}}>CART</span>
                                </div>
                            </div>
                            <div className="col-md-2">
                                <div>
                                    <span onClick={ () =>{this.gotoOrders()}}>ORDERS</span>
                                </div>
                            </div>
                            <div className="col-md-2">
                                <div>
                                    <span onClick={ () =>{this.logout()}}>LOGOUT</span>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
            );
        }
        else{
            return (
                <div className="row">
                    <div className="col-md-6">
                        <img src="./burger_logo.png"  className="logo"/>
                    </div>
                    <div className="col-md-6 menu-heading">
                        <div className="row">
                            <div className="col-md-4">
                                <div>
                                    <span onClick={ () =>{this.gotoHome()}}>HOME</span>
                                </div>
                            </div>
                            <div className="col-md-4">
                                <div>
                                    <span onClick={ () =>{this.gotoMenu()}}>MENU</span>
                                </div>
                            </div>
                            <div className="col-md-4">
                                <div>
                                    <span onClick={ () =>{this.gotoSignin()}}>SIGN IN</span>
                                </div>
                            </div>
                            </div>
                    </div>
                </div>
            );

        }

  }
}

export default withRouter(TopMenu);


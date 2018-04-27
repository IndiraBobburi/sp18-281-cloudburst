import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import * as API from '../api/API.js';
class Menu extends Component {
    state={
        zipcode:''
    }
    getRestuarants = () =>{
        var data = this.state.zipcode;
        API.getRestaurants(data)
            .then((res) => {
                console.log(res);
               // this.props.history.push("/");
            });
    }
  render() {
    return (  
        <div className="row">
            <div className="col-md-2">
            </div>
            <div className="col-md-8">
            <div className="login-block">

                <span className="login-heading-style">Find Locations Near You</span>
                <form>
                    <div>
                        <label className="login-heading-style" for="zip">Enter Zip Code or City and State:</label>
                        <input className="login-textbox-style" classname="login-textbox-style" id="zip" maxlength="64" name="zip" type="text" value=""></input>
                    </div>
                    <button onClick={ () =>{this.getRestuarants()}}  className="login-button" id="btnLogin" type="submit">GO</button>
                </form>
                </div>
            </div>
            <div className="col-md-2">
            </div>
        </div>
    );
  }
}

export default withRouter(Menu);


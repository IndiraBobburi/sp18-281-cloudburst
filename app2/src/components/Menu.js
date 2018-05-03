import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';
import * as API from '../api/API.js';
class Menu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            pincode: '',
            listOfRes: []
        }
    }
    getRestuarants = () =>{
        var data = this.state.pincode;
        var self = this.state;
        API.getRestaurants(data)
            .then((res) => {
                console.log(res);
                self.listOfRes = res.restaurantlist;
                this.setState(self);
            });
    }
    getMenu = (data) =>{
        //localStorage.setItem("ResId",data)
        this.props.history.push("/Items");
    }
  render() {
      var listofRes = [];
      var data = this.state.listOfRes;
      data.map(function (temp, index) {
          listofRes.push(
             <div className="row border-1-black margin-top-20">
                 <div className="col-md-8 div-res">
                 <div>
                     Name: {temp.name}
                 </div>
                     <div>
                         Address: {temp.address}
                     </div>
                     <div>
                         Phone: {temp.phone}
                     </div>
                 </div>
                 <div className="col-md-4">
                     <button onClick={ () =>{this.getMenu(temp.id)}}  className="login-button" id="btnLogin" type="button">SELECT</button>
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

                <span className="login-heading-style">Find Restaurants Near You</span>
                <form>
                    <div>
                        <label className="login-heading-style" for="zip">Enter Zip Code or City and State:</label>
                        <input className="login-textbox-style"  id="zip" maxLength="64"
                               name="zip" type="text"
                               value={this.state.pincode}
                               onChange={(event) => {
                                   this.setState({
                                       pincode: event.target.value
                                   });
                               }}></input>
                    </div>
                    <button onClick={ () =>{this.getRestuarants()}}  className="login-button" id="btnLogin" type="button">GO</button>
                </form>
                <div className="padding-top-20">
                    {listofRes}
                </div>
                </div>
            </div>
            <div className="col-md-2">
            </div>
        </div>
    );
  }
}

export default withRouter(Menu);


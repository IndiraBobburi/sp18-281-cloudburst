import { Route, withRouter,BrowserRouter } from 'react-router-dom';
import '../App.css';
import React, { Component } from 'react';

class Footer extends Component {
  render() {
    return (  
        <div className="row footer footer-style">
         <div className="col-md-6">
             <a target="_blank" href="https://www.olo.com/"> Online Ordering by <strong>Olo.com</strong></a>
   </div>
   <div className="col-md-6">
       <span>Participating Locations | User Agreement | Privacy Policy | Feedback | Contact Us</span>
     </div>
</div>
    );
  }
}

export default withRouter(Footer);


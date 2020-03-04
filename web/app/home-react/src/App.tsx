import React from 'react';
import './App.css';

import {withRouter} from 'react-router-dom' ;


import {Home} from "./components/home/Home";
import {Routers} from "./routers/Routers";


function App() {
    return (
        <div className="Home full-height">
            <Home></Home>
            <Routers></Routers>
        </div>
    );
}

export default withRouter(App);

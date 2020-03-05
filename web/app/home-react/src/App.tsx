import React from 'react';
import './App.css';

import {Route, HashRouter as Router, Switch} from 'react-router-dom' ;


import {Home} from "./components/home/Home";

import {Expense} from "./components/book/Expense";
import {NotFound} from "./components/exception/NotFound";

const routeList = [
    Expense,
];


function App() {
    return (
        <Router>
            <Switch>
                <Route path='/abc' component={Home}></Route>
                <Route path={"/expense"} component={Expense}></Route>
                <Route path="/about" render={() => <div title='About'>23232323</div>}/>
                <Route component={NotFound}/>
            </Switch>
        </Router>
    );
}

export default App;

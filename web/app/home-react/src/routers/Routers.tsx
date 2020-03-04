import * as React from "react";
import {
    Switch,
    Route,
} from "react-router-dom";
import {Expense} from "../components/book/Expense";

export interface RoutersProps {
}

class state {
}

export class Routers extends React.Component<RoutersProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {};

    render() {
        return (
            <Switch>
                <Route path="/expense">
                    <Expense/>
                </Route>
            </Switch>
        );
    }
}
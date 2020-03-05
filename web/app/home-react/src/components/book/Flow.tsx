import React from "react";
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link, useRouteMatch, useParams
} from "react-router-dom";

import {Button, Flex} from 'antd-mobile';


export interface FlowProps {
}

class state {
}

function Top() {
    return (
        <div style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
        }}>
            <Flex>
                <Flex.Item style={{
                    marginTop: "10px",
                }}>
                    <Link to="/expense"> 记一笔 </Link>
                </Flex.Item>
            </Flex>
        </div>
    )
}


export class Flow extends React.Component<FlowProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {};

    render() {
        return (
            <div>
                <Top></Top>
            </div>
        );
    }
}
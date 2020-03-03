import * as React from "react";
import {Tabs, WhiteSpace} from 'antd-mobile';

import {Flow} from "./Flow"
import {Statement} from "./Statement"

const tabs = [
    {title: "记账"},
    {title: "报表"},
];

export interface BodyProps {
}

class state {
    public selectedTab!: string;
}

export class Body extends React.Component<BodyProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {
        selectedTab: 'redTab',
    };

    render() {
        return (
            // todo 布局
            <Tabs tabs={tabs}
                  initialPage={0}
            >
                <Flow/>
                <Statement/>
            </Tabs>
        );
    }
}
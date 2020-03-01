import * as React from "react";

import {Flow} from "./Flow"
import {Statement} from "./Statement"

export interface BodyProps {
}

class state {
    public selectedTab!: string;
}

export class HBody extends React.Component<BodyProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {
        selectedTab: 'redTab',
    };

    render() {
        return (
            // todo 布局
            <div>
                <Flow/>
                <Statement/>
            </div>
        );
    }
}
import * as React from "react";

export interface FlowProps {
}

class state {
    public selectedTab!: string;
}

export class Flow extends React.Component<FlowProps, {}> {
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
                Body流水
            </div>
        );
    }
}
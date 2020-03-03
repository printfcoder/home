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
            <div className={"full-height box-full"} style={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                backgroundColor: '#fff'
            }}>
               流水记账
            </div>
        );
    }
}
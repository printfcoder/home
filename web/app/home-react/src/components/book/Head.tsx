import * as React from "react";

export interface HeadProps {
}

class state {
    public selectedTab!: string;
}

export class Head extends React.Component<HeadProps, {}> {
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
                头部
            </div>
        );
    }
}
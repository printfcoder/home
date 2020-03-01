import * as React from "react";

export interface StatementProps {
}

class state {
    public selectedTab!: string;
}

export class Statement extends React.Component<StatementProps, {}> {
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
                报表清单
            </div>
        );
    }
}
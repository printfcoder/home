import * as React from "react";

export interface DrawerProps {
}

class state {
    public selectedTab!: string;
}

export class Drawer extends React.Component<DrawerProps, {}> {
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
                左边抽屉
            </div>
        );
    }
}
import {Drawer} from './Drawer';
import {Body} from './HBody';

import * as React from "react";

export interface BookProps {
}

class state {
    selectedMe!: boolean;
}

export class Expense extends React.Component<BookProps, {}> {
    constructor(props: any) {
        super(props);
    }

    name = "edit"
    state: state = {
        selectedMe: true,
    };

    render() {
        return (
            // todo 布局
            <div
                className={"full-height body-" + (this.state.selectedMe ? "normal" : "hidden")}
                style={{}}>
                <Body/>
                <Drawer/>
            </div>
        );
    }
}
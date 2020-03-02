import {Drawer} from './Drawer';
import {HBody} from './HBody';
import {Head} from './Head';

import * as React from "react";
import store from "../../store"

export interface BookProps {
}

class state {
}

export class Book extends React.Component<BookProps, {}> {
    constructor(props: any) {
        super(props);
    }

    name = "book"
    state: state = {};

    render() {
        return (
            // todo 布局
            <div
                className={"body-" + (this.name === store.getState().home.bottomTabBar.selectTab ? "normal" : "hidden")}
                style={{}}>
                <Head/>
                <HBody/>
                <Drawer/>
            </div>
        );
    }
}
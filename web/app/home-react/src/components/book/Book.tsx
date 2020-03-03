import {Drawer} from './Drawer';
import {HBody} from './HBody';
import {Head} from './Head';

import * as React from "react";
import store from "../../store"

export interface BookProps {
}

class state {
    selectedMe!: boolean;
}

export class Book extends React.Component<BookProps, {}> {
    constructor(props: any) {
        super(props);
    }

    name = "zhangben"
    state: state = {
        selectedMe: false,
    };

    componentDidMount(): void {
        store.subscribe(() => {
            let selectedMe = false;
            if (store.getState().home.bottomTabBar.selectTab === this.name) {
                selectedMe = true
            }
            this.setState({
                ...this.state,
                selectedMe: selectedMe
            })
        })
    }

    render() {
        return (
            // todo 布局
            <div
                className={"body-" + (this.state.selectedMe ? "normal" : "hidden")}
                style={{}}>
                <Head/>
                <HBody/>
                <Drawer/>
            </div>
        );
    }
}
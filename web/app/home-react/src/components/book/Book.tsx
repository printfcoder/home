import {Drawer} from './Drawer';
import {HBody} from './HBody';
import {Head} from './Head';

import * as React from "react";

export interface BookProps {
}

class state {
    public display!: string;
}

export class Book extends React.Component<BookProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {
        display: "block",
    };

    render() {
        return (
            // todo 布局
            <div style={{display: this.state.display}}>
                <Head/>
                <HBody/>
                <Drawer/>
            </div>
        );
    }
}
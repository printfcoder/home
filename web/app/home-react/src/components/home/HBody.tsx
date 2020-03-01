import {Book} from '../../components/book/Book';

import * as React from "react";

export interface HBodyProps {
}

class state {
    public display!: string;
}

export class HBody extends React.Component<HBodyProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {
        display: "none",
    };

    render() {
        return (
            // todo 布局
            <div>
                <Book/>
            </div>
        );
    }
}
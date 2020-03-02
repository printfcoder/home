import {Book} from '../../components/book/Book';

import * as React from "react";

export interface HBodyProps {
}

class state {
}

export class HBody extends React.Component<HBodyProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {};

    render() {
        return (
            // todo 布局
            <div>
                <Book/>
            </div>
        );
    }
}
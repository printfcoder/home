import {BottomTabBar} from './BottomTabBar';
import {HBody} from './HBody';

import * as React from "react";

export interface HomeProps {
}

class state {
    public selectedTab!: string;
}

export class Home extends React.Component<HomeProps, {}> {
    constructor(props: any) {
        super(props);
    }

    state: state = {
        selectedTab: 'redTab',
    };

    render() {
        return (
            <div className={"full-height"}>
                <HBody/>
                <BottomTabBar/>
            </div>
        );
    }
}
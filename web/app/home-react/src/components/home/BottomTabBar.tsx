import {TabBar} from 'antd-mobile';

import HIcon from "../../components/icon/HIcon";


import * as React from "react";

export interface BottomTabBarProps {
}

class state {
    public selectedTab!: string;
}


export class BottomTabBar extends React.Component<BottomTabBarProps, {}> {
    constructor(props: any) {
        super(props);

        this.state = {
            selectedTab: "home",
        }
    }

    state: state;

    render() {
        return (
            <div style={{
                position: 'fixed',
                height: '100%',
                width: '100%',
                top: 0
            }}>
                <TabBar
                    unselectedTintColor="#949494"
                    tintColor="#33A3F4"
                    barTintColor="white"
                >
                    <TabBar.Item
                        title="首页"
                        key="home"
                        icon={<HIcon name="home"></HIcon>}
                        selectedIcon={<HIcon name="home-selected"></HIcon>}
                        selected={this.state.selectedTab === 'home'}
                        badge={1}
                        onPress={() => {
                            this.setState({
                                selectedTab: 'home',
                            });
                        }}
                        data-seed="home"
                    >
                    </TabBar.Item>
                    <TabBar.Item
                        title="账本"
                        key="zhangben"
                        icon={<HIcon name="zhangben"></HIcon>}
                        selectedIcon={<HIcon name="zhangben-selected"></HIcon>}
                        selected={this.state.selectedTab === 'zhangben'}
                        badge={1}
                        onPress={() => {
                            this.setState({
                                selectedTab: 'zhangben',
                            });
                        }}
                        data-seed="zhangben"
                    >
                    </TabBar.Item>
                    <TabBar.Item
                        title="计划"
                        key="plan"
                        icon={<HIcon name="plan"></HIcon>}
                        selectedIcon={<HIcon name="plan-selected"></HIcon>}
                        selected={this.state.selectedTab === 'plan'}
                        badge={1}
                        onPress={() => {
                            this.setState({
                                selectedTab: 'plan',
                            });
                        }}
                        data-seed="plan"
                    >
                    </TabBar.Item>
                    <TabBar.Item
                        title="我的"
                        key="me"
                        icon={<HIcon name="me"></HIcon>}
                        selectedIcon={<HIcon name="me-selected"></HIcon>}
                        selected={this.state.selectedTab === 'yellowTab'}
                        onPress={() => {
                            this.setState({
                                selectedTab: 'yellowTab',
                            });
                        }}
                    >
                    </TabBar.Item>
                </TabBar>
            </div>
        );
    }
}
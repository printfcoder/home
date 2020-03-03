import {TabBar} from 'antd-mobile';

import HIcon from "../../components/icon/HIcon";

import store from "../../store"
import {updateCart} from "../../reducers/home/Action"


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

    emitTabChangeEvent(tabName: string) {
        this.setState({selectedTab: tabName})
        store.dispatch(updateCart(tabName))
    }

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
                            this.emitTabChangeEvent("home")
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
                            this.emitTabChangeEvent("zhangben")
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
                            this.emitTabChangeEvent("plan")
                        }}
                        data-seed="plan"
                    >
                    </TabBar.Item>
                    <TabBar.Item
                        title="我的"
                        key="me"
                        icon={<HIcon name="me"></HIcon>}
                        selectedIcon={<HIcon name="me-selected"></HIcon>}
                        selected={this.state.selectedTab === 'me'}
                        onPress={() => {
                            this.emitTabChangeEvent("me")
                        }}
                    >
                    </TabBar.Item>
                </TabBar>
            </div>
        );
    }
}
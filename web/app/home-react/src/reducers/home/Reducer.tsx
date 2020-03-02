import {CHANGE_BOTTOM_TAB} from "./Action";
import {AnyAction} from "redux";

const homeState = {
    bottomTabBar: {
        selectTab: 'home',
    }
}

export default function homeReducer(state = homeState, action: AnyAction) {
    switch (action.type) {
        case CHANGE_BOTTOM_TAB:
            return {
                ...state,
                bottomTabBar: {
                    selectTab: action.payload.tabName,
                }
            }
        default:
            return state;
    }
    return state;
}
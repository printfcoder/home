import {createStore} from "redux";
import {composeWithDevTools} from 'redux-devtools-extension';

import rootReducer from './reducers/Reducers';

let store = createStore(rootReducer, composeWithDevTools());

export default store;
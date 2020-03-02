import {combineReducers} from 'redux';

import homeReducer from "./home/Reducer";
import bookReducer from "./book/Reducer";

const allReducers = {
    book: bookReducer,
    home: homeReducer,
}

const RootReducer = combineReducers(allReducers);

export default RootReducer
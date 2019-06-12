import { combineReducers } from 'redux'
import init from './initial';

/**
 * helper function turns an object whose values are different reducing functions
 * into a single reducing function you can pass to createStore.
 */
export default combineReducers({
  init
})

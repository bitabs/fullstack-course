import { createStore, applyMiddleware, compose } from 'redux'
import thunk from 'redux-thunk'
import rootReducer from './reducers'

const
  // the initial state of redux store
  initialState  = {},
  // Store enhancers are a formal mechanism for adding capabilities to Redux
  // itself.
  enhancers     = [],
  // provides a way to interact with actions that have been dispatched to the
  // store before they reach the store's reducer.
  middleware    = [thunk];

/**
 * If in development environment, enable redux devtool extension
 */
if (process.env.NODE_ENV === 'development') {
  const __REDUX_DEVTOOLS_EXTENSION__ = window.__REDUX_DEVTOOLS_EXTENSION__
  if (typeof __REDUX_DEVTOOLS_EXTENSION__ === 'function')
    enhancers.push(__REDUX_DEVTOOLS_EXTENSION__())
}

/**
 * functional programming utility, and is included in Redux as a convenience.
 * You might want to use it to apply several store enhancers in a row.
 */
const composedEnhancers = compose(
  applyMiddleware(...middleware),
  ...enhancers
)

/**
 * Creates a Redux store that holds the complete state tree of your app.
 * There should only be a single store in your app.
 * @type {Store<any, AnyAction> & Store<S & {}, A> & {dispatch: any}}
 */
const store = createStore(
  rootReducer,
  initialState,
  composedEnhancers
);

export default store;

import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'
import App from './components/app'
import store from './store'
import './index.css'

render(
  // we pass the redux store to the entire app
  <Provider store={store}>
    <App/>
  </Provider>,
  document.getElementById('root')
);

// if in development environment, allow hot reload for debugging purposes
if (process.env.NODE_ENV === 'development') {
  module.hot.accept();
}

import React from 'react'
import { render } from 'react-dom'
import ApolloClient from "apollo-boost"
import { ApolloProvider } from "react-apollo"
import App from './components/app'
import './index.css'

const client = new ApolloClient()

const Application = () => (
  <ApolloProvider client={client}>
    <App />
  </ApolloProvider>
)

render(<Application />, document.getElementById('root'))

// render(
//   // we pass the redux store to the entire app
//   <Provider store={store}>
//     <App/>
//   </Provider>,
//   document.getElementById('root')
// );

// if in development environment, allow hot reload for debugging purposes
if (process.env.NODE_ENV === 'development') {
  module.hot.accept();
}

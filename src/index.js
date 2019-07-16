import React from 'react'
import { render } from 'react-dom'
import App from './components/app'
import './index.css'
import {ApolloProvider} from "react-apollo";
import ApolloClient from "apollo-boost";

const client = new ApolloClient({
  uri: 'http://localhost:3000/graphql'
});


const Application = () => (
  <ApolloProvider client={client}>
    <App />
  </ApolloProvider>
)

render(<Application />, document.getElementById('root'))

// if in development environment, allow hot reload for debugging purposes
if (process.env.NODE_ENV === 'development') {
  module.hot.accept();
}

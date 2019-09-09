import React from 'react'
import { BrowserRouter } from "react-router-dom"
import { render } from 'react-dom'
import App from './components/app'
import './index.css'
import {ApolloProvider} from "react-apollo";
import ApolloClient, { gql } from "apollo-boost";

const client = process.env.NODE_ENV === 'development'
  ? new ApolloClient({ uri: "http://localhost:3000" })
  : new ApolloClient()


const Application = () => (
  <BrowserRouter>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </BrowserRouter>
)

render(<Application />, document.getElementById('root'))

// if in development environment, allow hot reload for debugging purposes
if (process.env.NODE_ENV === 'development') {
  module.hot.accept();
}

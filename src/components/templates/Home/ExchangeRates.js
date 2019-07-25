import React from 'react'
import { Query } from "react-apollo"
import { GET_TUTORIAL } from "gql/queries"

const ExchangeRates = () => (
  <Query query={GET_TUTORIAL}>
    {({ loading, error, data }) => {
      console.log(data)
      if (loading) return <p>Loading...</p>;
      if (error) return <p>Error :(</p>;
      // console.log(data)
      return null
    }}
  </Query>
)

export default ExchangeRates

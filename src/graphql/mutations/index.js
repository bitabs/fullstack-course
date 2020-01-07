import { gql } from "apollo-boost"

export const REGISTER_USER = gql`
  mutation Register($firstName: String!, $lastName: String!, $emailAddress: String!, $phoneNumber: String!, $username: String!, $password: String!) {
    registerUser(firstName: $firstName, lastName: $lastName, emailAddress: $emailAddress, phoneNumber: $phoneNumber, username: $username, password: $password) {
      firstName
      lastName
      emailAddress
      phoneNumber
      username
      password
    }
  }
`;

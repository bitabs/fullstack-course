import { gql } from "apollo-boost"

export const REGISTER_USER = gql`
  mutation Register (
    $first_name    : String = "Naseebullah", 
    $last_name     : String = "Ahmadi", 
    $email_address : String = "javastastic@gmail.com", 
    $phone_number  : String = "07473693312", 
    $username      : String = "naseebullahsafi", 
    $password      : String = "afghan500"
  ) {
    registerUser(firstName: $first_name, lastName: $last_name, emailAddress: $email_address, phoneNumber: $phone_number, username: $username, password: $password) {
      firstName
      lastName
      emailAddress
      phoneNumber
      username
      password
    }
  }
`;

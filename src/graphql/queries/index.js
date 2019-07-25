import { gql } from "apollo-boost"


export const GET_TUTORIAL = gql`  
    {
        tutorial(id: 1) {
            title
        }
    }
`;

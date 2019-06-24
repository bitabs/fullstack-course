import React from 'react'
import { Wrapper } from './Home.style'
import H1 from "components/atoms/H1";

const Home = ({children, ...props}) => (
  <Wrapper>
    <H1>Hello, World!</H1>
  </Wrapper>
);

export default Home;

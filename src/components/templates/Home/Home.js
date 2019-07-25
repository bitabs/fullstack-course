import React from 'react'
import { Wrapper } from './Home.style'
import H1 from "components/atoms/H1";
import ExchangeRates from "./ExchangeRates";

const Home = ({children, ...props}) => (
  <Wrapper>
    <H1>Hello, World!</H1>
    <ExchangeRates/>
  </Wrapper>
);

export default Home;

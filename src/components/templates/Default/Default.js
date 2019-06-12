import React from 'react'
import PropTypes from 'prop-types';
import { Wrapper } from './Default.style'

const Default = ({header, children, footer, ...props}) => (
  <Wrapper>
    {children}
  </Wrapper>
);

Default.protoTypes = {
  header    : PropTypes.node.isRequired,
  children  : PropTypes.any.isRequired,
  footer    : PropTypes.node.isRequired
}

export default Default;
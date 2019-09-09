import React from 'react'
import { FormGroup, FormLabel, FormInput } from './Input.style'

const Input = props => (
  <FormGroup>
    <FormLabel
      htmlFor={props.name}
    >{props.title}</FormLabel>
    <FormInput
      id            = { props.name }
      name          = { props.name }
      type          = { props.type }
      value         = { props.value }
      onChange      = { props.handleChange }
      placeholder   = { props.placeholder }
    />
  </FormGroup>
)

export default Input

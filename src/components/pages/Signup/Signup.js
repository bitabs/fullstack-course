import React, { PureComponent } from 'react'
import Input from "components/molecules/Input"
import { REGISTER_USER } from "../../../graphql/mutations";
import Mutation from "react-apollo/Mutation";

class Signup extends PureComponent {
  constructor(props) {
    super(props)
    this.state = {
      form: {
        firstName: {
          value : "",
          title: "First Name",
          name: "first-name",
          placeholder: "Enter your first name",
          handleChange: (p) => this.handleUserInput(p, 'firstName'),
          type: "input"
        },
        lastName: {
          value : "",
          title: "Last Name",
          name: "last-name",
          placeholder: "Enter your last name",
          handleChange: (p) => this.handleUserInput(p, 'lastName'),
          type: "input"
        },
        emailAddress: {
          value: "",
          title: "Email address",
          name: "email-address",
          placeholder: "Enter your email address",
          handleChange: (p) => this.handleUserInput(p, 'emailAddress'),
          type: "input"
        },
        phoneNumber: {
          value : "",
          title: "Phone number",
          name: "phone-number",
          placeholder: "Enter your phone number",
          handleChange: (p) => this.handleUserInput(p, 'phoneNumber'),
          type: "input"
        },
        userName: {
          value    : "",
          title: "Username",
          name: "username",
          placeholder: "Enter your username",
          handleChange: (p) => this.handleUserInput(p, 'userName'),
          type: "input"
        },
        password: {
          value    : "",
          title: "Password",
          name: "password",
          placeholder: "Enter your password",
          handleChange: (p) => this.handleUserInput(p, 'password'),
          type: "input"
        }
      }
    }
  }

  handleUserInput = (props, field) => {
    props.preventDefault()
    const { value } = props.target
    this.setState(prevState => ({
      form: {
        ...prevState.form,
        [field]: {
          ...prevState.form[field],
          value
        }
      }
    }))
  }

  handleSubmit = (e) => {
    e.preventDefault()

  }


  render() {
    const { handleSubmit } = this

    return (
      <Mutation mutation={REGISTER_USER}>
        {(registerUser, { data }) => {
          if (data)
            console.log(data);

          return (
            <div className="signup">
              <h1>Signup Form</h1>
              <form onSubmit={e => {
                e.preventDefault()
                registerUser({variables: {
                    firstName: "Haseebullah",
                    lastName: "Ahmadi",
                    emailAddress: "ezatsafi@hotmail.com",
                    phoneNumber: "024332423",
                    username: "jksdjhfds",
                    password: "dfsfsdf"
                  }})
              }}>
                {Object.values(this.state.form).map((field, i) => {
                  const { type } = field

                  if (type === "input") {
                    return <Input
                      key={i}
                      id={field.name}
                      name={field.name}
                      title={field.title}
                      handleChange={field.handleChange}
                    />
                  }
                })}
                <input type="submit" value="Submit" />
              </form>
            </div>
          )
        }}
      </Mutation>
    );
  }
}

export default Signup

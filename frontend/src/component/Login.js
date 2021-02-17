import React, { useState }from 'react';
import PropTypes from 'prop-types';

import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import NavbarComponent from './Navbar'

import '../App.css'
import Container from 'react-bootstrap/Container'

async function loginUser(credentials) {
    return fetch('http://localhost:8888/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(credentials)
    }).then(response => response.json())
    .then((responseData) => {
      return responseData;
    })
 }

export default function Login({ setToken }) {
  const [username, setUserName] = useState();
  const [password, setPassword] = useState();

  const handleSubmit = async e => {
    e.preventDefault();
    const token = await loginUser({
      username,
      password
    });
    setToken(token);
    window.location.reload(false);
    console.log(token)
  }

  return(
    <div>
      <NavbarComponent></NavbarComponent>
      <Container>
        <div className="login-padding">
          <Form onSubmit={handleSubmit}>
            <Form.Group controlId="Username">
              <Form.Label>Username</Form.Label>
              <Form.Control type="text" placeholder="Username" onChange={e => setUserName(e.target.value)} />
            </Form.Group>

            <Form.Group controlId="Password">
              <Form.Label>Password</Form.Label>
              <Form.Control type="password" placeholder="Password" onChange={e => setPassword(e.target.value)}/>
            </Form.Group>
            <Button variant="primary" type="submit">
              Login
            </Button>
          </Form>
        </div>
      </Container>
    </div>
  )
}

Login.propTypes = {
  setToken: PropTypes.func.isRequired,
}
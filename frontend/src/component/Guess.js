import React,{useState} from 'react';
import NavbarComponent from './Navbar'
import Button from 'react-bootstrap/button'
import Alert from 'react-bootstrap/alert'
import Form from 'react-bootstrap/Form'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

async function guessSend(number) {
  return fetch('http://localhost:8080/guess?number='+number, {
    method: 'GET',
    headers: {
      'Token' : localStorage.getItem('token'),
    },
  }).then(response => response.json())
  .then((responseData) => {
    return responseData;
  })
}


export default function Guess() {
  const [number, setNumber] = useState();
  const [message, setMessage] = useState();

  const handleGuess = async e => {
    e.preventDefault();
    const respMessage = await guessSend(number);
    setMessage(respMessage.message);
  }

  return(
    <div>
        <NavbarComponent></NavbarComponent>
        <div className="login-padding"></div>
        <center>
            <h2>Guessing Game Page</h2>
            <br/>
            <Alert variant='primary'>
            We have selected a random number between 1 and 100. See if you We'll tell you if your guess was too high or too low.
            </Alert>

            <br/>

            <Form onSubmit={handleGuess}>
              <Row className="justify-content-md-center">
                <Col md="auto">
                  <Form.Control size="lg" type="number" placeholder="Guess Number" onChange={e => setNumber(e.target.value)}/>
                </Col>
              </Row>
              <br />
              {message}
              <br />

              <Button variant="primary" type="submit">
              Guess Now
            </Button>
            </Form>

           
          
        </center>
    </div>
    
  );
}
import React,{useState} from 'react';
import NavbarComponent from './Navbar'
import Button from 'react-bootstrap/Button'
import Alert from 'react-bootstrap/Alert'
import Form from 'react-bootstrap/Form'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

async function guessSend(number) {
  return fetch('http://localhost:8889/guess?number='+number, {
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
  const [buttonTxt = "Guess Now!", setButton] = useState();
  const [alertType = "warning", setAlert] = useState();
  const [hidden = "d-none", setHidden] = useState();
  const [hint, setHint] = useState();

  const handleGuess = async e => {
    if(buttonTxt == "New Game!"){
      setButton("Guess Now!")
      setNumber(0)
      setHidden("d-none")
    }

    e.preventDefault();
    const respMessage = await guessSend(number);
    if(respMessage.status){
      setButton("New Game!")
      setAlert("success")
      setHint("")
    }else{
      setAlert("warning")
      setHint(" | Range "+respMessage.first+" to "+respMessage.last)
    }
      //setAlert("warning")
      setHidden("")
      setNumber(0)
      setMessage(respMessage.number+" "+respMessage.message);
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
                  <Form.Control size="lg" autoFocus type="number" placeholder="Guess Number" onChange={e => setNumber(e.target.value)}/> 
                </Col>
                <Col md='auto'>
                  <Button variant="primary" type="submit">
                    {buttonTxt}
                  </Button>
                </Col>
              </Row>
              <br/>
              <Row className="justify-content-md-center">
                <Col md='auto'>
                  <Alert className={hidden} variant={alertType}>{message}{hint}</Alert>
                </Col>
              </Row>
            </Form>

           
          
        </center>
    </div>
    
  );
}
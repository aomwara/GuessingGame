import React from 'react';
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'
import useToken from './useToken';
import logout from './logout'

export default function NavbarComponent() {
    const { token } = useToken();
    var text = ""
   
    if(token != ""){
        text = "logout"
    }else{
        text = "login"
    }

  return(
      <div>
        <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
            <Navbar.Brand href="#home">Guessing Game!</Navbar.Brand>
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">
                <Nav className="mr-auto"></Nav>
                <Nav>
                <Nav.Link onClick={logout}>{text}</Nav.Link>
                </Nav>
            </Navbar.Collapse>
        </Navbar>
      </div>
  );
}
import React, { Component } from "react"
import "./Register.css"
import { Button } from "react-bootstrap"
import { Form } from "react-bootstrap"
import firebase from "firebase"
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth"

firebase.initializeApp({
    apiKey: "AIzaSyC_u27mI5neEM6BVCM1yqrAN-WHTk8-QKE",
    authDomain: "course-online-b26e1.firebaseapp.com"
})

class Register extends Component {
    state = { isSigndIn: false}
    uiConfig = {
        signInFlow: "popup",
        signInOptions: [
            firebase.auth.FacebookAuthProvider.PROVIDER_ID
        ],
        callback: {
            signInSuccess: () => false
        }
    }

    componentDidMount = () => {

        firebase.auth().onAuthStateChanged(user =>{
            this.setState({isSigndIn:!!user})
        })
    }

    render(){
    return (
        <div>
            <Form>
                <h1>PERSONAL INFORMATION</h1>
                <br></br>
                <Form.Group controlId="formfName">
                    <Form.Label>Frist name</Form.Label>
                    <Form.Control type="fname" placeholder="Enter frist name" />
                </Form.Group>

                <Form.Group controlId="formlName">
                    <Form.Label>Lrist name</Form.Label>
                    <Form.Control type="lname" placeholder="Enter lrist name" />
                </Form.Group>

                <Form.Group controlId="formIdcard">
                    <Form.Label>ID card</Form.Label>
                    <Form.Control type="idcard" placeholder="Enter id card" />
                </Form.Group>

                {this.state.isSigndIn ? (
                    <span>
                        <div>Signed In!</div>
                        <h1>Welcome {firebase.auth().currentUser.displayName}</h1>
                        <img alt="profile picture" src={firebase.auth().currentUser.photoURL}/>
                    </span>
                ) : (
                    <StyledFirebaseAuth
                        uiConfig={this.uiConfig}
                        firebaseAuth={firebase.auth()}
                    />
                )}

                <Form.Group controlId="formBasicChecbox">
                    <Form.Check type="checkbox" label="Accept" />
                </Form.Group>

                <Button variant="primary" type="Confirm">
                    Confirm
                </Button>
            </Form>
    </div>
    )
  }
}
  
export default Register;
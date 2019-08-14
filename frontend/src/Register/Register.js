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
            firebase.auth.FacebookAuthProvider.PROVIDER_ID,
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
                        <h1>Signed In!</h1>
                        <h1>FaceBook: {firebase.auth().currentUser.displayName}</h1>
                        <h1>Email: {firebase.auth().currentUser.email}</h1>
                        <br></br>
                        <img class="profile" alt="profile picture" src={firebase.auth().currentUser.photoURL + "?height=500"}/>
                    </span>
                ) : (
                    <StyledFirebaseAuth
                        uiConfig={this.uiConfig}
                        firebaseAuth={firebase.auth()}
                    />
                )}

                <Form.Group controlId="formPhone">
                    <Form.Label>Phone</Form.Label>
                    <Form.Control type="phone" placeholder="Enter phone" />
                </Form.Group>

                <Form.Group controlId="formAddress">
                    <Form.Label>Address</Form.Label>
                    <Form.Control type="address" placeholder="Enter address" />
                </Form.Group>

                {/* <Form.Group controlId="formBasicChecbox">
                    <Form.Check type="checkbox" label="Accept" />
                </Form.Group> */}
                <center>
                <Button class="butt" variant="primary" type="Confirm">
                    Confirm
                </Button>
                </center>
                
            </Form>
    </div>
    )
  }
}
  
export default Register;
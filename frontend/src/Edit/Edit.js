import React, { Component } from "react"
import "./Edit.css"
import axios from 'axios'
import { Form } from "react-bootstrap"
import { Button } from "react-bootstrap"
class Edit extends Component {
    constructor(props) {
        super(props)

        this.handleDelete = this.handleDelete.bind(this);
    }

    handleDelete() {
        const { match: { params } } = this.props;
        axios.delete(`http://localhost:8081/employee/${params.member_id}`)
            .then(() => {
                console.log('member deleted');
            });
    }
    render() {
        return (
            <div>
                <Form>
                    <h1>Edit Member</h1>
                            <Form.Group controlId="formTHName">
                                <Form.Label>Thai name</Form.Label>
                                <Form.Control type="thname" placeholder="Enter Thai name" />
                            </Form.Group>

                            <Form.Group controlId="formENName">
                                <Form.Label>English name</Form.Label>
                                <Form.Control type="engname" placeholder="Enter English name" />
                            </Form.Group>

                            <Form.Group controlId="formFBName">
                                <Form.Label>Facebook name</Form.Label>
                                <Form.Control type="fbname" placeholder="Enter Facebook name" />
                            </Form.Group>

                            <Form.Group controlId="formEmail">
                                <Form.Label>Email</Form.Label>
                                <Form.Control type="email" placeholder="Enter Email" />
                            </Form.Group>

                            <Form.Group controlId="formIdcard">
                                <Form.Label>IDcard</Form.Label>
                                <Form.Control type="Idcard" placeholder="Enter IDcard" />
                            </Form.Group>

                            <Form.Group controlId="formPhone">
                                <Form.Label>Phone</Form.Label>
                                <Form.Control type="phone" placeholder="Enter Phone" />
                            </Form.Group>

                            <Form.Group controlId="formAddress">
                                <Form.Label>English name</Form.Label>
                                <Form.Control type="address" placeholder="Enter Address" />
                            </Form.Group>

                            <center>
                            <Button variant="success" type="Update" className="Button">
                                UPDATE
                            </Button>
                            <br></br>
                            <br></br>
                            <Button variant="danger" type="Delete" className="Button" onClick={this.handleDelete}>
                                DELETE
                            </Button>
                            </center>
                </Form>
            </div>
        )
    }

}

export default Edit;

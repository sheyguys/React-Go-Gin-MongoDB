import React, { Component } from "react"
import { Form } from "react-bootstrap"
import axios from 'axios';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { Button } from "react-bootstrap"
import "./Result.css"

class Result extends Component {
    constructor(prop) {
        super(prop)

        this.state = { row:[] }
    }

    componentDidMount() {
        var Allmember = []
        axios.get('http://localhost:8081/employee').then(response => {
          console.log(response.data);
          response.data.member.forEach(member => {
            Allmember.push(member)
          })
          this.setState({ row : Allmember })
        })
        .catch(error => {
          console.log(error);
        });
      }
    
    handleDelete(id){
        axios.delete('http://localhost:8081/employee/'+id)
        console.log(id);
        alert('Delete Complete!!');
        window.location.reload();
    }

    render(){
        return (
            <div>
                <Form>
                    <h1>EMPLOYEE RECORDS</h1>                
                    <br></br>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Eng Name</TableCell>
                                <TableCell>Thai Name</TableCell>
                                <TableCell>ID.card</TableCell>
                                <TableCell>Phone</TableCell>
                                <TableCell>Address</TableCell>
                                <TableCell>Email</TableCell>
                                <TableCell>Facebook</TableCell>
                                <TableCell>DELETE</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.row.map(member => (
                            <TableRow key={member.member_id}>
                                <TableCell>{member.member_name_eng}</TableCell>
                                <TableCell>{member.member_name_th}</TableCell>
                                <TableCell>{member.member_idcard}</TableCell>
                                <TableCell>{member.member_phone}</TableCell>
                                <TableCell>{member.member_address}</TableCell>
                                <TableCell>{member.member_email}</TableCell>
                                <TableCell>{member.member_facebook}</TableCell>
                                <TableCell><Button variant="danger" onClick={() => this.handleDelete(member.member_id)}>DELETE</Button></TableCell>
                            </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Form>
            </div>

        )
    }
}
export default Result;
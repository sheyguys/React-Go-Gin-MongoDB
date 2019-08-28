import React, { Component } from "react"
import { Form } from "react-bootstrap"
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import "./Result.css"

class Result extends Component {
    state = {
        members: []
    };

    componentDidMount() {
       fetch('http://localhost:8081/employee')
        .then((response) => response.json())
        .then((data) => {
                const members = data;
                this.setState({members});
                console.log(data);
        });
      }
    render(){
        return (
            <div>
                <Form>
                    <h1>PERSONAL INFORMATION</h1>                
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
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.members.map(row => (
                            <TableRow key={row.Member}>
                                <TableCell >{row.member_name_eng}</TableCell>
                                <TableCell >{row.member_name_th}</TableCell>
                                <TableCell >{row.member_idcard}</TableCell>
                                <TableCell >{row.member_phone}</TableCell>
                                <TableCell >{row.member_address}</TableCell>
                                <TableCell >{row.member_email}</TableCell>
                                <TableCell >{row.member_facebook}</TableCell>
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
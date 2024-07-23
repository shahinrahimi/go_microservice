import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Toaster from './Toaster.js';

import axios from 'axios';

/*
This is the react equivilent of the following HTML form
<form action="http://localhost:8000" method="post" enctype="multipart/form-data">
  <p><input type="text" name="id" value="">
  <p><input type="file" name="file">
  <p><button type="submit">Submit</button>
</form>
*/

const Admin = () => {
    const [validated, setValidated] = React.useState(false)
    const [id, setId] = React.useState("")
    const [buttonDisabled, setButtonDisabled] = React.useState(false)
    const [toastShow, setToastShow] = React.useState(false)
    const [toastText, setToastText] = React.useState("asd")
    const [file, setFile] = React.useState("")

    const handleSubmit = (e: any) => {
        e.preventDefault()
        if (e.currentTarget.checkValidity() === false) {
            e.stopPropagation()
            return
        }
        setButtonDisabled(true)
        setToastShow(false)

        const formData = new FormData()
        formData.append("file", file)
        formData.append("id", id)

        axios.post(
            "http://localhost:7001/",
            formData
        ).then(res => {
            console.log(res)
            var txt = "";
            if (res.status === 200) {
                txt = "Uploaded file";
            } else {
                txt = "Unable to upload file. Error:" + res.statusText;
            }
            setButtonDisabled(false)
            setToastShow(false)
            setToastText(txt)
        }).catch(error => {
            console.log("Err", error)
            setButtonDisabled(false)
            setToastShow(false)
            setToastText("Unable to upload file. Error:" + error)
        })
    }

    const handleChange = (e: any) => {
        if (e.target.name === "file") {
            setFile(e.target.files[0])
            return
        }
        setId(e.target.value)
    }

    return (
        <div>
            <h1 style={{ marginBottom: "40px" }}>Admin</h1>
            <Container className="text-left">
                <Form noValidate validated={validated} onSubmit={handleSubmit}>
                    <Form.Group as={Row} controlId="productID">
                        <Form.Label column sm="2">Product ID:</Form.Label>
                        <Col sm="6">
                            <Form.Control type="text" name="id" placeholder="" required style={{ width: "80px" }} value={id} onChange={handleChange} />
                            <Form.Text className="text-muted">Enter the product id to upload an image for</Form.Text>
                            <Form.Control.Feedback type="invalid">Please provide a product ID.</Form.Control.Feedback>
                        </Col>
                        <Col sm="4">
                            <Toaster show={toastShow} message={toastText} />
                        </Col>
                    </Form.Group>
                    <Form.Group as={Row}>
                        <Form.Label column sm="2">File:</Form.Label>
                        <Col sm="10">
                            <Form.Control type="file" name="file" placeholder="" required onChange={handleChange} />
                            <Form.Text className="text-muted">Image to associate with the product</Form.Text>
                            <Form.Control.Feedback type="invalid">Please select a file to upload.</Form.Control.Feedback>
                        </Col>
                    </Form.Group>
                    <Button type="submit" disabled={buttonDisabled}>Submit form</Button>
                </Form>
            </Container>
        </div>
    )
}

export default Admin;
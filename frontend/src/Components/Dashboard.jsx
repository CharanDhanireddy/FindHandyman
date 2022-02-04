import React, { Component } from "react";
import { Container, Card, Modal, Form, Button, Row, Col } from "react-bootstrap";
import axios from "axios";
import TimeSelect from "./TimeSelect";

class Dashboard extends Component {
    constructor(props) {
        super(props)
        this.state = {
            cityList: null,
            city: null,
            showModal: true,
            serviceList: null,
            service: null
        }
    }

    async componentDidMount() {
        let res = await axios.get("http://localhost:5000/city")
        let data = res.data
        let status = res.status
        console.log(data, status)
        this.setState({ cityList: data.cities })
    }

    async componentDidUpdate(_prevProps, prevState) {
        if (this.state.city !== prevState.city) {
            let res = await axios.get("http://localhost:5000/service", {
                params: {
                    city: this.state.city
                }
            })
            let data = res.data
            this.setState({ serviceList: data.services })
        }
    }

    handleClose = () => {
        // check if city not null
        if (this.state.city)
            this.setState({ showModal: false })
    }

    render() {
        const { city, cityList, service, serviceList, showModal } = this.state;
        if (!cityList) return null
        return (
            <Container >
                {serviceList ? (
                    <>
                        <h2 className='text-center'>Select a service from the following</h2>
                        <Container className='mt-4'>
                            <Row xs={2} md={3} lg={4}>
                                {serviceList.map(service => (
                                    <Col>
                                        <Card
                                            key={service.id}
                                            style={{ height: '6rem', margin: '0 1rem 1rem 0', cursor: 'pointer' }}
                                            className='bg-dark text-white border'
                                            onClick={() => { console.log('here'); }}
                                        >
                                            <Card.Body className='mx-auto'>
                                                <h3 text-center>{service.name}</h3>
                                            </Card.Body>
                                            {/* <Button size="sm" variant="outline-primary">Select</Button> */}
                                        </Card>
                                    </Col>
                                )
                                )}
                            </Row>
                        </Container>
                    </>
                ) : (null)}

                <TimeSelect service={service} />

                <Modal
                    show={showModal}
                    backdrop="static"
                    aria-labelledby="contained-modal-title-vcenter"
                    centered
                >
                    <Modal.Header closeButton>
                        <Modal.Title>Select a city</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <Form>
                            <Form.Select
                                onChange={(e) => { this.setState({ city: e.target.value }) }}
                            >
                                <option value={null}>Select</option>
                                {cityList.map(city => (
                                    <option key={city.id} value={city.name}>{city.name}</option>
                                ))}
                            </Form.Select>
                        </Form>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="primary" disabled={!city} onClick={this.handleClose}>
                            Confirm
                        </Button>
                    </Modal.Footer>
                </Modal>
            </Container>
        );
    }
}

export default Dashboard;
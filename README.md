# HandyConnect
Class project for CEN5035 - Software Engineering

## Project Idea:
Platform for connecting customers with skilled and experienced professionals providing services (handymen, salon, applaince repair, etc).

## Team Members:
| Name | Gatorlink Email | UFID | Responsibility |
| :--: | :--: | :--: | :--: |
| Sai Charan Reddy Dhanireddy | s.dhanireddy@ufl.edu | 33614329 | Frontend |
| Srikanth Rao Parcha | parcha.srikanthr@ufl.edu | 46324560 | Frontend |
| Venkata Shiva Reddy Manchala | manchala.v@ufl.edu | 61099624 | Backend |
| Venkat Dhavaleswarapu | vdhavaleswarapu@ufl.edu | 72167758 |  Backend |

## Stack:
- Backend: Go Lang
- Frontend: ReactJS
- Database: Postgres

## Setup guide:
- Frontend: Follow the steps [here](https://github.com/CharanDhanireddy/HandyConnect/tree/frontend-dev/frontend#readme).

---
## Sprint 2:
[Integration Demo](https://www.youtube.com/watch?v=2YWmvZNdyHM)

### Frontend with live Golang backend & unit tests and integration tests with Cypress

![sprint2CypressGif](https://user-images.githubusercontent.com/20516490/156865184-b26d30da-ba5b-4e2f-9c77-fd02d1a21a44.gif)

### Test log for Back-end Unit Testing
![Backend_UT_1](https://user-images.githubusercontent.com/89587982/156866040-d6da90bc-e48d-42c8-be9f-838dd9cbc254.png)
![Backend_UT_2](https://user-images.githubusercontent.com/89587982/156866041-33524d75-fdba-4e64-b3d4-ac3b1dc36307.png)

Back-end Unit Tests can be executed by running "go test -v" in the *HandyConenct/Backend* directory

- TestCityRet - Test for validating City list retrieval success
- TestCityRet_Absent - Test for validating City list retrieval failure in case of invalid city
- TestCust - Test for validating successful retreival of Customer Information
- TestVend - Test for validating successful retreival of Vendor Information
- TestServices - Test for validating successful retreival of available services in a city
- TestBookingCust - Test for validating the retreival of Customer view of their bookings
- TestBookingVend - Test for validating the retreival of Vendor view of their bookings

Tasks completed:

- Added unit tests for all UI pages
- Cypress integration tests for all the workflows
- Updated API sections of React application to integrate with Live Golang backend
- Increased fields for SignUp pages to match with latest requirement
- UX improvements in Dashboard, Header and other pages
- Revamped UI for city selection

---

### Backend with mocked-up frontend
[Backend Demo video](https://youtu.be/YhrciuviPqM)

Tasks completed: (APIs for the below functionality are implemented and tested using Postman (shown in video))

- Customer can sign up with name, email, phone, city, address
- Vendor can sign up with their name, email, phone, city, services provided
- Customer/Vendor can log into their account with their email and password (error handling for incorrect password implemented)
- Customer/Vendor can login and view their profiles to see the data they've used for signing up
- Customer can get a list of available timeslots for a specific service in a city of their choice
- Customer can schedule a slot based on the availability of the service in a 1 week period

---

## Sprint 1:

### Frontend with mocked-up backend
[Frontend Demo video](https://youtu.be/5rjAeEBbr_4)

Tasks completed:

- Vendor can signup with vendor personal details along with selecting service and city
- Vendor/Customer can signup with name, email and password
- Vendor/Customer can login with email and password
- Customer can select a city before viewing avaliable services
- Customer can select a service to see available timeslots over a week
- Customer can select a timeslot and add their address to confirm the booking

---

### Backend with mocked-up frontend
[Backend Demo video](https://youtu.be/YhrciuviPqM)

Tasks completed: (APIs for the below functionality are implemented and tested using Postman (shown in video))

- Customer can sign up with name, email, phone, city, address
- Vendor can sign up with their name, email, phone, city, services provided
- Customer/Vendor can log into their account with their email and password (error handling for incorrect password implemented)
- Customer/Vendor can login and view their profiles to see the data they've used for signing up
- Customer can get a list of available timeslots for a specific service in a city of their choice
- Customer can schedule a slot based on the availability of the service in a 1 week period

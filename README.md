# Calculator
Over-engineered calculator. A basic calculator with a history. Exposes a REST API

## Installation
Clone repository
```bash
git clone https://github.com/gmmads/Calculator.git
```
Running the following from project root installs all dependencies, and starts the server:
```bash
go run .
```

## Usage
Postman collection: https://www.getpostman.com/collections/56ad5017b22c096dbe75

Can see history by making GET-request to localhost:8000/calculate

Can submit new calculation by making POST request to localhost:8000/calculate with body being a json-string such as the one below: 
```json
{
  "expr":"(1+2) * (10 / 5) - 42"
}
```
Accepted expression symbols are +, -, *, /, (, ) and any number.

Calculation is integer arithmetic. 

## Backlog (if I had more time)
Overflow protection (right now overflow just results in a negative number), 
support for floats and more interesting operations such as ^,
user auth (and a designated history for each individual user), 
website for using service.


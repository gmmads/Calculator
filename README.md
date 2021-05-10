# Calculator
Over-engineered calculator. A basic calculator with a history. Exposes a REST API. Deployed to https://safe-journey-27266.herokuapp.com/

## Installation
We can run a local version of the service. To do this first install Go: https://golang.org/. Then clone repository:
```bash
git clone https://github.com/gmmads/Calculator.git
```
Then running the following from project root installs all dependencies, and starts a local version of the server:
```bash
go run . test
```

## Usage
Postman collection: https://www.getpostman.com/collections/56ad5017b22c096dbe75

Can send GET-requests to https://safe-journey-27266.herokuapp.com/ and https://safe-journey-27266.herokuapp.com/calculations, and POST-request to https://safe-journey-27266.herokuapp.com/calculations (or localhost:8000/ and localhost:8000/calculations if using a local server).

Submit new calculation by making a POST-request to /calculations with body being a JSON-string such as the one below: 
```json
{
  "expr":"(1+2) * (10 / 5) - 42"
}
```
Accepted expression symbols are +, -, *, /, (, ) and any integer.

Returns a JSON-object:
```json
{
    "expr": "(1+2) * (10 / 5) - 42",
    "result": -36
}
```
or something like:
```json
{
    "message": "some error"
}
```
if an error occured (e.g. the expression used an unrecognised symbol, expression did not parse, or division by zero occured) 

A GET-request to /calculations will return a JSON-object with the history of all the calculations performed so far.

## Backlog (if I had more time)
Overflow protection (right now overflow is just allowed to occur, and no warning or error is raised), 
support for more interesting operations such as ^,
user auth (and a designated history for each individual user), 
website for using service.


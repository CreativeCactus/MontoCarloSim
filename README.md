# MonteCarloSim*

Estimate Pi by Monte Carlo simulation. Kinda neat!

Given the size of a 2D plane, the diameter of a circle within that plane, 
and an arbitrary number of randomly placed points on the plane, we can infer an aproximation of Pi.

This application consists of the following components:

- A computation backend written in golang and designed for scalability.
    It takes the aforementioned arguments and returns JSON with an estimate and a base64 encoded PNG.
- A server written in NodeJS with a simple exec call to the backend.
    Replace with call to Mesos or similar for cluster computation.
- A web frontend in HTML/CSS/JS with a bland UI for tweaking arguments and viewing the result.

## Installation

`git clone https://github.com/CreativeCactus/MontoCarloSim.git`
`cd MontoCarloSim`
`chmod +x build.sh`
`./build.sh`

## Usage

### Web application

To test and run the application as a whole, simply run `./e2e.sh`. Only docker is required.

### Backend

To run the computation backend, see the binary included, or `go build ./compute.go`.

Call the compute backend with `bin/compute -h` for help, or `bin/compute -grid=1000 -circ=900 -pts=100 -its=10`.


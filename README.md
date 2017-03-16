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

To test and run the application as a whole, simply run `./build.sh`. Only docker is required.

Build will compile the dockerfile, pulling `golang` and installing `node` and `jq`,
then call the image with `e2e.sh`, forwarding any arguments to `build.sh`.

Only arguments handled by `e2e.sh` are `buildonly` and `testonly`.
If neither is provided then the application will be tested, build, e2e tested, and left running.
If both are provided then the first is used.

Once the build runs once, it should be much quicker for subsequent runs. It will end with the listening IP and server port:


```
Starting server on IPs: 172.17.0.3
Server listening on http://127.0.0.1:3000
```

Navigate to the above host:port to see the GUI.

<img src="https://raw.githubusercontent.com/CreativeCactus/MontoCarloSim/master/web/src/gui.png" alt="gui" style="height:150px; width:250px; right: 0px; position:absolute;"></img>

### Backend

Beware big base64 output.

To run the computation backend, see the binary included, or `go build ./compute.go`.

Call the compute backend with `bin/compute -h` for help, or `bin/compute -grid=10 -circ=9 -pts=100 -its=10`.


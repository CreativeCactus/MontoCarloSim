window.onload = function () {
    // Set up all inputs to trigger an update()(e)
    const inputs = document.querySelectorAll('input');
    for (let elem of inputs) elem.onchange = update(inputs);
};

// update returns a curry function to handle events,
// which will validate inputs and request an updated estimate.
function update(inputs) {
    return () => {
        try {
            output('Request pending...', 'black');

            // Gather values and validate
            const values = {};
            for (let elem of inputs) {
                const integer = parseInt(elem.value); 
                if (isNaN(integer)) throw 'Invalid input';
                values[elem.id] = integer;
            }
        
            const req = new XMLHttpRequest();
            const URL = `/pi/${values.grid}/${values.circ}/${values.pts}/${values.its}`;
            req.open('GET', URL, true);
            req.onload = redraw.bind({ req });
            req.onerror = e => outputError(`Request failed: ${e}`);
            // req.onreadystatechange = redraw(req);
            req.send(null);
        } catch (e) {
            outputError(e);
        }
    };
}

// redraw returns a curry function to handle reception of the result
function redraw(e) {
    if (this.req.readyState !== 4) return;
    if (this.req.status !== 200) outputError(`Respone failed: ${this.req.statusText}`);
    
    result = document.querySelector('#result');
    try {
        const response = JSON.parse(this.req.responseText);
        if (response.error) outputError(`Response contains error: ${response.error}`);
        result.innerHTML = `Approximation of Pi: ${response.pi} <br> 
            <img alt="Response Image" src="data:image/png;base64,${response.png}" class="col-xs-12"/>`;
        output('Updated successfully', 'green');
    } catch (e) {
        outputError(`Response malformed: ${e}`);
    }
}

// outputError updates the output to show a failure
function outputError(e) {
    output(`Error: ${e}`, 'red');
}

// output updates the output element with some tect and a color
function output(text, colour) {
    const output = document.querySelector('#output');
    output.innerHTML = text;
    output.style.color = colour;
}

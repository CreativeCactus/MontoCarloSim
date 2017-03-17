const exec = require('child_process').exec;
const express = require('express');

const app = express();
const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
    exec('hostname -I', (error, stdout, stderr) => {
        if (error || stderr !== "") {
            console.log(`Server listening on http://127.0.0.1:${PORT}
            Could not, however, detect the local IP. Please install 'hostname'.`);
        }
        let IPs = stdout.split(' ');
        IPs = IPs.filter(v => ['\n', ''].indexOf(v) < 0);
        IPs = IPs.map(v => `http://${v}:${PORT}`);
        IPs = IPs.join('\n\t');
        console.log(`Server listening on interfaces: \n\t${IPs}`);
    });
});

app.get('/', (req, res) => {
    res.sendFile(`${__dirname}/src/frontend.html`);
});

app.use('/', express.static(`${__dirname}/src`));

app.get('/pi/:grid/:circ/:pts/:its', (req, res) => {
    const grid = parseInt(req.params.grid);
    const circ = parseInt(req.params.circ);
    const pts = parseInt(req.params.pts);
    const its = parseInt(req.params.its);

    const check = grid * circ * pts * its;
    if (isNaN(check) || check === 0) errHandler('Argument', check, res);

    // TODO sanitize further. Performing exec on user input can be dangerous.

    const cmd = `${process.env.PWD}/bin/compute -grid=${grid} -circ=${circ} -pts=${pts} -its=${its}`;
    exec(cmd, { maxBuffer: 10000 * 1 << 10 }, (error, stdout, stderr) => {
        if (error) return errHandler('Execution', error, res);

        if (stderr !== '') {
            console.error(stderr);
            return errHandler('Validation', 'See logs or contact admin.', res);
        }

        try {
            JSON.parse(stdout);
        } catch (e) {
            return errHandler('Parse', e, res);
        }

        res.send(stdout); // About 15Kb for the defaults
    });
});

function errHandler(type, err, res) {
    res.status(500);
    res.send('{error:"Compute failed, check logs or contact admin."}');
    console.error(`Time: ${Date.now()}
    ${type || 'General'} error: ${err}\n`);
}

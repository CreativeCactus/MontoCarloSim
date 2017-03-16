const exec = require('child_process').exec;
const express = require('express');
const app = express();
const PORT = process.env.PORT || 3000;

app.listen(PORT, () => { console.log(`Server listening on http://127.0.0.1:${PORT}`); });

app.get('/', (req, res) => {
    res.sendFile(`${process.env.PWD}/frontend.html`);
});
app.get('/frontend.js', (req, res) => {
    res.sendFile(`${process.env.PWD}/frontend.js`);
});
app.use('/web', express.static(`${process.env.PWD}/web/`));

app.get('/pi/:grid/:circ/:pts/:its', (req, res) => {
    const grid = parseInt(req.params.grid),
        circ = parseInt(req.params.circ),
        pts = parseInt(req.params.pts),
        its = parseInt(req.params.its);

    // TODO sanitize further. Performing exec on user input is dangerous at best.
    const cmd = `${process.env.PWD}/compute -grid=${grid} -circ=${circ} -pts=${pts} -its=${its}`;
    exec(cmd, { maxBuffer: 10000 * 1 << 10 }, (error, stdout, stderr) => { 
        if (error) {
            res.status(500);
            res.send('{error:"Compute failed, check logs or contact admin."}');
            console.error(`Time: ${Date.now()}`);
            console.error(`Compute error: ${error}`);
            console.error(`Stderr was: ${stderr}\n`);
            return;
        }
        res.send(stdout); // About 15Kb for the defaults
    });
});

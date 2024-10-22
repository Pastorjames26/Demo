const fs = require('fs');

let debugMode = true;

function debug(message) {
    if (debugMode) {
        console.log('DEBUG:', message);
    }
}

try {
    const data = fs.readFileSync('config.txt', 'utf8');
    debugMode = data.slice(-1) === '1';
} catch (err) {
    //do something or assume debugMode stays false
    console.log(err.message);
}

debug('debug message');
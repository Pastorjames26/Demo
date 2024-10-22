const crypto = require('crypto');

function sha256(data) {
    if (!sha256.cache) {
        console.log('initializing cache');
        sha256.cache = {};
    }

    if (!sha256.cache[data]) {
        console.log('cache miss');
        const hash = crypto.createHash('sha256');
        hash.update(data);
        sha256.cache[data] = hash.digest('hex');
    } else {
        console.log('cache hit');
    }

    return sha256.cache[data];
}

console.log(sha256('hello')); //cache miss
console.log(sha256('hello')); //cache hit
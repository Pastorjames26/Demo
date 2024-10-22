http.get('http://localhost:3000', (res) => {
    console.log('statusCode:', res.statusCode);

    res.on('data', (chunk) => {
        const data = JSON.parse(chunk.toString('utf8'));

        data.users.forEach((user) => {
            console.log(user.name);
        });
    });
}).on('error', (err) => {
    console.error('Request Error:', err);
});
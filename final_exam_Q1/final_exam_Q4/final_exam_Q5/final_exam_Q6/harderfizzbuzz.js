function harderFizzBuzz(n) {
    let result = [];

    for (let i = 1; i <= n; i++) {
        let strI = i.toString();
        let fizz = (i % 3 === 0 || strI.includes('3'));
        let buzz = (i % 5 === 0 || strI.includes('5'));

        if (fizz && buzz) {
        result.push("FizzBuzz");
        } else if (fizz) { 
        result.push("Fizz");
        } else if (buzz) { 
        result.push("Buzz"); 
        } else {
            result.push(strI);  
        }
    }

    return result;
} 

module.exports = harderFizzBuzz; 
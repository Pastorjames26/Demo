function reduceAndSortByFrequency(objects) {
    // Create a map to store value and its cumulative count 
    const frequencyMap = new Map();

    // Iterate over each object in the input array 
    objects.forEach(({ value, count }) => {
        // Convert the value to lowercase to ensure case-insensitvity
        const lowerValue = value.toLowerCase();

        // If the value already exist in the map, add the count to the exitsting count 
        if (frequencyMap.has(lowerValue)) {
            frequencyMap.set(lowerValue, frequencyMap.get(lowerValue) + count);
        } else {
            // Otherwise, set the value with the current count 
            frequencyMap.set(lowerValue, count);
        }
    });

    // Convert the map into an array of objects 
    const resultArray = Array.from(frequencyMap, ([value, count]) => ({value, count }));
    
    // Sort the array by count in descending order, and by value alphabetically if counts are the same
    resultArray.sort((a,b) => {
        if (b.count !== a.count) { 
            return b.count - a.count; // Sort by count (descending)
        }
        return a.value.localCompare(b.value); // Sort by value (ascending)
    });

    return resultArray;
}

module.exports = reduceAndSortByFrequency;
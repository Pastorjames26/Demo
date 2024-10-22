function navigateMaze(maze, instructions) {
    // Initalize starting position at the top-left corner of the maze
    let x = 0
    let y = 0;

    // Define the directions for the movement 
    const moves = {
        "up": [-1,0],
        "down": [1,0],
        "left": [0,-1],
        "right": [0,1]
    };

    // Loop through each instruction 
    for (let instruction of instructions) {
        // Get the direction of change for the current instruction
        const [dx, dy] = moves[instruction];

        // Update the current position
        x += dx;
        y += dy;

        // Check if the new position is out of bounds or is a wall
        if (x < 0 || x >= maze.length || y < 0 || y >= maze[0].length || maze[x][y] === 1) { 
            return false;
        }
    }

        // After all instructions are done, check if the current position is the end (9)
       return maze[x][y] === 9;
    }

    
module.exports = navigateMaze;
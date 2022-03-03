function isValidSudoku(board: string[][]): boolean {
    const rowContains = [];
    const colContains = [];
    const boxContains = [];

    if (board.length > 9 || board[0].length > 9) return false;

    for (let i = 0; i < 9; i++) {
        rowContains[i] = {};
        colContains[i] = {};
        boxContains[i] = {};
    }

    for (let i = 0; i < 9; i++) {
        for (let j = 0; j < 9; j++) {
            const digit = board[i][j];
            if (digit != ".") {
                if (rowContains[i][digit]) return false;
                rowContains[i][digit] = true;

                if (colContains[j][digit]) return false;
                colContains[j][digit] = true;

                const boxIndex = Math.floor(i / 3) * 3 + Math.floor(j / 3);
                if (boxContains[boxIndex][digit]) return false;
                boxContains[boxIndex][digit] = true;
            }
        }
    }

    return true;
}

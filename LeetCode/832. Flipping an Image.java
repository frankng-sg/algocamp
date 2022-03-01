class Solution {
    public int[][] flipAndInvertImage(int[][] A) {
        int height = A.length;
        int width = A[0].length;

        for (int i = 0; i < height; i++) {
            for (int j = (width - 1) / 2; j >= 0; j--) {
                if (A[i][j] != A[i][width - 1 - j]) {
                    A[i][j] = 1 - A[i][j];
                    A[i][width - 1 - j] = 1 - A[i][j];
                }
            }
        }

        for (int i = 0; i < height; i++) {
            for (int j = 0; j < width; j++) {
                A[i][j] ^= 1;
            }
        }

        return A;
    }
}

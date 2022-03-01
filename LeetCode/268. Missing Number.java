import java.util.HashMap;

class Solution {
    public int missingNumber(int[] nums) {
        Map<Integer, Integer> lookup = new HashMap<>();
        int N = 0;
        for (int num : nums) {
            lookup.put(num, 1);
            if (num > N) {
                N = num;
            }
        }
        
        for (int i = 0; i <= N; i++) {
            if (!lookup.containsKey(i)) {
                return i;
            }
        }
        return N + 1;
    }
}
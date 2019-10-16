import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> lookup = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            lookup.put(nums[i], i);
        }
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (lookup.containsKey(complement) && lookup.get(complement) != i) {
                return new int[] { i, lookup.get(complement) };
            }
        }
        throw new IllegalArgumentException("No two sum solution");
    }
}

import java.util.HashMap;
class Solution {
    public int singleNumber(int[] nums) {
        Map<Integer, Integer> lookup = new HashMap<>();
        for(int i = 0; i < nums.length; i++) {
            if (lookup.containsKey(nums[i])) {
                lookup.remove(nums[i]);
            }
            else {
                lookup.put(nums[i], 1);
            }        
        }
        
        for (Integer key : lookup.keySet()) {
            return key.intValue();
        }
        throw new IllegalArgumentException();
    }
}
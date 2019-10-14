class Solution {
    public char findTheDifference(String s, String t) {
        char output = 0;
        for(int i = 0; i < s.length(); i++) {
            output ^= s.charAt(i) ^ t.charAt(i);
        }
        output ^= t.charAt(t.length() - 1);
        return output;
    }
}
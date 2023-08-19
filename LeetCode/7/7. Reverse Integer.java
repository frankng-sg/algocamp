import java.lang.*;

class Solution {
    public int reverse(int x) {
        if (x < -Integer.MAX_VALUE || x > Integer.MAX_VALUE)
        {
            return 0;
        }

        int sign = 1;
        if (x < 0)
        {
            x = -x;
            sign = -1;
        }

        int i;
        String strx = Integer.toString(x);
        String rev_x = "";
        for (i = strx.length() - 1; i >= 0; i--)
        {
            rev_x += strx.charAt(i);
        }
        try
        {
            return sign * Integer.parseInt(rev_x);
        }
        catch (Exception e)
        {
            return 0;
        }

    }
}
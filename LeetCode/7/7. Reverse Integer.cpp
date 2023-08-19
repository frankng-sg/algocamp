#define INT_MAX 2147483647

class Solution {
public:
    int reverse(int x) {
        bool sign;
        int output;

        if (x <= -INT_MAX)
            return 0;

        sign = (x < 0) ? true : false;
        x = (x < 0) ? -x : x;
        output = 0;
        while (x > 0) {
            if (output > INT_MAX / 10)
                return 0;

            output = output * 10 + (x % 10);
            x /= 10;
        }
        if (sign)
            return -output;
        else
            return output;
    }
};
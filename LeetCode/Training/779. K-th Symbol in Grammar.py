class Solution:
    def kthGrammar(self, N: int, K: int, toggle=0) -> int:
        if (N == 1 or K == 0):
            return toggle

        if (K % 2 == 0):
            return self.kthGrammar(N - 1, K // 2, 1 - toggle)
        else:
            return self.kthGrammar(N - 1, K // 2 + 1, toggle)
        

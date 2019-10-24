class Solution:
    def generateParenthesis(self, n: int):
        queue = []
        open_char = '('
        close_char = ')'
        queue.append([open_char])
        for chars in range(1, 2 * n):
            for i in range(0, len(queue)):
                open_chars = queue[i].count(open_char)
                close_chars = queue[i].count(close_char)
                if not close_chars:
                    close_chars = 0
                unclosed = open_chars - close_chars

                was_extended = False
                new = queue[i].copy()
                if open_chars < n:
                    queue[i].extend([open_char])
                    was_extended = True

                if unclosed > 0:
                    if was_extended:
                        new.extend([close_char])
                        queue.append(new)
                    else:
                        queue[i].extend([close_char])

        for i in range(0, len(queue)):
            queue[i] = ''.join(queue[i])

        return queue

func = Solution()
print(func.generateParenthesis(3))
class Solution:
    def generateParenthesis(self, n: int):
        output = []
        sequence = []
        open_char = '('
        close_char = ')'

        def recur_search(remained_char, remained_open_char, unclosed_chars):
            if remained_char == 0:
                output.append(''.join(sequence))
                return
            if remained_open_char > 0:
                sequence.append(open_char)
                recur_search(remained_char - 1, remained_open_char - 1, unclosed_chars + 1)
                sequence.pop()

            if unclosed_chars > 0:
                sequence.append(close_char)
                recur_search(remained_char - 1, remained_open_char, unclosed_chars - 1)
                sequence.pop()

        recur_search(2 * n, n, 0)
        return output

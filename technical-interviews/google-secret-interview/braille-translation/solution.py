def get_braille_dict():
    braille = dict()
    braille["a"] = "100000"
    braille["b"] = "110000"
    braille["c"] = "100100"
    braille["d"] = "100110"
    braille["e"] = "100010"
    braille["f"] = "110100"
    braille["g"] = "110110"
    braille["h"] = "110010"
    braille["i"] = "010100"
    braille["j"] = "010110"
    braille["k"] = "101000"
    braille["l"] = "111000"
    braille["m"] = "101100"
    braille["n"] = "101110"
    braille["o"] = "101010"
    braille["p"] = "111100"
    braille["q"] = "111110"
    braille["r"] = "111010"
    braille["s"] = "011100"
    braille["t"] = "011110"
    braille["u"] = "101001"
    braille["v"] = "111001"
    braille["w"] = "010111"
    braille["x"] = "101101"
    braille["y"] = "101111"
    braille["z"] = "101011"
    braille["!"] = "000001"
    braille[" "] = "000000"
    return braille


def solution(s):
    braille = get_braille_dict()
    out = ""
    for c in s:
        if c.isupper():
            out += braille["!"]
        out += braille[c.lower()]
    return out


if __name__ == "__main__":
    print(solution("code"))  # 100100101010100110100010
    print(solution("Braille"))  # 000001110000111010100000010100111000111000100010
    print(
        solution(
            "The quick brown fox jumps over the lazy dog"
        )  # 000001011110110010100010000000111110101001010100100100101000000000110000111010101010010111101110000000110100101010101101000000010110101001101100111100011100000000101010111001100010111010000000011110110010100010000000111000100000101011101111000000100110101010110110
    )

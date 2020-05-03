class Solution:

    def canConstruct(self, ransomNote: str, magazine: str) -> bool:

        letters = [0] * (ord('z') - ord('a') + 1)

        for letter in magazine:
            letters[ord(letter) - ord('a')] += 1

        for letter in ransomNote:

            index = ord(letter) - ord('a')
            if letters[index] == 0:
                return False

            letters[index] -= 1

        return True


from itertools import zip_longest

class Solution:

    def reader(self, s):

        back = 0
        for c in reversed(s):

            if c == '#':
                back += 1
                continue

            if back > 0:
                back -= 1
                continue

            yield c

        yield None


    def backspaceCompare(self, S: str, T: str) -> bool:

        for s, t in zip_longest(self.reader(S), self.reader(T)):
            if s != t:
                return False

        return True


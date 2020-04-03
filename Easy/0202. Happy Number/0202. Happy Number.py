class Solution:

    def isHappy(self, n: int) -> bool:

        def f(x: str):
            return int(x)**2

        buffer = set([n])
        while True:

            n = sum(map(f, list(str(n))))

            if n == 1:
                return True

            if n in buffer:
                return False

            buffer.add(n)


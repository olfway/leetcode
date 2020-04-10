class MinStack:

    def __init__(self):
        self.s = list()

    def push(self, x: int) -> None:
        if not self.s:
            self.s.append((x, x))
        else:
            self.s.append((x, min(x, self.s[-1][1])))

    def pop(self) -> None:
        self.s.pop()

    def top(self) -> int:
        return self.s[-1][0]

    def getMin(self) -> int:
        return self.s[-1][1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()

class Node:
    def __init__(self, value):
        self.prev = None
        self.next = None
        self.value = value

class FirstUnique:

    def __init__(self, nums: List[int]):

        self.uniq = dict()
        self.first = self.last = None

        for num in nums:
            self.add(num)

    def add(self, value: int) -> None:

        node = self.uniq.get(value, None)

        if node is None:

            node = Node(value)
            self.uniq[value] = node
            node.prev = self.last

            if self.first is None:
                self.first = node

            if self.last is not None:
                self.last.next = node

            self.last = node

            return

        if node is False:

            return

        self.uniq[value] = False

        if node.next:
            node.next.prev = node.prev
        else:
            self.last = node.prev

        if node.prev:
            node.prev.next = node.next
        else:
            self.first = node.next


    def showFirstUnique(self) -> int:

        if self.first:
            return self.first.value

        return -1


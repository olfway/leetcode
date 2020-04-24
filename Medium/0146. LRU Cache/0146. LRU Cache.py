class DoubleLinkedList:

    def __init__(self, key, value, next=None):

        self.key = key
        self.value = value
        self.prev = None
        self.next = None

        self.setNext(next)

    def pop(self):

        if self.prev is not None:
            self.prev.next = self.next

        if self.next is not None:
            self.next.prev = self.prev

        self.prev = self.next = None

    def setNext(self, next):

        if next is not None:
            self.next = next
            next.prev = self


class LRUCache:

    def __init__(self, capacity: int):

        self.first = None
        self.last = None
        self.values = dict()
        self.capacity = capacity

    def makeFirst(self, item):

        if self.first == item:
            return

        if self.last == item:
            self.last = item.prev

        item.pop()
        item.setNext(self.first)
        self.first = item

        if self.last is None:
            self.last = item

    def get(self, key: int) -> int:

        if key not in self.values:
            return -1

        self.makeFirst(self.values[key])

        return self.first.value

    def put(self, key: int, value: int) -> None:

        if key in self.values:
            self.makeFirst(self.values[key])
            self.first.value = value
            return

        if len(self.values) == self.capacity:
            last = self.values.pop(self.last.key)
            self.last = last.prev

        item = DoubleLinkedList(key, value, next=self.first)

        self.first = item
        if self.last is None:
            self.last = item

        self.values[key] = item


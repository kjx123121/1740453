class Stack:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.insert(0, item)

    def pop(self):
        try:
            return self.items.pop(0)
        except IndexError:
            return IndexError("pop from empty list")

    def peek(self):
        try:
            return self.items[0]
        except IndexError:
            return IndexError("list index out of range")

    def size(self):
        return len(self.items)

    def is_empty(self):
        return self.items == []



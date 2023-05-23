# this is a comment

"""
This is a multi-line comment
"""

def add(*args):
    sum = 0
    for x in args:
        sum += x
    return sum


print(add(10, 2, 30))

# data types
a = 10
b = 2.30
c = True
d = "Hello world"
e = None

print(10 / 20)
print(10 // 20)
print(15 % 3)
print(not False)
print(a == b)

l1 = [1,2,3,4]
l2 = l1
print(l1 == l2)
print(l1 is l2)

l2 = [1,2,3,4]
print(l1 == l2)
print(l1 is l2)

print(len(d))

name = "Mark"
print("Hello {name}".format(name=name))
print(f"Hello {name}")

l1.append(10)
poped = l1.pop()
print(poped)

print(l1[0], l1[-1], l1[1:3], l1[::2], l1[::-1])

l3 = l1[:] # l3 is l1 => False | l3 == l1 => True
del l1[0]
print(l1)

l1.insert(1, 10)
print(l1)
print(l1.index(10))

print(l1 + l2)
print(10 in l1)
print(len(l1))

tup = (1, 2, 3, 4, 5)
a, *b, c = tup
print(a, b, c)

my_dict = {"one": 1, "two": 2, "three": 3}
print(my_dict["one"])
keys = my_dict.keys() # iterable
for k in keys:
    print(k)
vals = my_dict.values()
for v in vals:
    print(v)

my_set = {1, 2, 3, 3, 4}
print(my_set)
my_set.add(10)
print(my_set)

for i, v in enumerate(l1):
    print(f"{i}: {v}")

import json

with open("myfile.json", "w") as fhand:
    json.dump(my_dict, fhand)

# lambda functions

sum = (lambda x, y: x + y)(10, 20)
print(sum)

# maps

def add_10(args):
    mod = []
    for x in args:
        mod.append(x + 10)
    return mod

sum = list(map(lambda x: x + 10, [1,2,3,4]))
print(sum)

newl = list(filter(lambda x: x < 3, [1,2,3,4,5]))
print(newl)

# list comprehension
print([x**2 for x in range(10) if x <= 5])

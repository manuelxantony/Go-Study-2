import sys

f = open("data.txt", "r")


def input():
    ff = f.readline()
    return ff


for _ in range(int(input())):
    n, m = map(int, input().split())
    c = 0
    for i in range(m):
        x, y = map(int, input().split())
        if x == y or (x > n and x > y and x > n + y):
            c = 1
    print("NO" if c else "YES")

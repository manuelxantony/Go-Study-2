def c():
    global a
    a = 4
    chan(a)
    print(a)


def chan(b):
    global a
    a = 8
    print(a)


c()

from sys import stdin
import bisect


def longest(arr, n, k):
    dyn = []
    dyn.append(arr[0])
    for i in range(1, n, 1):
        if arr[i] > dyn[-1]:
            dyn.append(arr[i])
        else:
            j = bisect.bisect_left(dyn, arr[i])
            if j > 0:
                if arr[i] - dyn[j - 1] >= k:
                    dyn[j] = arr[i]
            else:
                dyn[j] = arr[i]

    i = 1
    while i < len(dyn):
        if dyn[i] - dyn[i - 1] < k:
            dyn.pop(i)
        else:
            i += 1

    print(len(dyn))


if __name__ == "__main__":
    f = open("data.txt", "r")
    k = int(f.readline())
    n = int(f.readline())
    if n == 0:
        print(0)
    else:
        arr = list(map(int, f.readline().split()))
        longest(arr, n, k)

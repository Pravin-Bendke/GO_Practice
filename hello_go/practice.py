arr = [ 2,1, 3, 4, 7, 6, 7]

largest = arr[0]
second = arr[0]

for i in range(1, len(arr)):
    if arr[i] > largest:
        second = largest
        largest = arr[i]
    elif arr[i] > second and arr[i] != largest:
        second = arr[i]

print(second)


for i in range(65, 123):
    print(chr(i),end=" | ")


# maxnum = sorted(set(arr))[-2]
# print(maxnum)
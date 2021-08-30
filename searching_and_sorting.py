def binary_search(numbers, key):
    low = 0
    high = len(numbers) - 1

    while high >= low:
        mid = (high + low) // 2
        if numbers[mid] < key:
            low = mid + 1
        elif numbers[mid] > key:
            high = mid - 1
        else:
            return mid
    return -1  # not found



# selection sort
def selection_sort(numbers):
    for i in range(len(numbers) - 1):
        # Find index of smallest remaining element
        index_smallest = i
        for j in range(i + 1, len(numbers)):
            if numbers[j] < numbers[index_smallest]:
                index_smallest = j

        # Swap numbers[i] and numbers[index_smallest]
        temp = numbers[i]
        numbers[i] = numbers[index_smallest]
        numbers[index_smallest] = temp


numbers = [10, 2, 78, 4, 45, 32, 7, 11]
print('UNSORTED:', end=' ')
for num in numbers:
    print(str(num), end=' ')
print()

selection_sort(numbers)
print('SORTED:', end=' ')
for num in numbers:
    print(str(num), end=' ')
print()

# Insertion Sort
def insertion_sort(numbers):
    for i in range(1, len(numbers)):
        j = i
        # Insert numbers[i] into sorted part
        # stopping once numbers[i] in correct position
        while j > 0 and numbers[j] < numbers[j - 1]:
            # Swap numbers[j] and numbers[j - 1]
            temp = numbers[j]
            numbers[j] = numbers[j - 1]
            numbers[j - 1] = temp
            j = j - 1


numbers = [10, 2, 78, 4, 45, 32, 7, 11]
print('UNSORTED:', end=' ')
for num in numbers:
    print(str(num), end=' ')
print()

insertion_sort(numbers)
print('SORTED:', end=' ')
for num in numbers:
    print(str(num), end=' ')
print()



# quick sort
def partition(numbers, i, k):
    #  Pick middle element as pivot
    midpoint = i + (k - i) // 2
    pivot = numbers[midpoint]

    #  Initialize variables
    done = False
    l = i
    h = k
    while not done:
        #  Increment l while numbers[l] < pivot
        while numbers[l] < pivot:
            l = l + 1
        #  Decrement h while pivot < numbers[h]
        while pivot < numbers[h]:
            h = h - 1
        """  If there are zero or one items remaining,
              all numbers are partitioned. Return h """
        if l >= h:
            done = True
        else:
            """  Swap numbers[l] and numbers[h],
                  update l and h """
            temp = numbers[l]
            numbers[l] = numbers[h]
            numbers[h] = temp
            l = l + 1
            h = h - 1
    return h


def quicksort(numbers, i, k):
    j = 0
    """  Base case: If there are 1 or zero entries to sort,
          partition is already sorted  """
    if i >= k:
        return
    """  Partition the data within the array. Value j returned
          from partitioning is location of last item in low partition. """
    j = partition(numbers, i, k)
    """  Recursively sort low partition (i to j) and
          high partition (j + 1 to k) """
    quicksort(numbers, i, j)
    quicksort(numbers, j + 1, k)
    return


numbers = [10, 2, 78, 4, 45, 32, 7, 11]
print ('UNSORTED:', end=' ')
for num in numbers:
    print (str(num), end=' ')
print()

#  Initial call to quicksort
quicksort(numbers, 0, len(numbers) - 1)
print ('SORTED:', end=' ')
for num in numbers:
    print (str(num), end=' ')
print()

# merge sort
def merge(numbers, i, j, k):
    merged_size = k - i + 1  # Size of merged partition
    merged_numbers = []  # Temporary list for merged numbers
    for l in range(merged_size):
        merged_numbers.append(0)

    merge_pos = 0  # Position to insert merged number

    left_pos = i  # Initialize left partition position
    right_pos = j + 1  # Initialize right partition position

    #  Add smallest element from left or right partition to merged numbers
    while left_pos <= j and right_pos <= k:
        if numbers[left_pos] < numbers[right_pos]:
            merged_numbers[merge_pos] = numbers[left_pos]
            left_pos = left_pos + 1
        else:
            merged_numbers[merge_pos] = numbers[right_pos]
            right_pos = right_pos + 1
        merge_pos = merge_pos + 1

    #  If left partition is not empty, add remaining elements to merged numbers
    while left_pos <= j:
        merged_numbers[merge_pos] = numbers[left_pos]
        left_pos = left_pos + 1
        merge_pos = merge_pos + 1

    #  If right partition is not empty, add remaining elements to merged numbers
    while right_pos <= k:
        merged_numbers[merge_pos] = numbers[right_pos]
        right_pos = right_pos + 1
        merge_pos = merge_pos + 1

    #  Copy merge number back to numbers
    merge_pos = 0
    while merge_pos < merged_size:
        numbers[i + merge_pos] = merged_numbers[merge_pos]
        merge_pos = merge_pos + 1


def merge_sort(numbers, i, k):
    j = 0
    if i < k:
        j = (i + k) // 2  # Find the midpoint in the partition

        #  Recursively sort left and right partitions
        merge_sort(numbers, i, j)
        merge_sort(numbers, j + 1, k)

        #  Merge left and right partition in sorted order
        merge(numbers, i, j, k)


numbers = [10, 2, 78, 4, 45, 32, 7, 11]
print('UNSORTED:', end=' ')
for num in numbers:
    print(str(num), end=' ')
print()

#  initial call to merge_sort with index
merge_sort(numbers, 0, len(numbers) - 1)
print('SORTED:', end=' ')
for num in numbers:
    print(str(num), end=' ')
print()
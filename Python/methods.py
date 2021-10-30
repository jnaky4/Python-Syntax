# argument list *arg sets at tuple
def sandwich(bread, meat, *args):
    print('{} on {}'.format(meat, bread), end=' ')
    if len(args) > 0:
        print('with', end=' ')
    for extra in args:
        print(extra, end=' ')
    print('')


sandwich('sourdough', 'turkey', 'mayo')
sandwich('wheat', 'ham', 'mustard', 'tomato', 'lettuce')

# kwargs, passing dictionaries
def sandwich(bread, meat, **kwargs):
    print('{} on {}'.format(bread, meat))
    for category, extra in kwargs.items():
        print('   {}: {}'.format(category, extra))

sandwich('sourdough', 'turkey', sauce='mayo')
sandwich('wheat', 'ham', sauce1='mustard', veggie1='tomato', veggie2='lettuce')

# multi output methods
# returns tuple, have to unpack
student_scores = [75, 84, 66, 99, 51, 65]


def get_grade_stats(scores):
    # Calculate the arithmetic mean
    mean = sum(scores) / len(scores)

    # Calculate the standard deviation
    tmp = 0
    for score in scores:
        tmp += (score - mean) ** 2
    std_dev = (tmp / len(scores)) ** 0.5

    # Package and return average, standard deviation in a tuple
    return mean, std_dev


# Unpack tuple
average, standard_deviation = get_grade_stats(student_scores)

print('Average score:', average)
print('Standard deviation:', standard_deviation)
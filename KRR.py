import sys

def read_file(name):
    f = open(name, 'r')
    line = f.readline()
    f.close()
    return line

def get_numbers(line):
    numbers = []
    parts = line.split()
    for p in parts:
        try:
            numbers.append(float(p))
        except:
            pass
    return numbers

def do_sum(numbers):
    s = 0
    for n in numbers:
        s = s + n
    return s

def do_mul(numbers):
    r = 1
    for n in numbers:
        r = r * n
    return r

def do_avg(numbers):
    if len(numbers) == 0:
        return 0
    s = do_sum(numbers)
    return s / len(numbers)

def save_to_file(name, result):
    f = open(name, 'w')
    f.write(str(result))
    f.close()

args = sys.argv[1:]

input_file = 'input.txt'
output_file = None
operation = None

i = 0
while i < len(args):
    if args[i] == '--output':
        output_file = args [i + 1]
        i = i + 2
    elif args[i] == '-sum':
        operation = '-sum'
        i = i + 1
    elif args[i] == '-mul':
        operation = '-mul'
        i = i + 1
    elif args[i] == '-avg':
        operation = '-avg'
        i = i + 1
    else:
        i = i + 1

line = read_file(input_file)
numbers = get_numbers(line)
result = 0

if operation == '-sum':
    result = do_sum(numbers)
elif operation == '-mul':
    result = do_mul(numbers)
elif operation == '-avg':
    result = do_avg(numbers)

if output_file != None:
    save_to_file(output_file, result)
else:
    print(result)
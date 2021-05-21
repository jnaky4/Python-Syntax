import sys, os


#TODO add argparser

# first argument is the filename
if len(sys.argv) > 1:
    print(sys.argv[0], sys.argv[1])
    try:
        f = open(sys.argv[1], "r")
        print(f.readline())
    except Exception as e:
        print(e)



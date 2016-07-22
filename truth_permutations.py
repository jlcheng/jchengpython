#!//usr/bin/env python

from os.path import dirname, join
import sys
import pprint
import copy


def conv(intval):
    if intval == 0:
        return 'True'
    else:
        return 'False'

def main(args):
    m = create_matrix(args.count)
    for row in m:
        print ','.join(map(conv,row))

def create_matrix(size):
    if size == 1:
        return [[0],[1]]
    else:
        matrix = create_matrix(1)
        for i in range(size-1):
            matrix = grow_matrix(matrix)
        return matrix
    
def grow_matrix(input):
    output = []
    for row in input:
        row_copy = copy.deepcopy(row)
        row_copy.append(0)
        output.append(row_copy)
        row_copy = copy.deepcopy(row)
        row_copy.append(1)
        output.append(row_copy)

    return output

if __name__  == "__main__":
    import argparse

    parser = argparse.ArgumentParser()
    parser.description = 'Generate a table of permutations of True and False'
    parser.add_argument('--count', default=2, type=int, help='The width of the table, defaults to 2')
    args = parser.parse_args()
    main(args=args)

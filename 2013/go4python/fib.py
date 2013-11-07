#!/usr/bin/python

def fib(n):
    a, b = 0, 1
    for i in range(n):
        a, b = b, a + b
    return b

def fib_rec(n):
	if n <= 1:
		return 1
	else:
		return fib_rec(n-1) + fib_rec(n-2)

for x in range(10):
	print fib(x), fib_rec(x)

#!/usr/bin/python

def fib(n):
    a, b = 0, 1
    for i in range(n):
        a, b = b, a + b
        yield a

f = fib(10)
try:
	while True:
		print f.next()
except StopIteration:
	print 'done'

for x in fib(10):
	print x
print 'done'

#!/usr/bin/python

import urllib

auth_url = 'http://google.com'

def auth(usr):
	try:
		r = urllib.urlopen(auth_url + '/' + usr)
		return r.getcode() == 200
	except:
		return False

def say_hi(usr):
	if auth(usr):
		print 'Hi, %s' % usr
	else:
		print 'unknown user %s' % usr

def sayhitest():
	# Test authenticated user
	globals()['auth'] = lambda x: True
	say_hi('John')

	# Test unauthenticated user
	globals()['auth'] = lambda x: False
	say_hi('John')
#done OMIT

sayhitest()

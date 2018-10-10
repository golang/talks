#!/usr/bin/python

from BaseHTTPServer import BaseHTTPRequestHandler,HTTPServer
from urlparse import urlparse,parse_qs

PORT_NUMBER = 8080

def auth_required(myfunc):
	def checkuser(self):
		user = parse_qs(urlparse(self.path).query).get('user')
		if user:
			self.user = user[0]
			myfunc(self)
		else:
			self.wfile.write('unknown user')
	return checkuser


class myHandler(BaseHTTPRequestHandler):
	@auth_required
	def do_GET(self):
		self.wfile.write('Hello, %s!' % self.user)

try:
	server = HTTPServer(('', PORT_NUMBER), myHandler)
	server.serve_forever()

except KeyboardInterrupt:
	server.socket.close()

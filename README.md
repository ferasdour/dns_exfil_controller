# DNS Exfil over Interactsh
So, long story short here; I made a blog post (https://feemcotech.blogspot.com/2025/05/pyscript-nim-aaaaand-go.html) about my experiences learning different lanaguages building example red team malware. I then decided I'd take those and make a simple controller for them since I was able to show a client that detections worked in some scenarios but not others. 

That's what this is. But I didn't leave it at just that yet. I wanted to see how different detections work against different programming languages implementing the exact same logic and for the tool to be able to spin those up.

## to build controller
```
go build .
```

## to build samples
check out the dependencies, I had some issues with nim throwing windows.h and other c library files that were accessible and in the path and what not. So I had to add additional options to the nim compile to get it to work

## crypting
This wasn't considered when making this, but there are many ways to hide this.

## Reading the exfil file
I know I could implement better logic, maybe use a special encoded option for "start new file" then give name, then "start data". but it seemed so much easier to just grab everything into one file (since it generates new binary on each run of the controller, to a new interactsh client callback) and use binwalk, foremost, or strings, hexdump, whatever to pull information out of it. 

The intention for this isn't to be used maliciously, but to showcase an ability for real attackers to exfiltrate data in a way that immediately compromises everything that was collected. That along with testing behavioral alerting instead of file by file alerting (one of the many benefits to modern EDR), seemed like a tool worth sharing.

## Other ways to run
This can be imported into a container, I haven't made one yet, but it seems pretty stright forward; ubuntu latest, install stuff needed, start with a file mount. Maybe use a bash script to start this so it generates a new local file for each container as it runs to mount for the /tmp/exfil, and maybe add a variable for the .exe name to change with each container. idk, maybe i'll work on that later. For now, i'm having more fun just going through different languages running this. 

## running examples:
I did 3 uploads:
```
https://app.any.run/tasks/c7131923-861c-48c6-a802-5ba55904054b - go test
https://app.any.run/tasks/1fdfa6c1-aca0-453b-a556-25ac5741016e - nim test
https://app.any.run/tasks/ad36533e-6214-40c2-8fcc-762b3e7706e9 - second go test (longer time)

hexdump -Cv /tmp/exfil
00000000  15 00 01 00 02 00 02 00  00 00 00 00 b8 24 cc 60  |.............$.`|
00000010  10 00 00 00 00 00 00 00  00 00 00 ff ff ff ff 00  |................|
00000020  00 00 00 ff ff ff f7 00  00 00 00 00 00 00 00 00  |................|
00000030  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000060  ff fe 0d 00 0a 00 5b 00  2e 00 53 00 80 06 50 06  |......[...S...P.|
00000070  c0 06 c0 04 30 06 c0 06  00 73 00 73 00 49 00 6e  |....0....s.s.I.n|
00000080  00 66 00 6f 05 d0 00 d0  00 a0 04 c0 06 f0 06 30  |.f.o...........0|
00000090  05 d0 00 d0 00 a0 04 c0  06 f0 06 30 05 d0 00 d0  |...........0....|
000000a0  00 a0 04 c0 06 f0 06 30  05 d0 00 d0 00 a0 04 c0  |.......0........|
000000b0  06 f0 06 30 61 00 6c 00  69 00 7a 00 65 00 64 00  |...0a.l.i.z.e.d.|
000000c0  61 00 6c 00 69 00 7a 00  65 00 64 00 61 00 6c 00  |a.l.i.z.e.d.a.l.|
000000d0  69 00 7a 00 65 00 64 00  61 00 6c 00 69 00 7a 00  |i.z.e.d.a.l.i.z.|
000000e0  65 00 64 00 20 06 50 07  30 06 f0 07 50 07 20 06  |e.d. .P.0...P. .|
000000f0  00 65 00 4e 00 61 00 6d  00 65 00 3d 04 00 02 50  |.e.N.a.m.e.=...P|
00000100  05 30 07 90 07 30 07 40  65 00 6d 00 52 00 6f 00  |.0...0.@e.m.R.o.|
00000110  6f 00 74 00 50 05 c0 07  30 07 90 07 30 07 40 06  |o.t.P...0...0.@.|
00000120  00 6d 00 33 00 32 00 5c  00 77 00 69 06 e0 06 40  |.m.3.2.\.w.i...@|
00000130  06 f0 07 70 07 30 02 e0  73 00 74 00 6f 00 72 00  |...p.0..s.t.o.r.|
00000140  61 00 67 00 50 02 e0 06  40 06 c0 06 c0 02 c0 02  |a.g.P...@.......|
00000150  00 32 00 31 00 38 00 32  00 35 00 0d 00 a0 04 90  |.2.1.8.2.5......|
00000160  06 30 06 f0 06 e0 05 20  65 00 73 00 6f 00 75 00  |.0..... e.s.o.u.|
00000170  72 00 63 00 50 03 d0 02  50 05 30 07 90 07 30 07  |r.c.P...P.0...0.|
00000180  00 65 00 6d 00 52 00 6f  00 6f 00 74 02 50 05 c0  |.e.m.R.o.o.t.P..|
00000190  07 30 07 90 07 30 07 40  65 00 6d 00 33 00 32 00  |.0...0.@e.m.3.2.|
000001a0  5c 00 69 00 d0 06 10 06  70 06 50 07 20 06 50 07  |\.i.....p.P. .P.|
000001b0  00 2e 00 64 00 6c 00 6c  00 2c 00 2d              |...d.l.l.,.-|
```

# Credits:
- First and foremost, this all started by replicating with slight modifications a tool on https://github.com/byt3bl33d3r/OffensiveNim
- google
- The many other offensive(programminglanguage) repos out there

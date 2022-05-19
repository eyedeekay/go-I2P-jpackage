# go-I2P-jpackage

Embeddable, updatable alternative to NSIS based installer for portable Java I2P apps. Should
maybe work on Linux and Windows. OSX users should look to Zlatinb's Easy-Install DMG. Most
Windows users should probably refer to the Windows Easy-Install Bundle. Go developers who
are especially fond of Java I2P and who are building I2P apps might find something to offer
here.

**UNBELIEVABLY Experimental and probably totally unsuitable for use.**

It is impossible to use except with a specific branch of i2p.firefox, which allows overriding
the install and configuration path of the jpackaged I2P router. It is intended as part of a
"portable" in the sense that the installation is self-contained to a directory and includes
all the software which is required to adequately make use of it. Want to build an I2P-Based
communications app for Windows that your users run from a flash drive? This... might be the
way. If you're brave. Or an idiot. I don't know.

The executable produced by this is a standalone binary that can be built for any platform,
but it does not support cross-compilation. It is designed to be `go get` able and `import`
able on any platform supported by Java 17 an Jpackage, but it **REQUIRES** all the same
build dependencies as `i2p.firefox`. `i2p.firefox` uses `make` for a lot of stuff which
is obnoxious to use on Windows and totally my fault, and also totally unnecessary. If it
turns out to be a non-stupid idea, then I will cut those build dependencies down to just
`bash`. Because I literally cannot imagine the person who does Go development on Windows
without at least `git bash` and if I met them, I would probably not like them.

To produce the executable, run:

```sh
go generate
```

instead of go build. That's because it plays some tricks on Go to generate a jpackaged
router while being built.

In order to use this *as a library*, you will need Java 14+(17 recommended), bash, Go
and Git installed. The binary produced by importing the library will work **only** upon
the platform that it is built on. You should clone it into your GOPATH:

`git clone https://github.com/eyedeekay/go-I2P-jpackage $GOPATH/src/github.com/eyedeekay/go-I2P-jpackage`

then generate the code for your platform locally:

```bash
cd $GOPATH/src/github.com/eyedeekay/go-I2P-jpackage
```

finally, add a "replace" directive which refers to the local checkout of go-I2P-jpackage
to your `go.mod` file and run `go build`. Complete the indicated steps to update your go.sum
file.
s
## Why? In god's name why would I waste my time on some bullshit like this?

Well. The idea, stupid as it might be, is that everything that imports this package is
potentially a valid UpdatePostProcessor type 6(EXE) for the platform it's on. So as
long as the vendor keeps producing updates, they can in theory distribute them using
an I2P style signed newsfeed. Which is arguably a non-boring thing to be able to do.

As long as you also show people how to set up a signed newsfeed?
And you trust them to like, actually release updates in a timely and secure manner?

... It might even be a good idea.

# The Big TODO:

How to host a secure, signed newsfeed to provide updates for your app which embeds I2P?
Because that's what this does, at least, in a perfect world.
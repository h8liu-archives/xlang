[![BuildStatus](https://travis-ci.org/h8liu/xlang.png?branch=master)](https://travis-ci.org/h8liu/xlang)

# Xlang

Xlang is my toy language as my effort to enable writing comprehensible code,
where people can easily take some piece of code, read it, understand it,
use it, wrap it, or even change it, into better code.

The plan is:

- Pick a simple enough architecture.
- Design a language and write a compiler for the language.
- Write a simple operating system in the language.
- Rewrite the compiler in the language so that it is self-hosting.

The goal is to make all the code written in this project comprehensible for
human programmers and readers. At least, all the code that is written in the
new language needs to be comprehensible for human. I will put some certain care
on the language design and the project architecture so that it can be as
comprehensible as possible. I will also ask my friends to review my code and
evaluate if the code is comprehensible or not.

I plan to have a coding blog that logs the entire procedure.

I have several previous tries including [E8](http://github.com/e8vm) and
[Luna](http://github.com/h8liu/luna).  E8 is a simple toy architecture, and
Luna is a subset of ARM architecture, which I plan to play with on Raspberry
Pi's.  However, picking and simulating a simple architecture is relatively
simple, while building a programming language power enough to host itself and
an operating system (in a comprehensible way) is much harder. Plus, if I have a
working compiler with a relatively good IR design, porting it to multiple
architectures seems not a very complicated task.

So, I paused my construction on E8 and Luna, and started to just focus on my
programming langauge first.

# Design Thoughts

I like Go language very much, and in my opinion, it is the most readable and
comprehensible language in the world, for these features:

- It is a procedural language, which means the program describes what it does
  in the same way that most people think about algorithms: as steps, but not
  functions.
- It has a succinct, clean grammar, which means simple stuff is written in
  simple ways.
- It has very few language features, which means complex stuff (like error handling
  for corner cases) must be written in complex ways.
- It has no macros or templates or other meta-programming features, which means
  the source code is statically presented as it is to the code generator.
- It has packages but no circular dependencies, which means even for a very large
  project, you can always start reading (and playing with) the code from the 
  very bottom of the dependency tree.

As a result, Go language is the only language that I can really read and sort
of understand its stardard library.

There are still some nitches for Go language:

- A package can be really big and becomes hard to understand (like the
  `net/http` package), where splitting it to multiple smaller packages
  is sometimes not encouraged (because packages is essentially a library
  that does well on some particular single thing, and `http` is one big
  single thing.)
- Poor package management. The Go language guys at Google just do not care,
  and I really don't see how the open source community can figure this thing
  out automatically.

Also some other reasons that I want to invent (again) a new language:

- The Go language compiler is non-trivial to write (and read). Looking at the
  code, I don't understand how it works.
- Go language has garbage collection built-in in the runtime, so it is probably
  not the best language for writing an operating system.

So, here is Xlang, my own toy language as an effort to create human
comprehensible code at the very foundation of computer software systems:
compilers and operating systems.

I am no way a very good compiler writer, like I have never written a real
compiler IR optimizer before. So it won't be a very good compiler with
good codegen results. Really could need some help if you know more on
how to implement and optimize SSA's.

If you are interested (and want to help), please contact me: liulonnie@gmail.com

# *meltedbrain* - A Brainf*ck Interpreter made in Go

> This project is inspired by this [article](https://thorstenball.com/blog/2017/01/04/a-virtual-brainfuck-machine-in-go/) written by Throsten Ball.

## Overview
*meltedbrain* is a fast and simple interpreter for Brainf*ck programming language written in Go. This project is not meant to be used in production since it's just a learning project.

You can learn more about the Brainf*ck programming language here:
- [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck).
- [Video by Fireship](https://www.youtube.com/watch?v=hdHjjBS4cs8).

## Install

You can install *meltedbrain* with the `go install` command or build the binary yourself.

````bash
go install github.com/daniarlert/meltedbrain@latest
````

## Usage

If you just want to run a program written in Brainf*ck use:
````bash
meltedbrain "./examples/hello_world.bf"
````

If instead what you want is to see the set of instruction that the compiler generates use:
````bash
meltedbrain --d "./examples/hello_world.bf"
````
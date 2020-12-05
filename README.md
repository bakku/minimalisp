# Minimalisp

After reading [Crafting Interpreters](www.craftinginterpreters.com) Minimalisp is my small take
on writing a programming language from scratch. It mimicks the approach of the *jlox* implementation
and uses a tree-walking approach to interpret source code.

## Compiling

To compile Minimalisp just execute:

```
make
```

Golang must be installed, of course. This will create an `mlisp` executable in the `cmd/mlisp` folder.

## Usage

Minimalisp can be used either as a REPL or for running simple scripts. Here is an example of Minimalisp code:

```clojure
(defvar names '(1 2 3 4))
(defun quarter-of (n)
  (let (half (/ n 2))
    (/ half 2)))
  
(let (quarters (map quarter-of names))
  (if (= 0.25 (first quarters))
    (println "yes!")
    (println "no!")))
```

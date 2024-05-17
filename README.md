# Chui Compiler

> Cause it blazingly fast 

- A hobbyist functional programming language and interpreter project done in Golang as a way of understanding Golang.

- The goal of this project is to turn a `tree-walking, on-the-fly evaluating interpreter` into a `bytecode compiler` and a `virtual machine` that executes the bytecode.

- To mimic `Ruby`, `Lua`, `Python`, `JavaScript` implementations and `Perl`, even the mighty `Java Virtual Machine`. Bytecode compilers and virtual machines are everywhere - with good reason.

- The reason is that to provide a new layer of abstraction - the bytecode passed from the compiler to the virtual machine - makes the system more modular, the main appeal of this architecture lies in the performance. Bytecode interpreters are way faster than tree-walking interpreters!

Here's is a demo:

```shell
    $ ./chui-fibonacci -engine=eval
    engine=eval, result=9227465, duration=27.204277379s
    $ ./chui-fibonacci -engine=vm 
    engine=vm, result=9227465, duration=8.876222455s
```

- Yeap you read that right! 3 times faster without low-level tweaking or mind blowing optimizations.

- It contains:
    - [x] C-ish syntax
    - [x] Prefix-, infix- and index operators
    - [x] Conditionals
    - [x] Variable Bindings (Global and Local)
    - [x] Integers, Booleans and Arithmetic Expressions
    - [x] Built-In functions

    - [x] First-Class and Higher-Order Functions**
    - [x] Closures**
    - [] Macros*** (Elixir like)

    - [x] String Data Structure
    - [x] An Array Data Structure
    - [x] A Hash Data Structure

- How does it look?

```javascript
   let number = 5;
   let name = "Chui";
   let isTrue = true;

   // Arrays in Chui
   let numArray = [1, 2, 3, 4, 5];
    
   // Hashes in Chui
   let cyrus = {"name": "Cyrus", "age": 26};

   // Accessing elements in arrays
   numArray[0]; // 1

   // Accessing property values in hashes
   cyrus("name"); // Cyrus
```

- The `let` keyword is used to declare variables in Chui, above is how we bind variables to values, and some Data Structures. `let` can also be used to bind function names.

```javascript
    // Implicit returns are supported out of the box
    let add = func(x, y) {
        x + y;
    };

    let add = func(x, y) {
        return x + y;
    };

    // Calling the function
    add(5, 5); // 10
```

- We can also write complex functions in Chui, that use recursion:

```javascript
    let fibonacci = func(x) {
        if (x == 0) {
            return 0;
        } else {
            if (x == 1) {
                return 1;
            } else {
                fibonacci(x - 1) + fibonacci(x - 2);
            }
        }
    };

    fibonacci(10); // 55
```

- We are also working on Higher-Order functions. `These are functions that take other functions as arguments:

```javascript
    let applyFunc = func(f, x) {
        return f(f(x));
    };

    let addOne = func(x) {
        x + 1;
    };

    applyFunc(addOne, 5); // 6

    let map = func(arr, f) {
        let iter = func(arr, accumulated) {
            if (len(arr) == 0) {
                accumulated
            } else {
                iter(rest(arr), push(accumulated, f(first(arr))));
            }
        };
        5
        iter(arr, []);
    };

    let numbers = [1, 1 + 1, 4 - 1, 2 * 2, 2 + 3, 12 / 2];

    map(numbers, fibonacci); // => returns: [1, 1, 2, 3, 5, 8]
```

- From the above demo we can see applyFunc takes two arguments: a function `f` and a value `x`. It calls addOne 5 times with first 2 as argument and returns the result, 6.

- Why Golang? Well why not? Golang is a statically typed language, and it's fast. It's also a language that I have been wanting to learn for a while now. So why not kill two birds with one stone? Also Go has garbage collection, which means I don't have to worry about memory management (wink* memory leaks wink* :D)

- The project is still in its early stages, and I am still learning Golang. So if you have any suggestions, or you want to contribute, feel free to reach out to me. I am always open to learning new things.
# Chui Compiler

> 📝 **Note:** Cause it blazingly fast

- A hobbyist functional programming language and interpreter project done in Golang as a way of understanding Golang.

## Introduction
- The goal of this project is to turn a `tree-walking, on-the-fly evaluating interpreter` into a `bytecode compiler` and a `virtual machine` that executes the bytecode.

- To mimic `Ruby`, `Lua`, `Python`, `JavaScript` implementations and `Perl`, even the mighty `Java Virtual Machine`. Bytecode compilers and virtual machines are everywhere - with good reason.

- The reason is that to provide a new layer of abstraction - the bytecode passed from the compiler to the virtual machine - makes the system more modular, the main appeal of this architecture lies in the performance. Bytecode interpreters are way faster than tree-walking interpreters!

- Here's is a demo:

```shell
    $ ./chui-fibonacci -engine=eval
    engine=eval, result=9227465, duration=27.204277379s
    $ ./chui-fibonacci -engine=vm
    engine=vm, result=9227465, duration=8.876222455s
```

- Yeap you read that right! 3 times faster without low-level tweaking or mind blowing optimizations.

- To run the code above, you need to have Golang installed on your machine. You can download it [here](https://golang.org/dl/).

- Open you terminal/bash:

```shell
    # Clone the repo
    $ git clone git@github.com:Cyrus-0101/chui.git

    # Navigate to the src folder
    $ cd chui/src

    # Run the main application
    $ go run main.go

    # Build the application
    $ go build -o build main.go

    # Run the executable application if on Windows
    $ ./build/main.exe

    # Run the build on Linux/Mac OS
    $ ./build/main

    # Generate a coverage report
    $ go tool cover -html cover.out -o cover.html

    # Run the tests with a small terminal coverage report
    $ go test -v -coverprofile cover.out ./...
```

- It contains:
    - [x] C-ish syntax
    - [x] Prefix-, infix- and index operators
    - [x] Conditionals
    - [x] Variable Bindings (Global and Local)
    - [x] Integers, Booleans and Arithmetic Expressions
    - [x] Built-In functions

    - [x] First-Class and Higher-Order Functions**
    - [x] Closures**
    - [x] Macros*** (Elixir like)

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
let cyrus = { name: "Cyrus", age: 26 };

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

## Macros

- Macro systems can be defined as features of a language that concern itself with transforming code before execution- macros; that is how to define them, how to access them, how to evaluate them and how macros work.

- They can be divided into two broad categories:

### 1. Text-substitution macro systems (search-and-replace).

- Arguably simpler, and a good example is the C preprocessor. It allows you to generate and modify C code by using a seperate macro language in the rest of your normal C code.
- How it works is by parsing and evaluating the separate language before the actual C code is compiled. This is done by the preprocessor, which is a separate program that runs before the actual compiler.

```c
    #define GREETING "Hello there"

    int main(int argc, char *argv[])
    {
    #ifdef DEBUG
        printf(GREETING " Debug-Mode!\n");
    #else
        printf(GREETING " Production-Mode!\n");
    #endif

        return 0;
    }
```

- Instructions to the preprocessor are given in the `#` symbol. This means `GREETING` will be replaced with `Hello there` before the actual C code is compiled.

- In the 5th line, we check for the `DEBUG` variable, if present in the C libs or predefined, its presence defines either the Debug-Mode or Production-Mode statements will be printed.

- However, its an efficient system if used with care and restraint. It's limited, since code production is based on a textual level. In that regard, its closer to a templating system than a macro system.

### 2. Syntactic macro systems (code-as-data camps).

- They treat code as data, yes that's weird but its true. Think of how lexers and parsers turn source code from text to (Abstract Syntax Trees) ASTs.
- In our case, we turned the Chui source code, which was a string, into the structs in Go that make up our Chui AST;Then we could treat the code as data: we could pass around, modify and generate Monkey source code inside our Go program.
- Now languages with this type of macros can do it within the language itself and not just in an outerhost language. If a language has a syntactic macro system, you can use language X to work with source code in language X.
- This kind of makes the language more self aware; a step closer to AGI![Thanos Infinity Stones](https://media1.tenor.com/m/qUdBKJFbXzEAAAAd/thanos-avengers-infinity-war.gif)

- No, but seriously, macros allow you to inspect and modify code, kind of like how actors operate on themselves in movies.![Rambo 3 (1988) - Gunpowder Cauterization Scene (1080p) FULL HD](https://i.makeagif.com/media/6-21-2021/W_4mSa.gif)

- Ok enough gifs :D, this type of macro system was popularized and pioneered by Lisp, and is found in its descendants such as Clojurem Scheme, Racket and even non-Lisp languages such as Julia and Elixir.

- Lets see how an example looks like in Elixir, yeah it kind of looks like Ruby and Python but lets check it out. Elixir's `quote` function allows use to stop code from being evaluated - effectively converting code into data.

```elixir
    iex(1)> quote do: 10 + 5

    {:+, [context: Elixir, import: Kernel], [10, 5]}
```

- Above we pass the infix expression 10 + 5 to`quote` as a single argument block in a `do` block.
- But instead of 10 + 5 being evaluated - as args in normal function calls - `quote` returns a data structure that represents the expression, a `tuple` containing the operator `:+`, meta information like the context of the call and a list of the operands `[10, 5]`.
- This is Elixir's AST, and how code is represented all through [Elixir](https://hexassets.pm/elixir/introduction.html)

- The most important thing about Macros in Elixir is: everything is passed to macro as an argument is `quote` d. This means the macro's arguments are not evaluated and can be accessed like any other piece of data.

- This is what we will be working on to implement in Chui. It's not as easy as let's get to it, so many questions come along, "How? Why? What are the implications? Is this for clout?" Let's get a mental picture of what we expect to build:

- We will model a macro system after Elixir's, which itself is modelled after a simple define-macro system popular in the Lisp and Scheme world.

- We will definitely add the quote and unquote functions; which allow us to influence when exactly Chui code is evaluated. Here is a small illustration:

```shell
    $ go run main.go
    Hello DESKTOP-51RHGR7\cyrus! This is the Chui programming language!
    Feel free to type in commands
    >> quote(foobar);
    QUOTE(foobar)
    >> quote(10 + 5);
    QUOTE((10 + 5))
    >> quote(foobar + 10 + 5 + barfoo);
    QUOTE((((foobar + 10) + 5) + barfoo))
```

- Hear me out, `quote` will take one argument and stop it from being evaluated. It will return an object that represents the quoted code.
- The matching `unquote` code will allow us to circumvent `quote`:

```shell
    >> quote(8 + unquote(4 + 4));
    QUOTE((8 + 8))
```

- `unquote` will only be usable inside the expression that's passed to `quote`. But in there it will also be possible to `unquote` source code that's been quoted before:

```shell
    >> let quotedInfixExpression = quote(4 + 4);
    >> quotedInfixExpression;
    QUOTE((4 + 4))
    >> quote(unquote(4 + 4) + unquote(quotedInfixExpression));
    QUOTE((8 + (4 + 4)))
```

- This will be important for when we put in the final system, the `macro` literals which allow us to define macros in Chui:

```shell
    >> let reverse = macro(a, b) { quote(unquote(b) - unquote(a)); };
    >> reverse(2 + 2, 10 - 5);
    1
```

- Similar to functions they look like function literals, and once a macro is bound to a name we can call it in our code. Similar to `Elixir's` macros, the arguments to the macro are not evaluated before the macro is called. This is paired with the use of `quote` d and `unquote` to selectively evaluate macro arguments, which are just `quote` d code:

```shell
    >> let evalSecondArg = macro(a, b) { quote(unquote(b)) };
    >> evalSecondArg(puts("not printed"), puts("printed"));
    printed
```

- From the above illustration we can see that by returning code that only contains the second argument, the `puts("printed")` expression, the first argument is never evaluated.

- There are some challenges faced/trade-offs made: This is not production ready, not near what I've seen in open source repos, however as a first attempt, it will be a fully working macro system. Let's document code that writes code:

<center><video width="320" height="240" controls src="./assets/lfg.mp4"/></center>

#### Quote

- Let's define the quote function, it will only be used inside macros and its purpose is; when called it stops arguments from being evaluated and returns the AST node representing the arg.

- Every function in Chui returns values of the interface type `object.Object`, this is because our `Eval` function relies on Chui value being an `object.Object` to work.

- This means in order for quote to return an `ast.Node`, we need a simple wrapper that allows us to pass around `object.Object` containing an `ast.Node`:

<!-- TO-DO: Add unquote notes -->

## Compilers

- Below is a resource of a compiler demonstrating how code is mysteriously converted to machine code, no mere mortal can comprehend.

    ![Chui](/assets/compiler.png)

- Virtual machines are also largely misunderstood mystical virtual creatures. Some say they have something to do with compilers as shown above. Others claim that virtual machines allow them to run an operating system within an operating system. The truth is, all of that is true!

- Wikipedia defines a compiler as:

  > A compiler is computer software that transforms computer code written in one programming language (the source language) into another computer language (the target language). Compilers are a type of translator that support digital devices, primarily computers. The name compiler is primarily used for programs that translate source code from a high-level programming language to a lower level language(e g., assembly language, object code, or machine code) to create an executable program.

- Compilers are fundamentally about translation, because translation is how they implement a programming language.

- The above might sound counter-intuitive, that interpreters and compilers are opposites, but while their approach is different, they share a lot of things in their construction. They both have a frontend that reads in source code in the source language, and turns it into a data structure.

- In both, compiler and interpreter, the frontend is usually made up of a lexer and a parser, that generate a syntax tree. In the frontend they have similarities, after that when they both traverse the AST, their paths diverge.
- Let's take a look at the lifecycle of code being translated to machine code below:
  ![Compiler Lifecycle](/assets/compiler-lifecycle.png)

1. The source code is tokenized and parsed by the lexer and parser respectively. This is the frontend. The source code is turned from text to AST.

2. Then the component called optimizer (sometimes called the compiler), might translate the AST into another Internal Representation (IR). This additional IR might be another syntax tree, or maybe a binary format, or even a textual format. The main reason for this additional translation to another IR, is that the IR might lend itself better optimizations and translations to the target language.

3. This new IR then goes through an optimization phase: where dead code is eliminated, simple arithmetic is pre-calculated, ... ,a ton of other optimizations.

4. Finally the code generator, also called the backend, generated the code in the target language. This is where the compilation happens. The code hits the file system.

- The above is a very high level overview of the lifecycle of a compiler. There are thousands of variations possible. And then again, a compiler doesn’t even have to be a tool you run on the command line, that reads in source code and outputs code in a file, like gcc or go. It can just as well be a single function that takes in an AST and returns a string. That’s also a compiler. A compiler can be written in a few hundred lines of code or have millions of them. But underlying all of these lines of code is the fundamental idea of translation.

### Virtual Machines & Real Machines.

- We probably associate “virtual machine” with software like VMWare or Virtualbox. These are programs that emulate a computer, including a disk drive, hard drive, graphics card, etc. They allow you to, for example, run a different operating system inside this emulated computer. Yes, these are virtual machines.

- We will discuss (and later build), virtual machines used to implement programming languages. It's worth noting `VM's (in this instance) do not emulate an existing machine; They are the machine`. Virtual simply means software. [Here is a perfect example of a simple VM written in JS.](/assets/virtualMachine.js)

- The machine describes their behaviour, these software constructs act and mimic the behaviour of hwo a real machine works, computers. This means we need to understand how computers work in order to understand how virtual machines work.

### Real Machines

- Nearly every computer in our lives is built according to the Von Neumann Architecture, which describes a way to build a fully-functioning computer with a surprisingly tiny number of parts.

- In Von Neumann's model, a computer has twi central parts:

  1. A Processing Unit: Which contains an Arithmetic Logic Unit (ALU) and multiple process registers, and a control unit, with an instruction register and a program counter. Together they're known as the Central Processing Unit (CPU).

  2. A Memory: This is where the CPU reads and writes data. The memory is a long list of slots, each of which can hold a number. The CPU can read from memory, write to memory, and perform arithmetic on the numbers in memory. Such include: `Read Only Memory (ROM)`, `Random Access Memory (RAM)`, `Hard Drives`, etc.

- Below is a diagram of the Von Neumann Architecture:

  ![Von Neumann Architecture](/assets/von-neumann-architecture.png)

## Bytecode

- Here we simply want to compile and execute the following Chui expression:

> 2 + 2

- We will discuss how to:
1. Take the expression 2 + 2
1. Tokenize and Parse using our existing `lexer`,`token` and `parser` packages.
1. Take the resulting AST - nodes defined in `ast` package.
1. Pass it to the new compiler, which compiles the AST into bytecode.
1. Pass the bytecode to the new VM, which executes it.
1. Make sure the VM turned it into `4`.

- The above expression will travel through all the major parts of our new interpreter:

    ![Interpreter Flow](/assets/interpreter-flow.png)

- Lets see how the flow of data structures will look like:

    ![Data Structures Flow](/assets/data-structure-flow.png)

- From the above we can that we can handle the flow up till the AST; then we begin the hunger games of the compiler and VM, i.e define bytecode instruction, build the compiler and the VM.

- The architecture of the virtual machine is the biggest influence on what the bytecode will look like, this means a decision has to be made on the direcion of the build.

- We will build a stack-based VM, which is the most common type of VM. This means that the VM will have a stack, which is a data structure that allows you to push and pop values. The stack is the only place where the VM can store values. Less complex compared to register machine, more moving parts, more concepts.

- We will begin by gettinf the operands `2` and `2` onto the stack and tell the VM `add them`, which takes the two topmost values from the stack, adds them, and pushes the result back onto the stack. Here is a before and after of the stack:

    |  2  |  4  |
    | --  | --  |
    |  2  |     |
    |     |     |
    Before| After

- The instructions to our VM will be:
1. Push 2 to the stack
1. Push 2 to the stack
1. Add the two topmost values on the stack.

- We can see we need to implement two instruction types in total: One for pushing to the stack and another for adding values in the stack.

- Let's define the opcodes and how they are encoded in bytecode, then extend the compiler to generate instructions, then create a VM that decodes and executes the instructions. We'll create a new package `code` to define the bytecode instructions and the compiler.
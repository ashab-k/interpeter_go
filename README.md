# Simian Programming Language Interpreter

Simian is a lightweight and expressive custom programming language with a simple syntax designed for learning and experimentation. This interpreter, written in Go, demonstrates the implementation of a basic yet fully functional programming environment.

## Features

- **Dynamic Typing:** Support for integers, booleans, and functions.  
- **Control Structures:** Includes conditionals (`if/else`) and loops.  
- **Function Definitions:** First-class functions with lexical scoping.  
- **Arithmetic Operations:** Support for basic mathematical operations (`+`, `-`, `*`, `/`).  
- **Variable Bindings:** Flexible variable declarations.  
- **REPL (Read-Eval-Print Loop):** An interactive shell for immediate evaluation.

## Getting Started

### Installation

1. Clone the repository:

   ```bash
   git clone <your-repo-url>
   cd simian-lang
   ```

2. Run the interpreter:

   ```bash
   go run main.go
   ```

3. Access the REPL (Read-Eval-Print Loop) for immediate code execution.

### Usage

Once inside the REPL, you can execute Simian code directly:

```simian
let x = 5 * 3;  
x + 2;  
if (x > 10) { x - 1 } else { x + 1 };  
```

To exit the REPL, use `Ctrl+C` or type `exit`.

## Example Programs

### Fibonacci Sequence

```simian
let fib = fn(n) {  
  if (n < 2) {  
    n  
  } else {  
    fib(n - 1) + fib(n - 2);  
  }  
};  

fib(10);  
```

### Factorial

```simian
let factorial = fn(n) {  
  if (n == 0) {  
    1  
  } else {  
    n * factorial(n - 1);  
  }  
};  

factorial(5);  
```

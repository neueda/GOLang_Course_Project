## Introduction

When transitioning from one programming language to another, understanding the syntax similarities and differences is crucial for a smooth learning experience. This wiki page aims to provide a comparative analysis of the syntax of two popular languages: Java and GoLang (Go). By comparing their syntax, we aim to help Java developers grasp the key concepts and bridge the gap between these languages.

## Braces and Code Blocks

Both Java and GoLang use curly braces ({}) to define code blocks. This shared syntax promotes code readability and maintains consistency between the two languages. However, there are some subtle differences to be aware of:

### Java


    if (condition) {     
        // Code block 
    } else 
    {    
        // Code block 
    }

### GoLang

    if condition {
        // Code block 
    }else { 
        // Code block 
    }

In GoLang, the opening brace ({) should be placed on the same line as the statement or condition, while in Java, it is often placed on the next line.

## Variable Declarations

Variable declaration syntax varies between Java and GoLang. Let's compare how variables are declared in each language:

### Java

    int age = 25; String name = "John";

### GoLang

    var age int = 25 name := "John"

In Java, you explicitly declare the variable type before the variable name, whereas in GoLang, the type comes after the name. GoLang also supports short variable declarations using the := operator, which infers the type based on the assigned value.

## Control Flow Statements

Java and GoLang share similar control flow statements like if-else, for loops, and switch statements. Let's compare their syntax:

### Java

    if (condition) {    
        // Code block 
    } else if (condition) 
    {     
        // Code block 
    } else {     
        // Code block 
    }  
    
    for (int i = 0; i < 5; i++) 
    {     
        // Code block 
    }  
    switch (value) { 
        case 1:     
        // Code block 
            break;     
        case 2: 
            // Code block
            break;
        default:
            // Code block 
    }

### GoLang

    if condition {
        // Code block 
    } else if condition {    
        // Code block 
    } else {     
        // Code block 
    }  

    for i := 0; i < 5; i++ {     
        // Code block 
    }  switch value {  
        case 1:        
        // Code block     
        case 2:        
        // Code block    
        default:         
        // Code block 
    }

The syntax of control flow statements is quite similar in both languages, with the primary difference being the placement of parentheses in GoLang.

## Exception Handling

Exception handling is an essential aspect of programming. Let's examine the syntax differences in exception handling between Java and GoLang:

### Java
    try {     
        // Code block 
    } catch (Exception e) {     
        // Exception handling code
    } finally {    
        // Code block 
    }

### GoLang

GoLang does not have built-in exception handling like Java. Instead, GoLang adopts a unique approach that promotes explicit error handling through return values and the error type.

## Pointers

Pointers are variables that hold the machine address of some other value. References are variables that hold a locator (could be an address or something else) for another value.

### Java

Java uses the (safer) idea of references instead of pointers. The Java language does not provide pointers. Instead, all objects are handled by references.

### GoLang

Go has pointers. A pointer holds the memory address of a value.

The * operator denotes the pointer's underlying value. The type *T is a pointer to a T value. Its zero value is nil.

    var p *int
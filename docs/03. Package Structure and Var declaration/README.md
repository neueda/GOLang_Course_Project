## Introduction

When transitioning from Java to GoLang (Go), understanding the package structure and variable declaration syntax is crucial for a smooth onboarding experience. This wiki page aims to provide a comprehensive guide for Java developers, explaining the package structure in GoLang and the differences in variable declaration syntax between the two languages. By familiarizing yourself with these concepts, you'll be well-equipped to start working with GoLang confidently.

## Package Structure

In GoLang, every file belongs to a package, and packages are organized within a directory structure. Understanding the package structure is important for organizing and managing your GoLang projects effectively. Here are some key points to note:

- A GoLang package is defined using the ```package``` keyword followed by the package name. For example, ```package main``` is commonly used for executable programs.
- The package name should be meaningful and descriptive, indicating the purpose or functionality of the package.
- The package name should be lowercase, following the convention of lowercase package names in GoLang.
- Packages are organized within a directory hierarchy. By default, the name of the directory containing the GoLang source code should match the package name.
- GoLang provides a standard directory structure, where different packages reside in separate directories within the project.

For example, in a project with the following structure:

### Go

    myproject/ 
    ├── main.go 
    ├── utils/ 
    │   ├── util.go 
        │└── helper.go 
    └── api/  
        ├── handler.go
        └── routes.go

The ```main.go``` file belongs to the ```main``` package, the files within the ```utils``` directory belong to the ```utils``` package, and the files within the ```api``` directory belong to the ```api``` package.

## Variable Declaration

Variable declaration syntax in GoLang differs from Java and requires understanding the key differences. Here are some important points to note:

- GoLang uses the ```var``` keyword to declare variables explicitly, followed by the variable name and type. For example, ```var age int```.
- Unlike Java, the type in GoLang comes after the variable name.
- GoLang also supports short variable declarations using the ```:=``` operator. This syntax allows GoLang to infer the variable type based on the assigned value. For example, ```name := "John"``` is equivalent to ```var name string = "John"``` in Java.

It's important to note that GoLang encourages explicit type declarations to enhance code readability and maintainability. However, the short variable declaration syntax can be used for concise variable declarations when the type is obvious from the assigned value.

By understanding the differences in variable declaration syntax between Java and GoLang, Java developers can adapt their coding style to leverage GoLang's explicit and concise variable declaration options effectively.
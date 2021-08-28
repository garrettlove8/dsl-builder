# Overview
DSL Builder is a library for building DSLs written in C. While this is above all else meant as an interesting and fun project for me to learn C, it is also for enabling anyone to easily inlcude a DSL in their project.

**Note:** This project has been architected using the [Bullseye Model](https://github.com/garrettlove8/bullseye-model)

# 1. Project Scope
When it is all said and done, DSL Builder will provide the ability to build a full-featured DSL from scratch.

The goal of version one is to provide a way for implementing a semantic model through a language's lexicon.

In other words, when a user of a DSL type in a command, we need to answer a couple questions.

1. How is the command understood?
2. What is the command supposed to do?

These are the questions a semantic model will answer.

## Use Cases
Note: For my own understanding this **would not** be used to build a CLI, though perhaps it could. For that, their are other tools available (which are likely better anyways).

### Query Language
Let's say you are **building** a database and you'd like to access some data within it. Notice that I wrote "building", since in this case you are already in the database, you have complete control of the data. You could simply open the file where the data is stored, allocate some memory from the heap, store the data in that allocated memory block and manipulate it by referring directly to its address, using pointer arithmetic to move throughout it.

While this might, in fact, end up being what you'd do anyways, if you intended for other developers to be able to use your new shiny database as part of their applications you'd need to offer some way of externally interacting with your database. This is where a query language comes into play, which is simply an [external DSL](https://martinfowler.com/dsl.html).

### Rules Engine
Another potential use-case would be a rules engine. In this situation you essentially need to describe what should happen based on some contextual data.

# 2. Define Packages


# 3. Package Scope and Context

# 4. UML Diagrams

# 5. API Definitions

# 6. Process Flow

# 7. Data Flow
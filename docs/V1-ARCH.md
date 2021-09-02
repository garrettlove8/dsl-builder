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

## Command
A command is the full sequence of words and symbols provided by a client of the program.

### Context
A command is essentially the entry to using a DSL built using DSL Builder. Without a command, there is nothing to be parsed and acted on.

### Scope
The command package only cares about the initial command input from a client, as well as tokenizing that command.

Note the limitations of the tokenization process here. The `command` package should not know about the token package. All it has to do is take a command and divide it into the appropriate linguistic units.

### API
When using the command package, clients should be able to:

1. `Sequence` a command
	- Ie. ganarate a sequence of tokens from the command
	- `It might make sense to call this something else here since there is also the concept of a "token" in the form of a separate package`

## Token
A token is a lexical unit within the overall command that can be acted on.

```c
struct Token {
	// lexicalUnit is the the character sequence representation of a single lexical unit.
	// Array of char is used to allow a multi-character squences.
	// The number 13 is used because that's the average meaningful length word in English.
	char lexicalUnit[13];

	// args are additional configuration and context that can be provided to the lexicalUnit for consideration.
	// Array of char is used to allow for multi-character sequences.
	char args[100];

	// fn is a behavior function on the lexicalUnit
	// It accepts a pointer to the object instance and the args for its implementation details.
	// A pointer to the object instance is required to allow the function to modify the object in place.
	// This provides the implicit object encapsulation.
	fn (*fn)(Token*, args, ...);

	// parent is a pointer to this specific Token's parent token.
	// Together, with the child attribute, parent makes Token is a doubly-linked list.
	struct Token* parent;

	// child is a pointer to this specific Token's child token.
	// Together, with the parent attribute, child makes Token is a doubly-linked list.
	struct Token* child;
}
```

```
TODO: NOTE: I think starting to diagram the packages would help me keep track of what I'm typing up and structure everything. It would also give a head start on UML parts.
```

### Context
A command is made up of one or more single token. Furthuremore, a token tree is made up of one or more single tokens.

Generating tokens (via tokenization) is necessary for bridging the concept of a `command` and the concept of a `token tree`. Thus, without tokens no DSL can be successfully created and run.

### Scope
A token does nothing more than provide the base encapsulation for a command. An instance of a token object holds the necessary properties to declare the intention and facilitate the execution of a lexical unit. A token does not know, or care, about the implementation of an action to be performed by a DSL. Is is strictly a way of **declaring the intention and facilitate the execution of a lexical unit**.

### API

## Token Tree
### Context
### Scope
### API

## Token Options
The valid lexical units that are available in the DSL being built. For example, in SQL `SELECT`, `FROM`, and `WHERE` are all lexical units.
### Context
### Scope
### API


# 4. UML Diagrams

# 5. API Definitions

# 6. Process Flow

# 7. Data Flow
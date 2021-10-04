#ifndef SCRIPT_H
#define SCRIPT_H

// script_t represents a complete DSL script and provides an API for running a script.
typedef struct {
	// full is a pointer which points to the complete script.
	char* full;

	// sep is a character array which holds the character on which a developer intends
	// for statements to be terminated.
	char sep[2];
} script_t;

// run allows a developer to execute a DSL script.
void* run(script_t*);

#endif
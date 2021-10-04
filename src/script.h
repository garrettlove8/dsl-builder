#ifndef SCRIPT_H
#define SCRIPT_H

// script_t
typedef struct {
	char* full;
	char sep[2];
} script_t;

void* run(script_t*);

#endif
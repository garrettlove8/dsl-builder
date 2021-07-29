#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>

// TODO: Questions:
// 1. I'm currently just writing code, knowing that I'll have to refactor later, How should this be refactored to fit the strucutre of an idiomatic C program?
// 2. How do you use third-party packages (or are they called libraries in C)?
// 3. How do you create your on library in C?

/*
	--- Sample Input ---
	glove init c -name first
	glove init go -name first
*/

/*
	1. Create a map to be used as the root
	2. Tokenize args
	3. Create tree structure with args
		- command structure
			{ 
				cmd: string // command name as key
				child: cmd
				fn: function -> cmd
				...args
				name: string
				...args
			}
	4. Follow tree and executute functions that correspond to the commands
*/

typedef struct cmdNode {
	char *cmd;
	struct cmdNode *child;
	struct cmdNode (*fn)();
	char name;
} cmdNode;

cmdNode *rootCmd;
char *tokens[];

int createNode(char *text) {
	printf("text: %s\n", text);

	cmdNode *cmd = malloc(sizeof(cmdNode));

	cmd->cmd = text;

	printf("cmd: %s\n", cmd->cmd);

	return EXIT_SUCCESS;
}

int main(int argc, char *argv[]) {
	// Create the tree's root command
	rootCmd = malloc(sizeof(cmdNode));

	// for(int a = 0; a <= ) {

	// }

	int parseResult = parseArgs(argv[1]);
	
	return EXIT_SUCCESS;
}
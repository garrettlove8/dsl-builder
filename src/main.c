#include <stdio.h>
// #include <stdarg.h>

// int myPrintf(const char *format, ...) {
// 	va_list arg;
// 	int done;
 
//   	va_start (arg, format);
// 	done = vfprintf (stdout, format, arg);
// 	va_end (arg);

// 	return done;

	// printf("stdin: %s", "myPrintf");
	// return 0;
// }

int printConsole(char text) {
	putc(text, stdout);
	return 0;
}

int main(void) {
	// puts("Hello, this is the CLI");
	// printf("stdin: %s", "test input");
	// myPrintf("stdin: %s", "test input");
	printConsole(71);
	return 0;
}
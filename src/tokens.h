// token_t is extended by each mdoel in the user's program.
// token_t provides a default behavior function which the
// then overrides to provide the desired funcionality of a
// given token, which token_t calls on their behalf instead
// if its default behavior function.
typedef struct {
	char lexicalUnit;
} token_t;

typedef struct {
	int (*funcOne) (void*, ...);
	int (*funcTwo) (void* objPtr);
	int (*funcThree) (void* objPtr);
} symbols_t;
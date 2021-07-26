init:
	mkdir src
	mkdir docs
	mkdir bin
	mkdir .tmp
	touch README.md
	touch CHANGELOG.md
	touch src/main.c

make build:
	cc -o bin/glove src/main.c

make run:
	@printf '%s\n' "--- Building ---"
	cc -o .tmp/glove src/main.c

	@printf '\n%s\n' "--- Running ---"
	.tmp/glove

	@printf '\n%s\n' "--- Cleaning ---"
	rm .tmp/glove
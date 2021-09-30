init:
	mkdir src
	mkdir docs
	mkdir bin
	mkdir .tmp
	touch README.md
	touch CHANGELOG.md
	touch src/main.c

build:
	gcc -o bin/$(EXEC) src/*.c

run:
	@printf '%s\n' "--- Building ---"
	gcc -o .tmp/$(EXEC) src/*.c

	@printf '\n%s\n' "--- Running ---"
	.tmp/$(EXEC)

	@printf '\n%s\n' "--- Cleaning ---"
	rm .tmp/$(EXEC)
init:
	mkdir src
	mkdir docs
	mkdir bin
	mkdir .tmp
	touch README.md
	touch CHANGELOG.md
	touch src/main.c

make build:
	gcc -o bin/$(EXEC) src/main.c

make run:
	@printf '%s\n' "--- Building ---"
	gcc -o .tmp/$(EXEC) src/main.c

	@printf '\n%s\n' "--- Running ---"
	.tmp/$(EXEC)

	@printf '\n%s\n' "--- Cleaning ---"
	rm .tmp/$(EXEC)
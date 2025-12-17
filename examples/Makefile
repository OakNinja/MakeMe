.PHONY: all clean test

all: program

program: main.o utils.o
	@echo "Linking..."
	@cc -o program main.o utils.o
	@echo "Done."

main.o: main.c
	@echo "Compiling main.c..."
	@cc -c main.c

utils.o: utils.c
	@echo "Compiling utils.c..."
	@cc -c utils.c

clean:
	@echo "Cleaning up..."
	@rm -f program *.o main.c utils.c

test:
	@echo "Running tests..."

main.c:
	@echo "int main() { return 0; }" > main.c

utils.c:
	@echo "void utils() {}" > utils.c

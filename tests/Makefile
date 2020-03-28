blah: blah.o
	cc blah.o -o blah

blah.o: blah.c
	cc -c blah.c -o blah.o

blah.c:
	echo "int main() { return 0; }" > blah.c

.PHONY: foo
foo:
	rm -f some_file
	rm -f clean
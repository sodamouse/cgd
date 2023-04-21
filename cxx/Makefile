CC = g++
CFLAGS = -Wall -Wextra -Wpedantic
INPUTS = cgd.cpp girlboss.cpp
OUTPUT = cgd
LIBS = fmt

.PHONY: all

all:
	$(CC) $(CFLAGS) -o $(OUTPUT) -l$(LIBS) $(INPUTS)

install:
	cp $(OUTPUT) /usr/local/bin

uninstall:
	rm /usr/local/bin/$(OUTPUT)

clean:
	rm $(OUTPUT)

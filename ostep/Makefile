CC=gcc
CCOPTS=#-Wall -Wextra
LIBS=

SRCS=$(wildcard *.c)
TARGETS=$(SRCS:.c=)

.PHONY: all clean new

all: $(TARGETS)

clean:
	rm -f *.bin

new: clean all

%: %.c
	$(CC) $(CCOPTS) -o $@.bin $< $(LIBS)
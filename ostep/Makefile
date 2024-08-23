CC=gcc
CCOPTS=-Wall -Wextra
LIBS=

SRCS=$(wildcard *.c)
TARGETS=$(SRCS:.c=)

.PHONY: all clean pristine

all: $(TARGETS)

clean:
	rm -f $(TARGETS) & rm -f *.output

pristine: clean all

%: %.c
	$(CC) $(CCOPTS) -o $@ $< $(LIBS)
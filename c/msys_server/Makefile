# Define the compiler
CC = gcc

# Define the flags
CFLAGS = -Wall -Wextra

# Define the library flags
LDFLAGS = -lws2_32

# Define the target executable
TARGET = server

# Define the source file
SRC = webserver.c

# Default target
all: $(TARGET)

# Build step
$(TARGET): $(SRC)
	$(CC) $(CFLAGS) -o $(TARGET) $(SRC) $(LDFLAGS)

# Run the server
run: $(TARGET)
	./$(TARGET)

# Clean step to remove the executable
clean:
	rm -f $(TARGET)

.PHONY: all run clean

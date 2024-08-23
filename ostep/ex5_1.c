#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <unistd.h>

int main(int argc, char* argv[]) {
    int* x = (int*)malloc(sizeof(int));

    int rc = fork();

    if (rc == 0) {
        for (int i = 0; i <= 10; i++) {
            printf("child | x=%d\n", (*x)++);
            usleep(1);
        }

    } else {
        for (int i = 0; i <= 10; i++) {
            printf("parent| x=%d\n", (*x)++);
            usleep(1);
        }

        wait(NULL);
    }

    return 0;
}

/*
Q: Write a program that calls fork(). Before calling fork(), have the
main process access a variable (e.g., x) and set its value to something (e.g.,
100). What value is the variable in the child process? What happens to the
variable when both the child and parent change the value of x?

A:
No race condition happens because they don't share memory actually.
Child process gets its own "private memory" (copy of parent's address space)
Therefore parent & child modify their own copy of x, indepedently.

Things would break if they were using shared memory.

note: Even if x is a pointer, nothing changes!
*/
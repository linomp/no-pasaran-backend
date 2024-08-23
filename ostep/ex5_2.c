#include <assert.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

int main(int argc, char* argv[]) {
    int fd = open("./ex5_2.output", O_CREAT | O_WRONLY | O_TRUNC, S_IRWXU);
    int rc = fork();

    if (rc == 0) {
        char* msg = "msg from child\n";

        for (int i = 0; i < 5; i++) {
            write(fd, msg, strlen(msg));
            usleep(100);
        }

    } else {
        char* msg = "msg from parent\n";
        write(fd, msg, strlen(msg));

        for (int i = 0; i < 5; i++) {
            write(fd, msg, strlen(msg));
            usleep(200);
        }

        wait(NULL);
        close(fd);
    }

    return 0;
}

/*
Q: Write a program that opens a file (with the open() system call)
and then calls fork() to create a new process. Can both the child
and parent access the file descriptor returned by open()? What
happens when they are writing to the file concurrently, i.e., at the
same time?

A: the order in which the content will be written to the file is
undeterministic. I needed to introduce loops & delays to force the race
condition to happen, but it was clear
*/
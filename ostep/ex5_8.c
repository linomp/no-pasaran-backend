#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

int main() {
    int status;

    int pipe_fds[2];
    assert(pipe(pipe_fds) >= 0);

    int pipe_read_fd = pipe_fds[0];
    int pipe_write_fd = pipe_fds[1];

    pid_t rc1 = fork();
    assert(rc1 >= 0);

    if (rc1 == 0) {
        char* out_buffer = "msg from child 1!";

        printf("child 1 is writing...\n");
        write(pipe_write_fd, out_buffer, strlen(out_buffer));

        close(pipe_write_fd);
        printf("child 1 terminating...\n");

        exit(0);
    }

    pid_t rc2 = fork();
    assert(rc2 >= 0);

    if (rc2 == 0) {
        printf("child 2 blocked on read()...\n");

        char in_buffer[200];
        read(pipe_read_fd, in_buffer, 200);
        printf("child 2 finished reading...\n");

        char out_buffer[400];
        sprintf(out_buffer, "%s || msg from child 2!", in_buffer);

        printf("child 2 is writing...\n");
        write(pipe_write_fd, out_buffer, strlen(out_buffer));

        close(pipe_write_fd);
        printf("child 2 terminating...\n");

        exit(0);
    }

    // Parent process

    printf("parent waiting for child 1 (pid: %d) ...\n", rc1);
    waitpid(rc1, &status, 0);
    printf("parent waiting for child 2 (pid: %d) ...\n", rc1);
    waitpid(rc2, &status, 0);

    char buffer[200];
    read(pipe_read_fd, buffer, 200);
    printf("parent received: “%s”\n", buffer);

    close(pipe_write_fd);
    printf("parent terminating...\n");

    return 0;
}

/*
Q: Write a program that creates two children, and connects the standard output
of one to the standard input of the other, using the pipe() system call.
*/

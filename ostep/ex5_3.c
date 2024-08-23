#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

int main() {
    int pipe_fds[2];
    assert(pipe(pipe_fds) >= 0);

    int pipe_read_fd = pipe_fds[0];
    int pipe_write_fd = pipe_fds[1];

    int rc = fork();
    assert(rc >= 0);

    if (rc == 0) {
        sleep(2);  // sleep to prove that nothing happens in parent until child
                   // writes to pipe
        char* msg = "msg from child!";
        write(pipe_write_fd, msg, strlen(msg));
    } else {
        char buffer[200];
        printf("parent blocked on read() ...\n");
        read(pipe_read_fd, buffer, 200);
        printf("parent received: “%s”\n", buffer);
        printf("parent terminating\n");
    }

    return 0;
}

/*
Q: Write another program using fork(). The child process should
print “hello”; the parent process should print “goodbye”. You should
try to ensure that the child process always prints first; can you do
this without calling wait() in the parent?

A: can be achieved using a pipe. parent blocks on read() until child writes to
it.
*/
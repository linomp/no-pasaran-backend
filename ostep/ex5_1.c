#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>

int main(int argc, char *argv[]){

    printf("hello world (pid:%d)\n", (int) getpid());

    int x = 100;

    int rc = fork();

    if (rc < 0) {
        fprintf(stderr, "fork failed\n");
        exit(1);
    } else if (rc == 0) {
        printf("hello, I am child (pid:%d) // x=%d\n", (int) getpid(), x);
	    sleep(1);
    } else {
        int wc = wait(NULL);
        printf("hello, I am parent of %d (wc:%d) (pid:%d)  // x=%d\n",
	       rc, wc, (int) getpid(), x);
    }

    return 0;
}
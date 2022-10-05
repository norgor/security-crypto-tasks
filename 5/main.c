#include <stdio.h>
#include <stdlib.h>
#include "utility.h"

int main() {
    char buf[512];

    fgets(&buf, sizeof(buf), stdin);

    const char *esc = html_escape(buf);
    printf("%s\n", esc);
    free(esc);

    return 0;
}

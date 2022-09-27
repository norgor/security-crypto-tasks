#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* html_escape(const char *buf) {
    size_t nl = 64;
    char* new = malloc(nl);
    size_t np = 0;
    while (*buf) {
        const char *c = buf;
        size_t l = 1;
        switch (*c) {
        case '&':
            c = "&amp;";
            l = 5;
            break;
        case '<':
            c = "&lt;";
            l = 4;
            break;
        case '>':
            c = "&gt;";
            l = 4;
            break;
        }
        if (np + l > nl) {
            nl = nl * 2;
            new = realloc(new, nl);
        }
        memcpy(&new[np], c, l);
        np += l;
        buf++;
    }

    return new;
}

int main() {
    char buf[512];

    fgets(&buf, sizeof(buf), stdin);

    const char *esc = html_escape(buf);
    printf("%s\n", esc);
    free(esc);

    return 0;
}

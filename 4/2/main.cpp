#include <string>
#include <iostream>
#include <sstream>
#include <vector>
#include <algorithm>



std::string html_escape(std::string str) {
    struct R {
        char from;
        std::string to;
    };
    
    std::vector<R> replaces = { R{ '&', "&amp;" }, R{'<', "&lt;"}, R{'>', "&gt;"} };

    std::stringstream ss;
    for (char c : str) {
        auto r = std::find_if(replaces.begin(), replaces.end(), [&](auto &r){ return c == r.from; });
        if (r == replaces.end()) {
            ss << c;
        } else {
            ss << r->to;
        }
    }

    return ss.str();
}

int main() {
    std::string in;

    getline(std::cin, in);

    std::string escaped = html_escape(in);
    std::cout << escaped << std::endl;
}

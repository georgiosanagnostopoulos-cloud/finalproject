#include <iostream>
#include "User.h"

int main() {
    User user("Georgios");

    user.startTimer();
    std::cout << "Timer started." << std::endl;

    user.pauseTimer();
    std::cout << "Timer paused." << std::endl;

    user.logSession("work", 25);
    std::cout << "Session logged." << std::endl;

    user.resetTimer();
    std::cout << "Timer reset." << std::endl;

    return 0;
}

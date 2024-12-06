#ifndef USER_H
#define USER_H

#include "Timer.h"
#include "SessionLog.h"

class User {
private:
    std::string username;
    Timer timer;
    SessionLog sessionLog;

public:
    User(const std::string& username);

    void startTimer();
    void pauseTimer();
    void resetTimer();
    void logSession(const std::string& sessionType, int duration);
    SessionLog getSessionLog() const;
};

#endif // USER_H

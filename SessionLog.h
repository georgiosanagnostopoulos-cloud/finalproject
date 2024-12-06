#ifndef SESSIONLOG_H
#define SESSIONLOG_H

#include <vector>
#include <string>

struct Session {
    std::string sessionType; // "work" or "break"
    int duration;            // in minutes
    std::string timestamp;
};

class SessionLog {
private:
    std::vector<Session> logs;

public:
    void logSession(const std::string& sessionType, int duration, const std::string& timestamp);
    std::vector<Session> getSessionHistory() const;
};

#endif // SESSIONLOG_H

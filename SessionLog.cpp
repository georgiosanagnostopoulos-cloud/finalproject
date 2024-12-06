#include "SessionLog.h"

void SessionLog::logSession(const std::string& sessionType, int duration, const std::string& timestamp) {
    logs.push_back({sessionType, duration, timestamp});
}

std::vector<Session> SessionLog::getSessionHistory() const {
    return logs;
}

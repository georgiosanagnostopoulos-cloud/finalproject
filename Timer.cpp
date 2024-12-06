#include "Timer.h"

Timer::Timer(int workDuration, int breakDuration)
    : workDuration(workDuration), breakDuration(breakDuration), status("Idle"), remainingTime(0) {}

void Timer::start() {
    status = "Running";
    remainingTime = workDuration * 60; // Convert to seconds
}

void Timer::pause() {
    if (status == "Running") {
        status = "Paused";
    }
}

void Timer::reset() {
    status = "Idle";
    remainingTime = 0;
}

std::string Timer::getStatus() const {
    return status;
}

int Timer::getRemainingTime() const {
    return remainingTime;
}

void Timer::updateDurations(int newWorkDuration, int newBreakDuration) {
    workDuration = newWorkDuration;
    breakDuration = newBreakDuration;
}

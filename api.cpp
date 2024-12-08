#include "crow_all.h"
#include "ConfigManager.h"
#include "User.h"
#include "Timer.h"
#include "SessionLog.h"

// Instantiate user (mocked for demonstration)
User user("Georgios");

int main() {
    crow::SimpleApp app; // Crow application

    // Endpoint: Get user details
    CROW_ROUTE(app, "/api/user")
    ([]() {
        return crow::json::wvalue{{"username", user.getUsername()}};
    });

    // Endpoint: Start a Pomodoro timer
    CROW_ROUTE(app, "/api/timers/start").methods("POST"_method)
    ([](const crow::request& req) {
        user.startTimer();
        return crow::response(201, "Timer started");
    });

    // Endpoint: Pause the Pomodoro timer
    CROW_ROUTE(app, "/api/timers/pause").methods("POST"_method)
    ([](const crow::request& req) {
        user.pauseTimer();
        return crow::response(200, "Timer paused");
    });

    // Endpoint: Reset the Pomodoro timer
    CROW_ROUTE(app, "/api/timers/reset").methods("POST"_method)
    ([](const crow::request& req) {
        user.resetTimer();
        return crow::response(200, "Timer reset");
    });

    // Endpoint: Log a completed Pomodoro session
    CROW_ROUTE(app, "/api/sessions").methods("POST"_method)
    ([](const crow::request& req) {
        auto body = crow::json::load(req.body);
        if (!body) {
            return crow::response(400, "Invalid JSON");
        }

        std::string sessionType = body["sessionType"].s();
        int duration = body["duration"].i();
        user.logSession(sessionType, duration);

        return crow::response(201, "Session logged");
    });

    // Endpoint: Retrieve all session logs
    CROW_ROUTE(app, "/api/sessions")
    ([]() {
        auto logs = user.getSessionLog().getSessionHistory();
        crow::json::wvalue result;
        result["sessions"] = crow::json::wvalue::list();
        for (const auto& log : logs) {
            result["sessions"].push_back(log.dump());
        }
        return result;
    });

    // Run the server on port 8080
    app.port(8080).multithreaded().run();
}

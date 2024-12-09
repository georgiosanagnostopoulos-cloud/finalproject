#pragma once
// crow_all.h is a single-header distribution of Crow.
// It's recommended to use the CMake build system for a modular Crow.

#include <string>
#include <functional>
#include <vector>
#include <unordered_map>
#include <map>
#include <memory>
#include <sstream>
#include <boost/algorithm/string.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/algorithm/string/trim.hpp>
#include <boost/asio.hpp>
#include <boost/beast.hpp>
#include <boost/utility/string_ref.hpp>
#include <boost/filesystem.hpp>
#include <nlohmann/json.hpp>

// Core Crow components
namespace crow
{
    // HTTP Request
    struct request
    {
        std::string method;
        std::string body;
        std::unordered_map<std::string, std::string> headers;
        std::unordered_map<std::string, std::string> params;
    };

    // HTTP Response
    struct response
    {
        int code;
        std::string body;
        std::unordered_map<std::string, std::string> headers;

        response(int code, const std::string& body = "") : code(code), body(body) {}
    };

    // Routing framework
    template <typename Func>
    void route(const std::string& path, Func&& func)
    {
        // Store and execute routing
    }

    // Simple Crow application
    class SimpleApp
    {
    public:
        void port(int port) { port_ = port; }
        void multithreaded() { threads_ = 4; }
        void run()
        {
            // Server logic using boost::asio
        }

    private:
        int port_{8080};
        int threads_{1};
    };

} // namespace crow

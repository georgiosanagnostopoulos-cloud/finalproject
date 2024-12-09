all: main api

main: main.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp
	g++ -o main main.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp -I. -I./nlohmann -lpthread

api: api.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp
	g++ -o api api.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp -I. -I./nlohmann -lpthread

clean:
	rm -f main api

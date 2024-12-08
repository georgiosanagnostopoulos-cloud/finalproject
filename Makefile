all: main

main: main.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp
	g++ -o main main.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp -I. -I./nlohmann

clean:
	rm -f main

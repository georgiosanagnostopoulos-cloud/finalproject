all: main tests

main: main.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp
	g++ -o main main.cpp Timer.cpp SessionLog.cpp User.cpp ConfigManager.cpp -I. -lnlohmann_json

test_sessionlog: test_sessionlog.cpp SessionLog.cpp
	g++ -o test_sessionlog test_sessionlog.cpp SessionLog.cpp -I. -lnlohmann_json

test_timer: test_timer.cpp Timer.cpp
	g++ -o test_timer test_timer.cpp Timer.cpp -I. -lnlohmann_json

test_user: test_user.cpp User.cpp Timer.cpp SessionLog.cpp
	g++ -o test_user test_user.cpp User.cpp Timer.cpp SessionLog.cpp -I. -lnlohmann_json

clean:
	rm -f main test_sessionlog test_timer test_user

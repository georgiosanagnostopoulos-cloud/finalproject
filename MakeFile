all: main test_timer test_sessionlog test_user

main: main.cpp Timer.cpp SessionLog.cpp User.cpp
	g++ -o main main.cpp Timer.cpp SessionLog.cpp User.cpp

test_timer: test_timer.cpp Timer.cpp
	g++ -o test_timer test_timer.cpp Timer.cpp

test_sessionlog: test_sessionlog.cpp SessionLog.cpp
	g++ -o test_sessionlog test_sessionlog.cpp SessionLog.cpp

test_user: test_user.cpp User.cpp Timer.cpp SessionLog.cpp
	g++ -o test_user test_user.cpp User.cpp Timer.cpp SessionLog.cpp

clean:
	rm -f main test_timer test_sessionlog test_user

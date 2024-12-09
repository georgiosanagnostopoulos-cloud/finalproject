package main

import (
    "time"
    "net/http"    // Add this line
    "github.com/labstack/echo/v4"
)



// Simulated data structures, since original code is C++.
// In a production environment, you'd implement these more robustly.

// Timer structure
type Timer struct {
	WorkDuration  int    `json:"workDuration"`
	BreakDuration int    `json:"breakDuration"`
	Status        string `json:"status"`
	RemainingTime int    `json:"remainingTime"`
}

func NewTimer(work, breakDur int) *Timer {
	return &Timer{WorkDuration: work, BreakDuration: breakDur, Status: "Idle", RemainingTime: 0}
}

func (t *Timer) Start() {
	if t.WorkDuration > 0 && t.BreakDuration > 0 {
		t.Status = "Running"
		t.RemainingTime = t.WorkDuration * 60 // convert minutes to seconds for simulation
	} else {
		t.Status = "Idle"
	}
}

func (t *Timer) Pause() {
	if t.Status == "Running" {
		t.Status = "Paused"
	}
}

func (t *Timer) Reset() {
	t.Status = "Idle"
	t.RemainingTime = 0
}

func (t *Timer) UpdateDurations(work, breakDur int) {
	t.WorkDuration = work
	t.BreakDuration = breakDur
}

// SessionLog
type Session struct {
	Type      string `json:"sessionType"`
	Duration  int    `json:"duration"`
	Timestamp string `json:"timestamp"`
}

type SessionLog struct {
	Logs []Session `json:"logs"`
}

func (s *SessionLog) LogSession(sessionType string, duration int) {
	// For simplicity, we’ll use current date as timestamp
	s.Logs = append(s.Logs, Session{
		Type:      sessionType,
		Duration:  duration,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	})
}

func (s *SessionLog) GetHistory() []Session {
	return s.Logs
}

// User
type User struct {
	Username   string
	Timer      *Timer
	SessionLog *SessionLog
}

func NewUser(username string, workDur, breakDur int) *User {
	return &User{
		Username:   username,
		Timer:      NewTimer(workDur, breakDur),
		SessionLog: &SessionLog{Logs: []Session{}},
	}
}

func (u *User) StartTimer() {
	u.Timer.Start()
}

func (u *User) PauseTimer() {
	u.Timer.Pause()
}

func (u *User) ResetTimer() {
	u.Timer.Reset()
}

func (u *User) UpdateTimerDurations(workDur, breakDur int) {
	u.Timer.UpdateDurations(workDur, breakDur)
}

// Global user (for demo; in production you would have real authentication)
var currentUser = NewUser("Georgios", 25, 5)

func main() {
	e := echo.New()

	// POST /api/timers
	e.POST("/api/timers", func(c echo.Context) error {
		var req struct {
			WorkDuration  int `json:"workDuration"`
			BreakDuration int `json:"breakDuration"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		}

		if req.WorkDuration <= 0 || req.BreakDuration <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Durations must be > 0"})
		}

		currentUser.UpdateTimerDurations(req.WorkDuration, req.BreakDuration)
		currentUser.StartTimer()

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":        currentUser.Timer.Status,
			"workDuration":  currentUser.Timer.WorkDuration,
			"breakDuration": currentUser.Timer.BreakDuration,
		})
	})

	// GET /api/timers
	e.GET("/api/timers", func(c echo.Context) error {
		t := currentUser.Timer
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":        t.Status,
			"remainingTime": t.RemainingTime,
		})
	})

	// PUT /api/timers
	e.PUT("/api/timers", func(c echo.Context) error {
		var req struct {
			WorkDuration  int `json:"workDuration"`
			BreakDuration int `json:"breakDuration"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		}
		if req.WorkDuration <= 0 || req.BreakDuration <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Durations must be > 0"})
		}
		currentUser.UpdateTimerDurations(req.WorkDuration, req.BreakDuration)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":        currentUser.Timer.Status,
			"workDuration":  currentUser.Timer.WorkDuration,
			"breakDuration": currentUser.Timer.BreakDuration,
		})
	})

	// DELETE /api/timers
	e.DELETE("/api/timers", func(c echo.Context) error {
		currentUser.ResetTimer()
		return c.NoContent(http.StatusNoContent)
	})

	// POST /api/sessions
	e.POST("/api/sessions", func(c echo.Context) error {
		var req struct {
			SessionType string `json:"sessionType"`
			Duration    int    `json:"duration"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		}
		if req.SessionType == "" || req.Duration <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid session data"})
		}

		currentUser.SessionLog.LogSession(req.SessionType, req.Duration)
		lastSession := currentUser.SessionLog.Logs[len(currentUser.SessionLog.Logs)-1]
		return c.JSON(http.StatusCreated, lastSession)
	})

	// GET /api/sessions
	e.GET("/api/sessions", func(c echo.Context) error {
		return c.JSON(http.StatusOK, currentUser.SessionLog.GetHistory())
	})

	// Start server on port 8080
	e.Logger.Fatal(e.Start(":7080"))
}

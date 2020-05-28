package main

import (
	"fmt"
	"net/http/httputil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

// openAccessLog opens the access log file for writing
func openAccessLog(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // #nosec
    if err != nil {
        log.Fatal(fmt.Sprintf("Can't open access log file: %s", err))
	}
	
	return f
}

// formatAccessLogMessage formats access log messages in Common Log Format
func formatAccessLogMessage(c echo.Context, responseStatus int, responseLength int) string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Can't retrieve current user")
	}

	loggedResponseLength := strconv.Itoa(responseLength)
	if(responseLength == 0) {
		loggedResponseLength = "-" 
	}
	
	timeNow := time.Now()

	dump, err := httputil.DumpRequest(c.Request(), false)
	requestStatusLine := strings.Split(string(dump), "\n")[0]
	requestStatusLine = strings.Replace(requestStatusLine, "\r", "", -1)


	logMessage := fmt.Sprintf(
		"%s - %s [%s] \"%s\" %d %s \"%s\"",
		c.RealIP(),
		currentUser.Username,
		timeNow.Format("02/Jan/2006:15:04:05 -0700"),
		requestStatusLine,
		responseStatus,
		loggedResponseLength,
		c.Request().UserAgent(),
	)

	return logMessage
}

// openAccessLog opens the error log file for writing
func openErrorLog(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // #nosec
    if err != nil {
        log.Fatal(fmt.Sprintf("Can't open error log file: %s", err))
	}
	
	return f
}

// formatErrorLogMessage formats access log messages in Common Log Format
func formatErrorLogMessage(c echo.Context, message string, responseStatus int) string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Can't retrieve current user")
	}
	
	timeNow := time.Now()

	dump, err := httputil.DumpRequest(c.Request(), false)
	requestStatusLine := strings.Split(string(dump), "\n")[0]
	requestStatusLine = strings.Replace(requestStatusLine, "\r", "", -1)


	logMessage := fmt.Sprintf(
		"%s - %s [%s] \"%s\" %d %s \"%s\"",
		c.RealIP(),
		currentUser.Username,
		timeNow.Format("02/Jan/2006:15:04:05 -0700"),
		requestStatusLine,
		responseStatus,
		message,
		c.Request().UserAgent(),
	)

	return logMessage
}

func writeAccessLog(accessLog *os.File, errorLog *os.File, c echo.Context, responseStatus int, responseLength int) {
	_, err := accessLog.WriteString(formatAccessLogMessage(c, responseStatus, responseLength) + "\n")
	if err != nil {
		writeErrorLog(errorLog, c, fmt.Sprintf("Error writing to access log: %s", err))
		return500(c)
	}
}

func writeErrorLog(errorLog *os.File, c echo.Context, message string) {
	_, err := errorLog.WriteString(formatErrorLogMessage(c, message, 500) + "\n")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error writing to error log: %s", err))
	}
}
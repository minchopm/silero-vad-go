package speech

/*
#include "ort_bridge.h"
*/
import "C"

type LogLevel int

const (
    LogLevelFatal LogLevel = iota
    LogLevelError
    LogLevelWarning
    LogLevelInfo
    LogLevelVerbose
    // Add other levels as needed
)

// Helper function to convert LogLevel to C.OrtLoggingLevel
func (l LogLevel) toCLogLevel() C.OrtLoggingLevel {
    switch l {
    case LogLevelFatal:
        return C.ORT_LOGGING_LEVEL_FATAL
    case LogLevelError:
        return C.ORT_LOGGING_LEVEL_ERROR
    case LogLevelWarning:
        return C.ORT_LOGGING_LEVEL_WARNING
    case LogLevelInfo:
        return C.ORT_LOGGING_LEVEL_INFO
    case LogLevelVerbose:
        return C.ORT_LOGGING_LEVEL_VERBOSE
    default:
        return C.ORT_LOGGING_LEVEL_WARNING // Default level
    }
}

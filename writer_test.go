package logger

import "testing"

func BenchmarkEnableDirect(b *testing.B) {
	consoleEnabler := func(l Level, scope string) bool {
		switch scope {
		case "ice", "sctp", "dtls", "datachannel":
			return l > WarnLevel
		default:
			return true
		}
	}
	// influxEnabler := func(l log.Level, s string) bool { return l >= log.ErrorLevel }
	stackEnabler := func(l Level, scope string) bool {
		return l >= ErrorLevel
	}

	w := ConsoleWriter(false, consoleEnabler, stackEnabler)
	for i := 0; i < b.N; i++ {
		w.enabler(WarnLevel, "ice")
	}
}

func BenchmarkEnableMap(b *testing.B) {
	consoleEnabler := func(l Level, scope string) bool {
		switch scope {
		case "ice", "sctp", "dtls", "datachannel":
			return l > WarnLevel
		default:
			return true
		}
	}
	// influxEnabler := func(l log.Level, s string) bool { return l >= log.ErrorLevel }
	stackEnabler := func(l Level, scope string) bool {
		return l >= ErrorLevel
	}

	w := ConsoleWriter(false, consoleEnabler, stackEnabler)
	for i := 0; i < b.N; i++ {
		w.isEnable(WarnLevel, "ice")
	}
}

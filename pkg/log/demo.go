package log

func demo() {
	defer Sync()
	Info("demo1:", String("app", "start ok"), Int("major version", 2))
	Warn("demo2:", String("app", "end over"), Int("major version", 3))
}

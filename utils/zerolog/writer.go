package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

var writerMap = make(map[string]io.Writer)

func AppendToConsole() {
	writerMap["console"] = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = "15:04:05"
	})
	refreshLogger()
}

func AppendToFile(logfile string) {
	if logfile == "" {
		return
	}
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("打开日志文件失败")
	}
	writerMap["file"] = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = "15:04:05"
		w.NoColor = true
		w.Out = file
	})
	refreshLogger()
}

func AppendToWriter(writer io.Writer) {
	writerMap["writer"] = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = "15:04:05"
		w.NoColor = true
		w.Out = writer
	})
	refreshLogger()
}

func refreshLogger() {
	var writers []io.Writer
	for _, w := range writerMap {
		writers = append(writers, w)
	}
	writer := io.MultiWriter(writers...)
	log.Logger = log.Output(writer)
}

package server

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sre-works/metric"

	"strconv"

	"go.opentelemetry.io/otel"
)

const name = "practice.roll_dice"

var tracer = otel.Tracer(name)

func Rolldice(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), name)
	defer span.End()

	metric.APICounter.Add(r.Context(), 1)

	roll := 1 + rand.Intn(6)

	span.AddEvent("roll dice called")

	var msg string
	if player := r.PathValue("player"); player != "" {
		msg = player + " is rolling the dice"
	} else {
		msg = "Anonymous player is rolling the dice"
	}
	fmt.Println(msg)
	// zaplog.Logger.Info("roll", zap.String("player", msg), zap.Int("result", roll))

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v", err)
	}
}

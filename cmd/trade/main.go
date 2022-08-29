package main

import (
	"net/http"
	"os"
	"time"

	"git.mnc.sh/ilazarev/trade"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.StandardLogger()

	value = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "moex",
			Subsystem: "stocks",
			Name:      "value",
			Help:      "Current stock market values",
		},
		[]string{"security"},
	)

	price = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "moex",
			Subsystem: "stocks",
			Name:      "price",
			Help:      "Current stock market prices",
		},
		[]string{"security", "variety"},
	)
)

func init() {
	prometheus.MustRegister(value, price)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	log.SetOutput(os.Stdout)

	exc := trade.NewExchange()

	eng, err := exc.Engine("stock")
	if err != nil {
		log.Errorf("main: %v", err)
		os.Exit(1)
	}

	log.Infof("eng: %+v", eng)

	mkt, err := eng.Market("shares")
	if err != nil {
		log.Errorf("main: %v", err)
		os.Exit(1)
	}

	log.Infof("mkt: %+v", mkt)

	brd, err := mkt.Board("TQBR")
	if err != nil {
		log.Errorf("main: %v", err)
		os.Exit(1)
	}

	log.Infof("brd: %+v", brd)

	go func(brd trade.Board) {
		ticker := time.NewTicker(time.Second * 10)
		for ; true; <-ticker.C {
			start := time.Now()

			s, err := brd.Securities()
			if err != nil {
				log.Errorf("main: ticker: %v", err)
			}

			for _, m := range s {
				open, low, high, last, v, size, err := m.Rates()
				if err != nil {
					log.Errorf("%v", err)
					continue
				}

				id := m.Id()

				value.WithLabelValues(id).Set(v)
				price.WithLabelValues(id, "open").Set(open)
				price.WithLabelValues(id, "low").Set(low)
				price.WithLabelValues(id, "high").Set(high)
				price.WithLabelValues(id, "last").Set(last)

				log.Infof("%s;%f;%f;%f;%f;%f;%f", id, open, low, high, last, v, size)
			}

			log.Infof("rates gathered in %s", time.Since(start))
		}
	}(brd)

	r.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":3000", r)
}

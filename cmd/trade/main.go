package main

import (
	"net/http"
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
			Name:      "value_total",
			Help:      "Current stock market values",
		},
		[]string{"security"},
	)

	price = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "moex",
			Subsystem: "stocks",
			Name:      "price_total",
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

	e := trade.NewExchange()
	go func(e trade.Exchange) {
		ticker := time.NewTicker(time.Minute)
		for ; true; <-ticker.C {
			start := time.Now()

			s, err := e.Securities()
			if err != nil {
				log.Errorf("%v", err)
				continue
			}

			data := make(map[string][]trade.Marketdata)
			for _, sec := range s {
				if _, ok := data[sec.Id()]; !ok {
					mkd, err := sec.Marketdata()
					if err != nil {
						log.Errorf("%v", err)
						continue
					}

					data[sec.Name()] = mkd
				}
			}

			for _, m := range data {
				for _, mkd := range m {
					id, open, low, high, last, market, today, v, _, err := mkd.Rates()
					if err != nil {
						log.Errorf("%v", err)
						continue
					}

					value.WithLabelValues(id).Set(v)
					price.WithLabelValues(id, "open").Set(open)
					price.WithLabelValues(id, "low").Set(low)
					price.WithLabelValues(id, "high").Set(high)
					price.WithLabelValues(id, "last").Set(last)
					price.WithLabelValues(id, "market").Set(market)
					price.WithLabelValues(id, "today").Set(today)
				}
			}

			log.Infof("rates gathered in %s", time.Since(start))
		}
	}(e)

	r.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":3000", r)
}

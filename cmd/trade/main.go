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
	price = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   "moex",
			Subsystem:   "stocks",
			Name:        "price",
			Help:        "Current stock market prices",
			ConstLabels: map[string]string{},
		},
		[]string{"security", "variety"},
	)

	value = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   "moex",
			Subsystem:   "stocks",
			Name:        "value",
			Help:        "Current stock market values",
			ConstLabels: map[string]string{},
		},
		[]string{"security"},
	)

	log = logrus.StandardLogger()
)

func init() {
	prometheus.MustRegister(price, value)
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

			rates := make(map[string]map[string]float64)
			for _, m := range data {
				for _, mkd := range m {
					id, open, low, high, last, market, today, v, _, err := mkd.Rates()
					if err == nil {
						if _, ok := rates[id]; !ok {
							value.WithLabelValues(id).Set(v)

							price.WithLabelValues(id, "open").Set(open)
							price.WithLabelValues(id, "low").Set(low)
							price.WithLabelValues(id, "high").Set(high)
							price.WithLabelValues(id, "last").Set(last)
							price.WithLabelValues(id, "market").Set(market)
							price.WithLabelValues(id, "today").Set(today)
						}
					}
				}
			}

			log.Infof("rates gathered in %s", time.Since(start))
		}
	}(e)

	r.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":3000", r)
}

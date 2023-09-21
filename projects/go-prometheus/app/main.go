package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Employee struct {
	Name        string `json:"name"`
	Age         int    `json:"int"`
	Designation string `json:"designation"`
}

type metrics struct {
	emps prometheus.Gauge
	info *prometheus.GaugeVec //metadat of the app
}

var emp []Employee
var info string

func init() {
	info = "jai-Mahishmathi"
	emp = []Employee{
		{"Mike", 23, "Blockchain Developer"},
		{"Mawa", 25, "Backend Developer"},
		{"John", 27, "Full Stack Developer"},
	}
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		emps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "application",
			Name:      "employees_list",
			Help:      "active employees",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "kingdom",
			Help: "Hail Mahishmathi",
		},
			[]string{"info"}),
	}
	reg.MustRegister(m.emps, m.info)
	return m
}

func main() {
	// fmt.Println("---go-prometheus---")
	reg := prometheus.NewRegistry()
	//optional
	reg.MustRegister(collectors.NewGoCollector())
	m := NewMetrics(reg)
	m.emps.Set(float64(len(emp)))
	m.info.With(prometheus.Labels{"info": info})

	//multiple servers..
	dMux := http.NewServeMux()
	dMux.HandleFunc("/emp", getEmployees)

	pMux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":1357", dMux))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":3579", pMux))
	}()

	select {}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	z, err := json.Marshal(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(z)

}

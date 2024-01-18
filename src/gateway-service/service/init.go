package service

import (
	circuitbreaker "lab2/src/circuit-breaker"
	job_scheduler "lab2/src/job-scheduler"
	"time"
)

var ticketcb *circuitbreaker.CircuitBreaker
var flightcb *circuitbreaker.CircuitBreaker
var bonuscb *circuitbreaker.CircuitBreaker

var jobscheduler *job_scheduler.JobScheduler

func init() {
	var ticketst circuitbreaker.Settings
	ticketst.Name = "Tickets Circuit Breaker"
	ticketst.ReadyToTrip = func(counts circuitbreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	ticketcb = circuitbreaker.NewCircuitBreaker(ticketst)

	var flightst circuitbreaker.Settings
	flightst.Name = "Flights Circuit Breaker"
	flightst.ReadyToTrip = func(counts circuitbreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	flightcb = circuitbreaker.NewCircuitBreaker(flightst)

	var bonusst circuitbreaker.Settings
	bonusst.Name = "Bonuses Circuit Breaker"
	bonusst.ReadyToTrip = func(counts circuitbreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	bonuscb = circuitbreaker.NewCircuitBreaker(bonusst)

	jobscheduler = job_scheduler.NewJobScheduler(10 * time.Second)
	jobscheduler.Start()
}

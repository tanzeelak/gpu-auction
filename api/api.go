package api

import (
	"github.com/tanzeelak/gpu-auction/services"
)

type API struct {
	scheduler *services.Scheduler
}

func New(s *services.Scheduler) *API {
	return &API{
		scheduler: s,
	}
}

func (a *API) makeAccount(name string) {

}

func (a *API) login(username string) {

}

func (a *API) getSchedule() map[string][]services.GPUBid {
	return a.scheduler.GetSchedule()
}

func bidOnGPU(username string, start int, end int, bid int, sameCluster bool) bool {
	return false
}

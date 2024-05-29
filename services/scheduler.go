package services

type GPUBid struct {
	id  string
	bid string
}

type Scheduler struct{}

func (s *Scheduler) New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) GetSchedule() map[string][]GPUBid {
	/*
		- Get Database from SQL
	*/
	return make(map[string][]GPUBid)
}

func (s *Scheduler) AddBid(
	username string,
	startTime int,
	endTime int,
	startGPU int,
	endGPU int,
	bid int,
	sameCluster bool) bool {
	/*
		- get the schedule
		- iterate over start to end times
			- iterate over start to end GPUs
				- key to the time + index to GPU index
				- place bid in schedule
	*/

	return true
}

func (s *Scheduler) DecideBid() bool {
	/*
		init winners array
		key to the first time
		for each GPU
			extract winner (higher price)
		remove most recent timeslot
		add another timeslot to the arr
		return winners array
	*/
	return true
}

package db

type Database struct {
}

// access DB

func New() *Database {
	return &Database{}
}

// get accounts
func (d *Database) GetAccounts() {

}

// add account
func (d *Database) CreateAccount() {

}

// get schedule
func (d *Database) GetSchedule() map[string]map[string]string {
	return make(map[string]map[string]string)
}

// add schedule
func (d *Database) CreateBid() {
}

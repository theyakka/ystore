// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore_test

import (
	"github.com/theyakka/ystore"
	"os"
	"testing"
)

type Person struct {
	FirstName  string   `ystore:"first_name"`
	LastName   string   `ystore:"last_name"`
	Age        int      `ystore:"age"`
	CurrentJob Job      `ystore:"current_job"`
	PastJobs   []Job    `ystore:"past_jobs"`
	Skills     []string `ystore:"skills"`
}

type Job struct {
	Company     string `ystore:"company"`
	YearStarted int    `ystore:"year_started"`
}

func setupStore() *ystore.Store {
	store := ystore.NewStore()
	return store

}

func makePerson() *Person {
	return &Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       30,
		CurrentJob: Job{
			Company:     "ACME Co",
			YearStarted: 2019,
		},
		PastJobs: []Job{
			{
				Company:     "Things R Us",
				YearStarted: 2017,
			},
			{
				Company:     "Good Peeps Inc",
				YearStarted: 2014,
			},
		},
		Skills: []string{
			"Professional", "Hard Working", "Independent", "Strong",
		},
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

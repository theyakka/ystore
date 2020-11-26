// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore_test

import (
	"github.com/theyakka/ystore"
	"log"
	"testing"
)

func TestAddStructs(t *testing.T) {
	store := ystore.NewStore()
	person := makePerson()
	err := store.AddData(person, nil)
	if err != nil {
		log.Println(err)
		return
	}
	jj := store.Get("current_job.company")
	log.Println(jj)
	jobObj := store.GetIndexed("past_jobs", 1, "company")
	if jobObj != nil {

	}

}

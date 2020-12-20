// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore_test

import (
	"log"
	"testing"

	"github.com/theyakka/ystore"
)

func TestAddStructs(t *testing.T) {
	store := ystore.NewStore()
	person := makePerson()
	err := store.AddData(person, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var xx string
	qerr := store.NewQuery().
		Get("data.monkey.something").
		Run(&xx, nil)
	log.Println("q1 ---")
	log.Println(xx, qerr)
	log.Println("---")

	var yy []string
	qerr2 := store.NewQuery().
		Get("past_jobs.company").
		Run(&yy, nil)
	log.Println("q2 ---")
	log.Println(yy, qerr2)
	log.Println("---")

	//jobObj := store.GetIndexed("past_jobs", 1, "company")
	//if jobObj != nil {
	//
	//}

}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		store := ystore.NewStore()
		person := makePerson()
		err := store.AddData(person, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

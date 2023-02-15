package k4

import (
	"log"
	"testing"
)

func TestID(t *testing.T) {
	routines := 50
	iter := 100
	test := routines * iter
	idStream := make(chan ID, test)

	for i := 0; i < routines; i++ {
		go func() {
			for i := 0; i < iter; i++ {
				gen := NewID()
				idStream <- gen
				log.Println(gen.Full())
			}
		}()
	}

	generated := map[string]bool{}
	lastProcess := 0
	for processed := 0; processed < test; processed++ {
		gen := <-idStream
		if _, found := generated[gen.UID()]; found {
			t.Fatal("Duplicate ID generated ", gen)
		} else {
			generated[gen.UID()] = true
		}
		lastProcess = processed
		if !IDFromString(gen.Full()).HasValidVerification() {
			t.Fatal("Invalid verification")
		}
	}

	log.Println("Processed ", lastProcess+1, " IDs")

}

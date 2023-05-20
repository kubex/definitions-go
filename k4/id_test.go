package k4

import (
	"log"
	"testing"
)

func benchmarkIDGeneration(idGen IDGenerator, b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		idGen.New()
	}
}
func BenchmarkIDGeneration(b *testing.B) {
	benchmarkIDGeneration(globalIDHost, b)
}

func BenchmarkIDGenerationNano(b *testing.B) {
	idGen := DefaultIDGenerator()
	idGen.idLength = 16
	idGen.SetTimeSize(TimeGeneratorNano)
	benchmarkIDGeneration(idGen, b)
}

func BenchmarkIDGenerationMicro(b *testing.B) {
	idGen := DefaultIDGenerator()
	idGen.SetTimeSize(TimeGeneratorMicro)
	benchmarkIDGeneration(idGen, b)
}
func BenchmarkIDGenerationMilli(b *testing.B) {
	idGen := DefaultIDGenerator()
	idGen.SetTimeSize(TimeGeneratorMilli)
	benchmarkIDGeneration(idGen, b)
}

func BenchmarkIDGenerationSmall(b *testing.B) {
	idGen := DefaultIDGenerator()
	idGen.SetBaseLength(6)
	idGen.SetTimeSize(TimeGeneratorSecond)
	benchmarkIDGeneration(idGen, b)
}

func TestID(t *testing.T) {
	routines := 1000
	iter := 10000
	test := routines * iter
	idStream := make(chan ID, test)

	for i := 0; i < routines; i++ {
		go func(generator IDGenerator) {
			for i := 0; i < iter; i++ {
				gen := generator.New()
				idStream <- gen
				//if i%500 == 0 {
				//log.Println(gen.String())
				//log.Println(gen.UUID())
				//}
			}
		}(DefaultIDGenerator())
	}
	generated := map[string]bool{}
	lastProcess := 0
	for processed := 0; processed < test; processed++ {
		gen := <-idStream
		if _, found := generated[gen.String()]; found {
			t.Fatal("Duplicate ID generated ", gen)
		} else {
			generated[gen.String()] = true
		}
		lastProcess = processed
		if !IDFromString(gen.String()).IsValid() {
			t.Fatal("Invalid verification")
		}
	}

	log.Println("Processed ", lastProcess+1, " IDs")

}

func TestIDUUID(t *testing.T) {
	id := NewID()
	originalUUID := id.UUID()
	log.Println(originalUUID)
	id2 := IDFromUUID(originalUUID)
	log.Println(id2.UUID())
	if id.String() != id2.String() {
		log.Fatal("UUID conversion failed")
	}
	if id.IsValid() && !id2.IsValid() {
		log.Fatal("UUID conversion failed checksum")
	}
}

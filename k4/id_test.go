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
	idGen := IDGen()
	idGen.SetBaseLength(16)
	idGen.SetTimeSize(TimeGeneratorNano)
	benchmarkIDGeneration(idGen, b)
}

func BenchmarkIDGenerationMicro(b *testing.B) {
	idGen := IDGen()
	idGen.SetBaseLength(16)
	idGen.SetTimeSize(TimeGeneratorMicro)
	benchmarkIDGeneration(idGen, b)
}
func BenchmarkIDGenerationMilli(b *testing.B) {
	idGen := IDGen()
	idGen.SetBaseLength(16)
	idGen.SetTimeSize(TimeGeneratorMilli)
	benchmarkIDGeneration(idGen, b)
}

func BenchmarkIDGenerationSmall(b *testing.B) {
	idGen := IDGen()
	idGen.SetBaseLength(6)
	idGen.SetTimeSize(TimeGeneratorSecond)
	benchmarkIDGeneration(idGen, b)
}

func TestID(t *testing.T) {
	routines := 200
	iter := 1000
	test := routines * iter
	idStream := make(chan ID, test)

	idGen := IDGen()
	for i := 0; i < routines; i++ {
		go func() {
			for i := 0; i < iter; i++ {
				gen := idGen.New()
				idStream <- gen
				log.Println(gen.String())
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
		if !IDFromString(gen.String()).IsValid() {
			t.Fatal("Invalid verification")
		}
	}

	log.Println("Processed ", lastProcess+1, " IDs")

}

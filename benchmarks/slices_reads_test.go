package benchmarks

import "testing"

func BenchmarkReadSlicePointers(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(*structRecordsPtr); j++ {
			_ = (*structRecordsPtr)[j]
		}
	}
}

func BenchmarkReadSliceNoPointers(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(structRecords); j++ {
			_ = structRecords[j]
		}
	}
}

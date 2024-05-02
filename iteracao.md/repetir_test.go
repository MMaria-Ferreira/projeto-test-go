package interacao

import "testing"

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}
func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}

}

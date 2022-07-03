package repeater

import ("testing"
		"fmt")

func TestIterator(t *testing.T){



	repeated := Repeat("a")
	expected := 5

	if expected != repeated {
		t.Errorf("expected %d but got repeated %d",expected,repeated)
	}


}

func BenchmarkRepeat(b *testing.B){
	for i:=0; i < b.N ; i++ {
		Repeat("a")
	}
}

func ExampleRepeat (){
	repeated := Repeat("b")
	fmt.Println(repeated)
	// Output: 5
}
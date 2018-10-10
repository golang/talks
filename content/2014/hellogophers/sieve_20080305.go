// +build ignore,OMIT

package Main

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch *chan> int) {
	for i := 2; ; i++ {
		>ch = i;	// Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in *chan< int, out *chan> int, prime int) {
	for ; ; {
		i := <in;	// Receive value of new variable 'i' from 'in'.
		if i % prime != 0 {
			>out = i;	// Send 'i' to channel 'out'.
		}
	}
}
// BREAK OMIT
// The prime sieve: Daisy-chain Filter processes together.
func Sieve() {
	ch := new(chan int);  // Create a new channel.
	go Generate(ch);      // Start Generate() as a subprocess.
	for ; ; {
		prime := <ch;
		printf("%d\n", prime);
		ch1 := new(chan int);
		go Filter(ch, ch1, prime);
		ch = ch1;
	}
}

func Main() {
	Sieve();
}

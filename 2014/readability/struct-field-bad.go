// +build OMIT

package sample // OMIT

type Modifier struct {
	pmod   *profile.Modifier
	cache  map[string]time.Time
	client *client.Client
	mu     sync.RWMutex // HL
}

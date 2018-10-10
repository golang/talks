// +build OMIT

package sample // OMIT

type Modifier struct {
	client *client.Client

	mu    sync.RWMutex // HL
	pmod  *profile.Modifier
	cache map[string]time.Time
}

// +build OMIT

package mart

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"appengine"
	"appengine/datastore"
	"appengine/delay"
	"appengine/mail"
	"appengine/user"

	"github.com/mjibson/appstats"
)

func init() {
	http.HandleFunc("/", front)
	http.Handle("/checkout", appstats.NewHandler(checkout))
	http.HandleFunc("/admin/populate", adminPopulate)
}

func front(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello, welcome to Gopher Mart!")
}

const (
	numItems = 100 // number of different items for sale

	appAdmin = "noreply@google.com" // an admin of this app, for sending mail
)

// Item represents an item for sale in Gopher Mart.
type Item struct {
	Name  string
	Price float64
}

func itemKey(c appengine.Context, i int) *datastore.Key {
	return datastore.NewKey(c, "Item", fmt.Sprintf("item%04d", i), 0, nil)
}

func checkout(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	num, err := strconv.Atoi(r.FormValue("num"))
	if err == nil && (num < 1 || num > 30) {
		err = fmt.Errorf("%d out of range [1,30]", num)
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("bad number of items: %v", err), http.StatusBadRequest)
		return
	}

	// Pick random items.
	keys := make([]*datastore.Key, num)
	for i := range keys {
		keys[i] = itemKey(c, rand.Intn(numItems))
	}

	// Dumb load.
	var items []*Item
	for _, key := range keys {
		item := new(Item)
		if err := datastore.Get(c, key, item); err != nil {
			// ...
			http.Error(w, fmt.Sprintf("datastore.Get: %v", err), http.StatusBadRequest) // OMIT
			return                                                                      // OMIT
		}
		items = append(items, item)
	}

	// Print items.
	var b bytes.Buffer
	fmt.Fprintf(&b, "Here's what you bought:\n")
	var sum float64
	for _, item := range items {
		fmt.Fprintf(&b, "\t%s", item.Name)
		fmt.Fprint(&b, strings.Repeat("\t", (40-len(item.Name)+7)/8))
		fmt.Fprintf(&b, "$%5.2f\n", item.Price)
		sum += item.Price
	}
	fmt.Fprintln(&b, strings.Repeat("-", 55))
	fmt.Fprintf(&b, "\tTotal:\t\t\t\t\t$%.2f\n", sum)

	w.Write(b.Bytes())

	sendReceipt.Call(c, user.Current(c).Email, b.String())
}

var sendReceipt = delay.Func("send-receipt", func(c appengine.Context, dst, body string) {
	msg := &mail.Message{
		Sender:  appAdmin,
		To:      []string{dst},
		Subject: "Your Gopher Mart receipt",
		Body:    body,
	}
	if err := mail.Send(c, msg); err != nil {
		c.Errorf("mail.Send: %v", err)
	}
})

func adminPopulate(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	for i := range [numItems]struct{}{} { // r hates this. tee hee.
		key := itemKey(c, i)
		good := goods[rand.Intn(len(goods))]
		item := &Item{
			// TODO: vary names more
			Name:  fmt.Sprintf("%s %dg", good.name, i+1),
			Price: float64(rand.Intn(1999)+1) / 100,
		}
		if _, err := datastore.Put(c, key, item); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	fmt.Fprintf(w, "ok. %d items populated.", numItems)
}

var goods = [...]struct {
	name string
}{
	{"Gopher Bran"},
	{"Gopher Flakes"},
	{"Gopher Grease"},
	{"Gopher Litter"},
}

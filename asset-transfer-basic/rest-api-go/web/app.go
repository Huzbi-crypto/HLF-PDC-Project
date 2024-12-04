package web

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// OrgSetup contains organization's config to interact with the network.
type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
}

// Serve starts HTTP web server with parallelized query and invoke handlers.
func Serve(setups OrgSetup) {
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		// Use goroutine to handle query requests in parallel
		go setups.Query(w, r)
	})

	http.HandleFunc("/invoke", func(w http.ResponseWriter, r *http.Request) {
		// Use goroutine to handle invoke requests in parallel
		go setups.Invoke(w, r)
	})

	fmt.Println("Listening (http://localhost:3000/)...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}

// Query handles query requests.
func (setups *OrgSetup) Query(w http.ResponseWriter, r *http.Request) {
	// Create a wait group to synchronize goroutines in query.
	var wg sync.WaitGroup
	// Assuming there are parallelizable tasks inside the query.
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Your query logic here (e.g., interacting with Fabric network)
		fmt.Fprintln(w, "Query executed!")
	}()
	wg.Wait()
}

// Invoke handles invoke requests.
func (setups *OrgSetup) Invoke(w http.ResponseWriter, r *http.Request) {
	// Create a wait group to synchronize goroutines in invoke.
	var wg sync.WaitGroup
	// Assuming there are parallelizable tasks inside the invoke.
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Your invoke logic here (e.g., interacting with Fabric network)
		fmt.Fprintln(w, "Invoke executed!")
	}()
	wg.Wait()
}

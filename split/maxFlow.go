package split

import (
	"fmt"
	"ledger/models"
	"math"
	"sort"
)

func indexMap(names []string) (map[string]int, map[int]string) {
	nameToId := make(map[string]int)
	idToName := make(map[int]string)
	for i, name := range names {
		nameToId[name] = i
		idToName[i] = name
	}
	return nameToId, idToName
}

// does not work. do not use
func SimplifyDebtsMax(balances map[string]float64) {
	var debtors, creditors []string
	for name, bal := range balances {
		if math.Abs(bal) < models.EPSILON {
			continue
		}
		if bal < 0 {
			debtors = append(debtors, name)
		} else {
			creditors = append(creditors, name)
		}
	}

	all := append(debtors, creditors...)
	sort.Strings(all)
	nameToId, idToName := indexMap(all)
	n := len(all)

	graph := models.NewGraph(n)
	graph.Names = idToName
	graph.Ids = nameToId

	for _, debtor := range debtors {
		for _, creditor := range creditors {
			if debtor != creditor {
				graph.AddEdge(nameToId[debtor], nameToId[creditor], math.Min(-balances[debtor], balances[creditor]))
			}
		}
	}

	result := make(map[[2]string]float64)
	for _, debtor := range debtors {
		for _, creditor := range creditors {
			if debtor == creditor {
				continue
			}
			tempGraph := models.NewGraph(n)
			tempGraph.Names = graph.Names
			tempGraph.Ids = graph.Ids
			for i := range n {
				for _, e := range graph.Adj[i] {
					tempGraph.AddEdge(e.Origin, e.To, e.Cap)
				}
			}
			flow := tempGraph.MaxFlow(nameToId[debtor], nameToId[creditor])
			if flow > models.EPSILON {
				key := [2]string{debtor, creditor}
				result[key] += flow
			}
		}
	}

	fmt.Println("Simplified Transactions (using max-flow logic):")
	for k, v := range result {
		if v > models.EPSILON {
			fmt.Printf("%s pays %f to %s\n", k[0], v, k[1])
		}
	}
}

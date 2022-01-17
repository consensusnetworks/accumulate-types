package main

import (
	"fmt"
	"github.com/AccumulateNetwork/accumulate/types/api"
	"github.com/AccumulateNetwork/accumulate/types/api/query"
	"os"

	"github.com/32leaves/bel"
)

func main() {
	var structs []interface{}

	// which are the needed types and where are they?
	structs = append(structs, api.APIDataResponse{})
	structs = append(structs, api.APIDataResponsePagination{})
	structs = append(structs, api.APIRequestChainId{})
	structs = append(structs, api.APIRequestKeyPage{})
	structs = append(structs, api.APIRequestRaw{})
	structs = append(structs, api.APIRequestURL{})
	structs = append(structs, api.APIRequestRawTx{})
	structs = append(structs, api.APIRequestTxId{})
	structs = append(structs, api.APIRequestURLPagination{})
	structs = append(structs, api.MerkleState{})
	structs = append(structs, api.Signer{})
	structs = append(structs, query.Query{})
	structs = append(structs, query.RequestDataEntry{})
	structs = append(structs, query.RequestTxHistory{})
	structs = append(structs, query.ResponseTxHistory{})
	structs = append(structs, query.ResponseByTxId{})
	structs = append(structs, query.ResponseByChainId{})
	structs = append(structs, query.ResponseKeyPageIndex{})

	const source = "types.d.ts"

	f, err := os.Create(source)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var tsTypes []bel.TypescriptType

	fmt.Println("Generating types for:")
	for _, t := range structs {
		extract, err := bel.Extract(t)
		if err != nil {
			panic(err)
		}
		tsTypes = append(tsTypes, extract[0])
		fmt.Printf("🔎 %s \n", extract[0].Name)
		// wrap in a namespace, add to the preamble and write to file
	}

	err = bel.Render(tsTypes,
		bel.GenerateAdditionalPreamble("// Generated TypeScript types for Accumulate API \n"),
		//bel.GenerateNamespace("accumulate"),
		bel.GenerateOutputTo(f),
	)

	fmt.Printf("Result saved to: %s\n", f.Name())
	if err != nil {
		panic(err)
	}

}

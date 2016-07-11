Simple go client for the ReliefWeb API
======================================

This package provides a simple interface to query the ReliefWeb API (https://api.reliefweb.int/v1) in [go](https://golang.org).

Check the [ReliefWeb API documention](http://apidoc.rwlabs.org) for additional information on available resources, fields etc.

**Note:** the client itself doesn't do any validation of the query. It's the responsibility of the program using this client. Queries should be built accordingly to the information available in the documentation (ex: resources, field names, filter value data types).

Example usage:

```go
package main

import (
  "bytes"
  "encoding/json"
  "log"

  "github.com/reliefweb/api-go-client"
)

// Define a resource structure.
// See the ReliefWeb API resources and fields description
// at http://apidoc.rwlabs.org.
type RWAPIReport struct {
  Title    string `json:"title"`
  Headline struct {
    Title string `json:"title"`
  } `json:"headline"`
  Source []struct {
    Shortname string `json:"shortname"`
  } `json:"source"`
}

func main() {
  client := rwapi.NewClient()

  // Filter to retrieve all OCHA headlines.
  filter := rwapi.NewFilter()
  filter.AddCondition("headline", nil, "", false)
  filter.AddCondition("source", "OCHA", "", false)
  filter.SetOperator("AND")

  // Aggregation of countries for thoses headlines.
  facet := rwapi.NewFacet()
  facet.SetField("country")
  facet.SetSort("count", "asc")
  facet.SetLimit(10)

  // Aggregation of the publication years for those headlines.
  facet2 := rwapi.NewFacet()
  facet2.SetField("date")
  facet2.SetSort("value", "desc")
  facet2.SetInterval("year")

  // Main query.
  query := rwapi.NewQuery()
  // Retrieve a maximum of 5 headlines.
  query.SetLimit(5)
  // Retrieve the headline title and source short name for each headline.
  query.SetFields([]string{"headline.title", "source.shortname"}, nil)
  query.SetFilter(filter)
  query.AddSort("date", "desc")
  query.AddFacet(facet)
  query.AddFacet(facet2)

  // Query the API at the "reports" endpoint and return the
  // raw JSON output.
  raw, err := client.QueryRaw("reports", query)
  if err != nil {
    log.Fatal(err)
  }

  // "Prettify" the json response payload.
  var data bytes.Buffer
  if err := json.Indent(&data, raw, "", "\t"); err != nil {
    log.Fatal(err)
  }
  log.Println(data.String())

  // Unserialize the json response payload.
  var result *rwapi.Result
  if err := json.Unmarshal(raw, &result); err != nil {
    log.Fatal(err)
  }

  // Print the first date facet year.
  log.Println(result.Embedded.Facets["date"].Data[0].Value)

  // Print the headline titles.
  for _, item := range result.Data {
    var report *RWAPIReport
    if err := json.Unmarshal(item.Fields, &report); err != nil {
      log.Fatal(err)
    }
    log.Println(item.Id + " - " + report.Source[0].Shortname + ": " + report.Headline.Title)
  }
}
```

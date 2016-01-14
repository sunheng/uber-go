# uber-go
uber-go is an Uber API implemented using Google's Golang. The current support allows users to connect to Uber API and query the pricing and time estimate for a given location and destination.

##Usage
Server token can be retrieved from Uber at https://developer.uber.com/.
Start by creating the client and providing the server token:

```
options := uber.Options{
	ServerToken: "XXXXXXXXXXXXX",
}
client := uber.NewClient(&options)

```

#####To get a list of pricings, provide the starting and ending coordinates.
```
prices, err := client.GetPriceEstimates(&uber.PriceQuery{startLat, startLong, endLat, endLong})
```

To see all available support, see uber.go.




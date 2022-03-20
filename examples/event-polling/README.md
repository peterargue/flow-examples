# event-polling

This example shows how to periodically poll for events from an Access Node. Rather than requesting events for each block individually, you can query for your events over a range of block heights.
This allows you to use time based polling to avoid unnecessary processing.

Note: This is a very simple example. If your use case requires multiple modules tracking different events, or you want more robust error handling, check out the [flow-event-poller](https://github.com/peterargue/flow-event-poller) library.


## Usage:
```
go run main.go
```
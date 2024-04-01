---
name: "Build Scalable Real-time Applications on Fly.io + Remix!"
date: "December 04 2023"
tags: ["remix.js", "fly.io", "article"]
---

# Scalability... what is it?

You often think of vertical scaling for something to scale, adding more resources to whatever issue you encounter (CPU, RAM, etc.). While this is the simplest way to scale an application or service, it can only get you so far. Horizontal scaling involves adding more machines or "nodes" running the same application. These nodes sit behind a load balancer responsible for directing client traffic to the appropriate machine to optimize load. This will work well if the application is meant to do simple CRUD operations for individual users. What if you need collaboration in real time? This complicates things...

# The real-time problem:

Real-time applications require a pub/sub or publish and subscribe model. A client will send a request to the application to perform an operation. Once done, the server will broadcast an event to all subscribing clients to trigger a re-fetch of data. In the case of a multi-node application, you need to use a service outside of your nodes to synchronize messages across all nodes.

# The stack:

For this stack, I chose Remix for its close adherence to web standards and easy support for server-sent events. These web socket connections work one way: from server to client. Next, we must synchronize all Server Sent events across different requests to a single node. For this, Node.js has its own Event Emitter API, which we can use. Now, we can use something like Redis and its Pub/Sub commands for multi-node setups to broadcast across nodes.

This is what it would look like:
![Diagram](https://link.storjshare.io/s/jwmhimh32pura4pyr5h5luou6qla/atridad%2Farticles/scalability.png?wrap=0)

# How does it work?

Once a client connects to a page with real-time enabled, a persistent connection via EventStreams is made. The client would request to make a real-time update. Once the endpoint completes the request, it triggers a Node.js event, which the EventStreams endpoint listens to. Once received, the application sends an event down to the client via a server-sent event while also passing that request on to Redis pub/sub. Every node listens to this Redis event stream, so every node will receive the event and trigger the event using Node.js events, which then start server-sent events.

As you can see, there are many moving parts here, and it can get quite complicated. I have a repo called Atash, which acts as a template to get started. You can check it out [here](https://atash.atri.dad)!

If you found this helpful, please let me know by email at [me@atri.dad](mailto:me@atri.dad). Until next time!

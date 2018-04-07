# Composition in Distributed Systems

Distributed systems are composed of many smaller systems. The three fundamental composition patterns are:

- Load balancer with multiple backend replicas
- Server with multiple backends
- Server tree


## Load Balancer with Multiple Backend Replicas

![alt text](http://ptgmedia.pearsoncmg.com/images/chap1_9780321943187/elementLinks/01fig01_alt.jpg)

In this type of composition, for each request, the loadbalancer selects one backend and forwards
the request there. The response comes back to the load balancer server, which in turn relays it to the
original requester.

The backends are called replicas because they are all clones or replications of each other. A request
sent to any replica should produce the same response.

To know which backends are alive and ready to accept requests, Load
balancers send health check queries dozens of times each second and stop sending traffic to that
backend if the health check fails.

To pick which backend to send a query, there are various methods : 

* round-robin
* least loaded scheme
* slow start algorithm

Round robin is the traditional way of selecting each backend in a loop.

Least loaded scheme is a complex one and in this type of alogorithm, a load balancer
tracks how loaded each backend is and always selects the least loaded one.
Selecting the least loaded backend sounds reasonable but a naive implementation can be a disaster.

To avoid this, we implement an alogorithm, where we would have some kind of control over the number of requests send to the backend, this is slow-start alogorithm.


## Server with Multiple Backends

![alt text](http://ptgmedia.pearsoncmg.com/images/chap1_9780321943187/elementLinks/01fig02_alt.jpg)

## Server Tree

![alt text](http://www.informit.com/content/images/chap1_9780321943187/elementLinks/01fig03_alt.jpg)

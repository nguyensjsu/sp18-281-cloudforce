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

In this kind of composition, a server receives a request, sends
queries to many backend servers, and composes the final reply by combining those answers. 


In Figure a, different backends do their own functions and after they are done, they send back the information.
Once the replies are received, the frontend makes up the search results by combining these, which is then
sent as the reply.


Figure b illustrates the same architecture with replicated, load-balanced, backends. The same
principle applies but the system is able to scale and survive failures better.
This kind of composition has many advantages. The backends do their work in parallel, the system is loosely
coupled, one backend can fail and the page can still be constructed.


The term fan out refers to the fact that one query results in many new queries, one to each backend.
The queries “fan out” to the individual backends and the replies “fan in” as they are set up to the frontend
and combined into the final result.


## Server Tree

![alt text](http://www.informit.com/content/images/chap1_9780321943187/elementLinks/01fig03_alt.jpg)

This composition pattern is the server tree. In this scheme a number of servers work cooperatively with one as the root of the tree, parent servers below it, and leaf servers at the bottom of the tree. 
Typically this pattern is used to access a large dataset or corpus. The corpus is larger than any one machine can hold;
thus each leaf stores one fraction or shard of the whole.

To query the entire dataset, the root receives the original query and forwards it to the parents. The
parents forward the query to the leaf servers, which search their parts of the corpus. Each leaf sends its
findings to the parents, which sort and filter the results before forwarding them up to the root. The root
then takes the response from all the parents, combines the results, and replies with the full answer.

The primary benefit of this pattern is that it permits parallel searching of a large corpus.



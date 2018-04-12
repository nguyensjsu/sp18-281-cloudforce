MicroServices general Architecture:

These are process that communicate with each other in many ways. Popularly these communicate through HTTP or sharing memory between processes. 

Services in a microservice architecture should be independently deployable.
The services are easy to replace.
Services are organized around capabilities, e.g., user interface front-end, recommendation, logistics, billing, etc.
Services can be implemented using different programming languages, databases, hardware and software environment, depending on what fits best.
Services are small in size, messaging enabled, bounded by contexts, autonomously developed, independently deployable, decentralized and built and 
released with automated processes.

A microservices-based architecture:

Naturally enforces a modular structure.
Lends itself to a continuous delivery software development process.
 A change to a small part of the application only requires one or a small number of services to be rebuilt and redeployed.
Adheres to principles such as fine-grained interfaces (to independently deployable services), business-driven development 
(e.g. domain-driven design), IDEAL cloud application architectures, polyglot programming and persistence, lightweight 
container deployment, decentralized continuous delivery, and DevOps with holistic service monitoring.
Provides characteristics that are beneficial to scalability.


Benefits of MicroServices:
Enables the continuous delivery and deployment of large, complex applications.
Better testability - services are smaller and faster to test
Better deployability - services can be deployed independently
It enables you to organize the development effort around multiple, auto teams. 
It enables you to organize the development effort around multiple teams.
 Each team is owns and is responsible for one or more single service. 
 Each team can develop, deploy and scale their services independently of all of the other teams.
Each microservice is relatively small
Easier for a developer to understand
The IDE is faster making developers more productive
The application starts faster, which makes developers more productive, and speeds up deployments
Improved fault isolation. For example, if there is a memory leak in one service then only that service will be affected. 
The other services will continue to handle requests. In comparison, one misbehaving component of a monolithic architecture can bring down the entire system.
Eliminates any long-term commitment to a technology stack. 
When developing a new service you can pick a new technology stack. 
Similarly, when making major changes to an existing service you can rewrite it using a new technology stack.

Challenges and Driving forces:
One challenge with using this approach is deciding when it makes sense to use it.
 When developing the first version of an application, you often do not have the problems that this approach solves. 
 Moreover, using an elaborate, distributed architecture will slow down development. 
 This can be a major problem for startups whose biggest challenge is often how to rapidly evolve the business model and accompanying application. 
 Using Y-axis splits might make it much more difficult to iterate rapidly. 
 Later on, however, when the challenge is how to scale and you need to use functional decomposition, the tangled dependencies might
 make it difficult to decompose your monolithic application into a set of services.
 
 Other Important challenges faced are:
 1. When to use MicroService Architecture
 2. How to decompose an application into MicroServices
 3.How to maintain Data consistency
 4. How to implement queries
 
 
 
 Setting Up Mongo in AP mode:
 The CAP Theorem is: where C is consistency, A is availability, and P is partition tolerance, you can't have a system that has all three. (It gets to be called a theorem because it has been formally proved.)

Roughly speaking:

Consistency means that when two users access the system at the same time they should see the same data.
Availability means up 24/7 and responds in a reasonable time.
Partition Tolerance means if part of the system fails, it is possible for the system as a whole to continue functioning.
If you have a web app backed by a SQL database, most likely, it is CA.

It is C because it's transaction-based. So when you update the database, everything stops until you've finished. So anything reading from the database will get the same data.

It can be A, but it won't be P because SQL databases tend to run on single nodes.

If you want your application to be P, according to the CAP theorem, you have to sacrifice either A or C.

With MongoDB, in order to gain P, you sacrifice C. There are various ways to set it up, but in our application we have one master database, that all writes go to, and several secondaries (as can be seen from the diagram: M is the Master, the Rs are the secondaries – also called replicas, or slaves). Reads may come from the secondaries. So it is possibly that one or more of the secondary nodes could be disconnected from the application by some kind of network failure, but the application will not fall over because the read requests will just go to another node. Hence P.

The reason this sacrifices C is because the writes go to the master, and then take some time to filter out to all the secondaries. So C is not completely sacrificed – there is just a possibility that there may be some delay. We are not allowing a situation where the secondaries are permanently out of synch with the master – there is "eventual consistency".

So you might use this in applications where, for example, you are offering the latest news story. If User A gets the latest news 10 seconds earlier than User B, this doesn't really matter. Of course, if it was a day later, then that would be a problem. The failure case of C is just around the time of the write and you want to keep that window of consistency small.

There is also a concept of durability, which you can also be flexible with.
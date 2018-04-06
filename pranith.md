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
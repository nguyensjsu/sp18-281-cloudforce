## AKF scaling Cube






![AKF](https://ranjithabalaraman.files.wordpress.com/2014/10/scaledb.png?raw=true)







- ***X Horizontal duplication***

**Horizontal duplication** increases throughput by replicating the service. It is also known as **horizontal
scaling** or **scaling out**.A group of shared resources is called a resource pool. When adding resources to a pool, it is
necessary for each replica to be able to handle the same transactions, resulting in the same or equivalent
results.The x-axis does not scale well with increases in data or with complex transactions that require special
handling. If each transaction can be completed independently on all replicas, then the performance
improvement can be proportional to the number of replicas. There is no loss of efficiency at scale.

- ***Y Functional or Service Splits***

 **functional or service split** means scaling a system by splitting out each individual function so that it
can be allocated additional resources.
Separating the functions requires making them less tightly coupled to each other. When they are
loosely coupled, it becomes easier to scale each one independently. For example, we could apply x-axis
scaling techniques to a single subsystem. Scaling individual parts has advantages. It is less complicated
to replicate a small part rather than an entire system. It is also often less expensive to replicate one part
that needs more capacity than the entire system, much of which may be performing adequately at the
current scale.

- ***Z lookup oriented split***

A **lookup-oriented split** scales a system by splitting the data into identifiable segments, each of which
is given dedicated resources. z-axis scaling is similar to y-axis scaling except that it divides the data
instead of the processing.
A simple example of this is to divide, or segment, a database by date. If the database is an
accumulation of data, such as log data, one can start a new database server every time the current one
fills up. There may be a database for 2013 data, 2014 data, and so on. Queries that involve a single year
go to the appropriate database server. Queries that span years are sent to all the appropriate database
servers and the responses are combined. If a particular year’s database is accessed so often that it
becomes overloaded, it can be scaled using the x-axis technique of replication. Since no new data is
written to past years’ servers, most servers can be simple read-only replicas


## Scaling a SAAS Application


Can your application scale? Out of the box, many applications cannot scale well. They are not designed to handle scenarios like multiple users accessing the same pieces of data, load balanced servers, etc. Well-designed SaaS applications will need to have a strong layer between the data and the application so the data backend can scale separately from the business logic and presentation layers. They will also need to be able to provide their business logic via APIs so that mobile applications (and other non-Web clients) can be built, as well as enable integrations with partners' and customers' systems. Again, these are all things that are typically not addressed with "shrink wrapped" software, but for SaaS they are virtually mandatory.

## Why application scalability

An application scalability is the potential of an application to grow in time, being able to efficiently handle more and more requests per minute (RPM). It’s not just a simple tweak you can turn on/off, it’s a long-time process that touches almost every single item in your stack, including both hardware and software sides of the system.

In case of problems you can keep adding new CPUs or increase memory limits, but by doing so, you’re just increasing the throughput, not the app performance. It’s not the way you should stick to when you see your app is starting to have efficiency problems. Scaling the app is not an easy thing and thus you should know your app very well before starting to think about how and when to scale it.





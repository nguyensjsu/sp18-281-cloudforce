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






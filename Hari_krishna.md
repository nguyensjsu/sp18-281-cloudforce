
## Addressing X-axis in AKF Cube.

The most commonly used approach of scaling an solution is by running multiple identical copies of the application behind a load balancer also known as X-axis scaling. That’s a great way of improving the capacity and the availability of an application.

When using X-axis scaling each server runs an identical copy of the service (if disaggregated) or monolith. One benefit of the X axis is that it is typically intellectually easy to implement and it scales well from a transaction perspective.  Impediments to implementing the X axis include heavy session related information which is often difficult to distribute or requires persistence to servers – both of which can cause availability and scalability problems.  Comparative drawbacks to the X axis is that while intellectually easy to implement, data sets have to be replicated in their entirety which increases operational costs.  Further, caching tends to degrade at many levels as the size of data increases with transaction volumes.  Finally, the X axis doesn’t engender higher levels of organizational scale.

## Addressing Y-axis in AKF Cube.

Y-axis scaling (think services oriented architecture, micro services or functional decomposition of a monolith) focuses on separating services and data along noun or verb boundaries.  These splits are “dissimilar” from each other.  Examples in commerce solutions may be splitting search from browse, checkout from add-to-cart, login from account status, etc.  In implementing splits,  Y-axis scaling splits a monolithic application into a set of services. Each service implements a set of related functionalities such as order management, customer management, inventory, etc.  Further, each service should have its own, non-shared data to ensure high availability and fault isolation.  Y axis scaling shares the benefit of increasing transaction scalability with all the axes of the cube.

Further, because the Y axis allows segmentation of teams and ownership of code and data, organizational scalability is increased.  Cache hit ratios should increase as data and the services are appropriately segmented and similarly sized memory spaces can be allocated to smaller data sets accessed by comparatively fewer transactions.  Operational cost often is reduced as systems can be sized down to commodity servers or smaller IaaS instances can be employed.

## Addressing Z-axis in AKF Cube.

Whereas the Y axis addresses the splitting of dissimilar things (often along noun or verb boundaries), the Z-axis addresses segmentation of “similar” things.  Examples may include splitting customers along an unbiased modulus of customer_id, or along a somewhat biased (but beneficial for response time) geographic boundary.  Product catalogs may be split by SKU, and content may be split by content_id.  Z-axis scaling, like all of the axes, improves the solution’s transactional scalability and if fault isolated it’s availability. Because the software deployed to servers is essentially the same in each Z axis shard (but the data is distinct) there is no increase in organizational scalability.  Cache hit rates often go up with smaller data sets, and operational costs generally go down as commodity servers or smaller IaaS instances can be used.


## What are the Problems with App Scaling

Scalability issues when your project grows too large, but the application ability to scale is more about the whole system architecture, not only the framework itself. Building the project using The Rails Way is certainly not the best approach when your app is evolving rapidly, but it doesn’t mean that scaling an app is always a pain.

Even if you don’t have performance or scalability problems like Twitter or Shopify, planning and developing the application in a proper way is priceless, regardless of the nature of potential problems. You may face dozens of different issues when it comes to scaling. A few general sources of your problems may be related to:

limited physical resources like memory, CPUs etc.,

wrong memory management,

inefficient database engine,

complicated database schema, bad indexing,

poorly performed database queries,

wrong server configuration,

app server limitations,

overall spaghetti code,

inefficient caching,

lack of monitoring tools,

too many external dependencies,

improper background jobs design.



## Addressing X-axis in AKF Cube.

The most commonly used approach of scaling an solution is by running multiple identical copies of the application behind a load balancer also known as X-axis scaling. That’s a great way of improving the capacity and the availability of an application.

When using X-axis scaling each server runs an identical copy of the service (if disaggregated) or monolith. One benefit of the X axis is that it is typically intellectually easy to implement and it scales well from a transaction perspective.  Impediments to implementing the X axis include heavy session related information which is often difficult to distribute or requires persistence to servers – both of which can cause availability and scalability problems.  Comparative drawbacks to the X axis is that while intellectually easy to implement, data sets have to be replicated in their entirety which increases operational costs.  Further, caching tends to degrade at many levels as the size of data increases with transaction volumes.  Finally, the X axis doesn’t engender higher levels of organizational scale.

## Addressing Y-axis in AKF Cube.

Y-axis scaling (think services oriented architecture, micro services or functional decomposition of a monolith) focuses on separating services and data along noun or verb boundaries.  These splits are “dissimilar” from each other.  Examples in commerce solutions may be splitting search from browse, checkout from add-to-cart, login from account status, etc.  In implementing splits,  Y-axis scaling splits a monolithic application into a set of services. Each service implements a set of related functionalities such as order management, customer management, inventory, etc.  Further, each service should have its own, non-shared data to ensure high availability and fault isolation.  Y axis scaling shares the benefit of increasing transaction scalability with all the axes of the cube.

Further, because the Y axis allows segmentation of teams and ownership of code and data, organizational scalability is increased.  Cache hit ratios should increase as data and the services are appropriately segmented and similarly sized memory spaces can be allocated to smaller data sets accessed by comparatively fewer transactions.  Operational cost often is reduced as systems can be sized down to commodity servers or smaller IaaS instances can be employed.


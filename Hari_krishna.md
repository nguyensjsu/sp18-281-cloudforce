
## Addressing X-axis in AKF Cube.

The most commonly used approach of scaling an solution is by running multiple identical copies of the application behind a load balancer also known as X-axis scaling. That’s a great way of improving the capacity and the availability of an application.

When using X-axis scaling each server runs an identical copy of the service (if disaggregated) or monolith. One benefit of the X axis is that it is typically intellectually easy to implement and it scales well from a transaction perspective.  Impediments to implementing the X axis include heavy session related information which is often difficult to distribute or requires persistence to servers – both of which can cause availability and scalability problems.  Comparative drawbacks to the X axis is that while intellectually easy to implement, data sets have to be replicated in their entirety which increases operational costs.  Further, caching tends to degrade at many levels as the size of data increases with transaction volumes.  Finally, the X axis doesn’t engender higher levels of organizational scale.

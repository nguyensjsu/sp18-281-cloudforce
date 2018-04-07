Good Practices:
1. Use a whiteboard, not a keyboard
Starting a microservices initiative by coding is the wrong way to go. While creating a single service may be easy, the end state for a microservices initiative is a potentially large number of services used in aggregate to deliver a functional application.

The right way to start is to define the overall application functionality and then identify the individual services required—and the best place to do that is on a whiteboard. Start with a list of application features and then break those down into lower-level services. Iterate that process until you can’t think of any more low-level services.

In an e-commerce application, for example, one important task might be “order product.” A slew of functional items might support that task, such as checking inventory, accepting credit/debit cards, obtaining shipping addresses, calculating shipping cost, and so on. You need to identity each lower-level function so you can capture them in a service. 

Don’t expect this process to be finished in a single day. Also, you'll need to expose people from all across the organization to it so that you can include their viewpoints and requirements. After several iterations, you'll be ready to move to the next step.

2. Follow a Conway’s Law anti-pattern
Conway’s Law states that system designs inevitably reflect the structure of the organization that promulgates them. Regarding microservices, many people recommend creating microservice partitions that mirror your organizational structure. For example, if you have a shipping group, you should have a shipping microservice.

That’s not necessarily bad advice, but you shouldn’t slavishly follow it, because an organization structured to support a monolithic application architecture may be entirely wrong for a microservices move. Many application organizations have a database group that handles all database access services. But if you mirror that organization into your microservices application architecture, you’ll have one service providing data functions for every other service, creating a single point of failure. Every microservice must have its own storage arrangement so that it can continue operating even when another service fails. A single data service is the wrong way to go.

Rather than taking today’s organizational structure as the jumping-off point for architectural design, evaluate your structure in light of the functions you’ll need to deliver as part of your new architecture.

3. Incorporate operations as well as functionality
The previous two steps focus on designing an application to provide the right kind of functionality and ensuring that your organization is structured to align with it. The next step is to ensure that you have operational support for the new application architecture.

Management and monitoring will probably need to evolve to support a new, more distributed, and more dynamic application topology. It’s important to be able to track transient resources and keep a history of application metrics that you can use to perform time-series comparisons.

Part of your design should incorporate unique identifiers you pass as part of service calls and use when logging information. You can use these identifiers to correlate events or anomalies that cascade across multiple services. In this way, you can identify the root cause of a problem. For example, if you have a situation where one problem has caused a set of errors, you won't view them as just individual events.

You can facilitate this part of the design process immeasurably by discovering the best practices that other companies that have moved further into their microservices journeys have used. Many presentations at DevOps and Webscale events focus on how companies support the operations aspects of their microservices environments. Learn how they’ve done it and apply those lessons so you get this aspect of microservices design right.

4. Integrate performance and resilience
One of the main drivers behind the move to microservices is the desire to create applications that are better suited to customer-facing digital business offerings. While customers have come to accept, and even expect, that internally focused applications will run slowly or be unavailable for lengthy periods of time, response time and availability expectations for externally focused applications are much higher. This means that each microservice must deliver high performance and continue to operate even in the face of an underlying resource failure.

For example, each service should leverage a caching layer to avoid, to the greatest extent possible, calls to storage that require network and disk access. Also, a service must not hang if the caching or storage layers are unavailable. Every service should have two or three alternative mechanisms it can use to continue operating when underlying resources fail, and it should quickly cycle through them if one fails to respond within an acceptable time frame.

It's important to address performance and resilience as part of the design process to ensure that those potential issues are considered early. A common problem arises from the fact that microservice designs focus primarily on functionality. When the system is placed into production and receives a load for the first time, it may exhibit slow response times or just hang due to resource saturation. Thinking about how to prepare your service for 10X load variations and operation continuance in the face of underlying resource failure will help ensure that your application stands up to real-world challenges.

5. Plan for the future
An inevitable aspect of applications is that they must evolve. Developers add new functionality, refactor code, and replace core components with newer alternatives all the time. This is even more true with microservices applications. In fact, it might be more correct to say that microservices applications are constantly morphing, given how commonly this architecture is associated with continuous deployment initiatives.

When an application experiences code updates several times a day, it’s probably better to recognize that change is constant, rather than an occasional interruption to an ongoing, steady state. As such, it’s important to integrate the flexibility necessary for change right from the start of the design process. One aspect of this is to anticipate version changes to service APIs so that individual services can continue to interact with other services, even as they change. You can also use version control to allow services to include old and new service interfaces. In this way, you can expose new functionality to calling services that are aware of it, while ensuring that calls to services that have not yet been upgraded to work with the new functionality continue to operate.

Enabling a data storage evolution is more challenging. Upgrading database schemas to support new functionality traditionally has been the hardest part of evolving an application, and microservices do not make this easier. However, the new breed of NoSQL databases are far more flexible in terms of adding new fields without disrupting the existing arrangement. If you expect your data storage requirements to evolve (and who doesn't?), you should incorporate evolvable data storage as part of your microservices design effort.

Keys to success
Success with microservices doesn’t just happen. While the benefits are significant, microservices require careful planning if you want to achieve good outcomes. I've reviewed five key areas that you should address as you plan your move to microservices, but experienced professionals may have additional design tips.

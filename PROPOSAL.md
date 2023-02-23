# Technical Proposal Document

## What

An HTTP server is an essential part of the world wide web. It's the fabric between your user interfaces, your business logic and your data. It's platform agnostic, can be deployed in a wide range of systems, and serves web content for all your requiring clients. Designing and developing the web service of your business is a very important process, that should be deeply thought of and discussed. It will be the layer through which your business expresses itself in the world, and might be the breaking point between an awesome or awful experience.

There are usually two paths to take when building your web service; use an existing HTTP server framework/library, or the standard library of your programming language of choice (Go in our case).

It provides tools and libraries that solve common problems when dealing with these solutions. A couple of important features when looking for a framework of this kind are:

- Routing: the mechanism behind matching routes to their specific handlers. This can include string routes, parameterized routes, regex pattern routes, and so on.
- Middlewares: frameworks usually provide the ability to add middlewares, which act as a pipeline through which the incoming requests pass. This is often used to add authentication logic, logging capabilities, request/response parsing, etc.
- Request Handling: the ability to parse request data bodies, validate headers and data, and generating proper HTTP responses.
- Templating: a more rare feature, but some frameworks provide the ability to generate dynamic content in the form of HTML templates, which can be used to generate pages, emails, etc.
- Integration with existing services: some frameworks go above and beyond to provide integration with databases, auth providers, and other services.

It is important to consider the alternative of using the standard library of the language to implmenet your server needs. In Go's case, this would include using `net/http` from stdlib to implement our HTTP server. This is a good solution when your scope is smaller and would not require many complex features. Also, there is always the risk of using third-party libraries/frameworks and they get discontinued with their support and ongoing development. However, going down the path of using the standard library would require a lot of development resources to get done properly, and would still be a risky endeavour, considering that well estabilished frameworks have been extensively tested and scrutinized already. Some of the features we would need to implement if we went this path:

- Manual error handling
- Templating
- Common middlewares like logging and crash recovery
- Big test coverage needed for reliability
- Increased maintenance costs and overhead



### Options

- Gin
- Echo
- Chi
- Fiber


- Fiber
  


### Options Evaluated


## Why

### Purpose


### Benefits


## Who

### Intended Audience


### Interaction and Observation


### Integrating Applications and External Parties


## How

### Success Criteria



### Analysis Process


### Analysis Results


### Recommendations


## Architectural Impact


## Associated Costs


## Security Considerations


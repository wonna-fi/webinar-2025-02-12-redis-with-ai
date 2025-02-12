## Research

### Prompt 1:

I'm implementing a lightweight Redis server clone. What resources should I research to get an understanding of how a Redis server should work?

### Prompt 2

Thank you! Please summarize the most important parts of the RESP protocol with examples in less than 300 words.


## Programming language choice

I'm going to write a Redis server clone. Can you please help me choose the programming language that I'm going to implement it in? Here's what I would like you to do:

1) Assess what the technical requirements of a Redis server are
2) Based on the technical requirements, pick the most suitable modern programming languages for the job
3) Further evaluate how well those programming languages are suited for AI-assisted development, ie. how much high quality teaching material is available and how complex the programming language is for an LLM to work with. Tune your picks based on this evaluation.

Please provide as output two different options for me to weigh, listing the pros and cons of both of them.


## Implementation plan

We are going to implement a lightweight Redis server in Go. It's not going to support all of the commands, but it's going to support concurrency. We are going to use the official Redis CLI to interact with the server, so it must comply with the RESP specification: https://redis.io/docs/latest/develop/reference/protocol-spec/

Please make a high level plan for in which order the various parts of the server is best to be implemented, so that I can get a good foundation to work on. Provide the plan as a series of implementation steps, each of them having a small and clear goal. We want have something working end-to-end as soon as possible, so please split the work in narrow complete vertical slices.

The goal is to create a simple Redis server that supports the following use cases:

- connecting with the Redis CLI client
- concurrency
- all data types, as defined in the RESP specification
- the following commands: PING, ECHO, GET, SET, DEL
- other necessary commands for redis-cli and redis-benchmark tools to be able to work with server

Take the following steps while working out the plan:

1. Familiarize yourself with the RESP protocol
2. Familiarize yourself with how the Redis CLI connects with the server
3. Research what needs to be implemented in order to support the required use cases
4. Plan a simple architecture for the server that is easy to extend to support more commands and features. Don't over-engineer it.
5. Make a plan by planning narrow vertical slices from the work that needs to be done to implement the server's features and arhitecture.
6. Validate the plan by double checking that it complies with the RESP spec and other resources
7. Output the plan as a series of small steps that contain a vertical slice to implement, each building on each other.


## Setup project

Hello, my friendly paircoder! We are going to build a Redis server clone in Golang. Please help me initialize a barebones Go project for that purpose inside the current folder (the current folder is the project folder, so please initialize it right here without an additional subfolder). make the application itself a simple "Hello world!" application, so that we can make sure that our environment works.
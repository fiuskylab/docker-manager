## Docker API
Developing a RESTful API to manage Docker Containers

## Thought Flow.

### First Things First!
This section is the most important one, trying to understand what's happening

#### Is there a Docker SDK?
Yes! Of course there is, how could I even ask that question?

#### Final Goal?
Develop a REST API to manage Docker Containers, with the following actions:
- Create
- Read
- Update (???)
- Delete

and learn something new (or improve it)!

#### Information Available
- App
  - Status (health check BABYYYYYYYY!!!)
- Container (keeping it simple)
  - Status
  - Name
  - Ports
  - Image
  - ID
  - CreatedAt

#### Am I comfortable with everything above?
Right now(didn't wrote a line of code yet)? Yes, sound good

### Technical - Part One! (keep it simple)
Making it simple and working

#### Development Workflow
1. Write Code
2. Write Tests
3. Run the Tests
  - Failed? Fix it
  - Success? Go to the first line, and repeat

#### Packages:
- Logging: [uber-go/zap](https://github.com/uber-go/zap)
  - Low allocation
  - Easy to use/implement
  - Great docs
  - Well maintained
- Web Framework: [gofiber/fiber](https://github.com/gofiber/fiber)
  - User friendly (express-like)
  - Great docs
  - Well maintained
  - Because I'm not implementing Swagger, it should be fine.
- Testing Toolkit: [stretchr/testify](https://github.com/stretchr/testify)
  - Easy to use/implement
  - Support for Suite Testing
  - Great docs
  - Well maintained
- Docker Client: [docker/client](https://pkg.go.dev/github.com/docker/docker/client#section-readme)
  - Official Client to communicate with the Docker Engine API

#### Workflow Sample (Create)
1. Receive API request
2. Validate the request
3. "Mount" the Data model
4. Create the Container

#### HTTP EndPoints
- `/health`
  - Check connections
- `/container` - Route
  - `/create` - POST
    - Create a Container
  - ` ` - GET
    - Retrieve all Containers
  - `?image=` - GET
    - Retrieve all Containers by image
  - `/:name` - GET
    - Retrieve Container by name
  - `/:name` - PUT/PATCH
    - Update Container with given name
  - `/:name` - DELETE
    - Update Container with given name


### Technical - Part Two! (let's improve it)
This section is not as important as the first one, but could bring greater quality to the app.

#### Improvement Workflow
1. Write code
2. Run the tests
  - Failed? Fix it
  - Success? Go to the first line, and repeat

#### What could be improved?
- Tests:
  - Check the coverage
  - If it's only Unit Tests, what about Feature Tests?
- Container CRUD:
  - More options to customize the Container
- Controller
  - Concurrency?
    - Race Condition
    - Make it safe/safer

### Technical - Part Three! (let's have more fun)
This part is just for funsies, to see where I can reach while having fun.

### What else could be added?
- Container CRUD:
  - CPU/MEM Limit
- Observer
  - Keep checking all containers (maybe Docker API has somewhere that we can Sub)
- Internal Storage
  - Can be a JSON file, no need to be _fAnCy_
  - Only check the containers created by my Controller
- gRPC Server
  - What if we had a notification system
    - e.g: `container XYZ broke`
    - e.g `container XYZ memory/CPU usage surpassed 30%`

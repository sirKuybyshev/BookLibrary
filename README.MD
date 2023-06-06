## Librarian 
Book management service.

### Launch

#### Unit-tests
`make unit-tests`

#### DemoRun
`make demo_run` launch server on `localhost:55001` with example database.

#### Run
`make run config=<path_to_your_config>` launch server with provided config.

### Configuration
Example config is placed in `cmd/config`. All fields are required.

### GRPC
Proto-spec is placed in `proto/protos`. There are 2 methods: `GetAuthor` return author of book, provided in request and `GetBooks` return all existing in db books of provided author. Now service do not use TLS

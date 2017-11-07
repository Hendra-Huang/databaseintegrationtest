This is a sample of parallel database integration test

The main idea is separating data for every function test. In this example, the database is PostgreSQL and the data is separated by schema.

testingutil package contains helper function for printing error.

database package contains function for creating a test-database connection.

testdata folder contains dummy data for testing.

To run the test:
1. Start the docker by executing `docker-compose up -d`
2. Run the test by executing `go test -v ./... -tags integration`

For more information, you can read the blog post in [medium](https://medium.com/@hendra_zhong/parallel-database-integration-test-on-go-application-8706b150ee2e)

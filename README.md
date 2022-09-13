Rest API Client Interface

Here our RestAPIClient package's WebCallV2 function calls directly Do function call on the http.Client instance, any tests we write that trigger code pathways using this client will actually make a real web request. This means we will be spamming the real system with our fake test data.

Further, relying on real Real API interactions makes it difficult for us to write legible tests in which a given set of input results in an expected outcome. We end up with tests that are less declarative, that force other developers reading our code to infer an outcome based on the input to a particular HTTP request.

So, how can we write clear and declarative tests that avoid sending real web requests?

We'll build a mock HTTP client and configure our tests to use that mock client. Then, each test can clearly declare the mocked response for a given HTTP request. But, In order to build our mock client and teach our code when to use the real client and when to use the mock, we'll need to build an interface. Refer WebCallV1 function which uses interface to make Do Function call on http.Client Instance.

# References
1) https://github.com/stretchr/testify
2) https://pkg.golang.ir/github.com/stretchr/testify/mock

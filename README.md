# GoLang Unit Tests
A sample project for Software Engineers looking to perform unit tests in Go
<br>
This example shows a manual test, but can easily be automated via CICD pipelines or runners by DevOps engineers.

## How to execute the test:
1. Clone the repo to a folder on your PC
2. Open Command Prompt
3. Navigate to the folder in Step 1, from within Command Prompt
4. Run: <code>go test main.go main_test.go </code>
5. You should get an output that says "ok", or an error message showing where the test failed

## Notes:
1. To create a Test file, create a new file called *_test.go, where * is the name of the file you're creating tests for.
2. Each test method has to start with "Test", and the following word should be capitalized.
3. Each test method should be within the same package as the method/function it's testing.
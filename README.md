Para rodar:

Sem docker:
  - go build -o ./cli .
  - ./cli --url=http://google.com --requests=50 --concurrency=10

Via docker:
  - docker build -t go-stress-test-cli .
  - docker run go-stress-test-cli --url=http://google.com --requests=50 --concurrency=10

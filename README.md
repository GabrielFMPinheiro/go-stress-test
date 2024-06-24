Para rodar:

Sem docker:
  - go build -o ./cli .
  - ./cli --url=http://google.com --requests=50 --concurrency=10

Via docker:
  - docker build -t stress-test .
  - docker run stress-test --url=http://google.com --requests=50 --concurrency=10

### GO-REACT

A simple React app with a Go back end I put together after being interrupted by
my kids during a coding test.

### HOW TO RUN

- Clone repo.
- Go into `client` folder.
- Run `yarn install`.
- Run `yarn run:server && yarn run:client`.
- Navigate to `http://localhost:3000`.
- Profit.
- Let me know if you want me to add a GraphQL server.

### SOME NOTES

- The React app uses hooks and a functional component, no classes or lifecycle hooks.
- The Go server sets an open CORS header so the client can talk to the server.

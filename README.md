# wabisabi

`wabisabi` is a stand-alone micro-service for managing JWT-based sessions, written in GoLang.

It **is not** an authentication service. Instead, `wabisabi` provides a ready to go light-weight server for creating and managing JWT-based sessions.

## Why?

#### *JWT-based sessions* sounds like an oxymoron.

The promise of JWTs was to be stateless while sessions, are by definition a form of state and introduce scaling concerns... The bottom line is there is [no practical means to revoke a JWT](https://redis.io/blog/json-web-tokens-jwt-are-dangerous-for-user-sessions/) once it's issued without introducing state.

But there is a potential middle ground here.

#### The Middle-Ground

`wabisabi` strikes a balance by utilizing JWTs to increase the scalability of a session-based system while keeping the security posture of one by allowing revocation.

It increases scalability because JWTs are signed, self-contained tokens with additional claims we can assert against without ever hitting I/O.

This means we can filter at the edge and only forward valid JWTs to `wabisabi`, removing I/O and overhead and even attacks that could exhaust resources just by sending invalid sessions that would otherwise result in a DB hit.


## Getting Started

You can test this out by cloning the repo and running:

```shell
go run cmd/http/main.go
```

Utilize the flag `-u` to define the format for your userId.


**WIP**: Pre-built binary releases are also available for use in building your own images/deployment strategies.


### Advanced Usage

Usage with other services looks like this:


Essentially, whenever a JWT is succesfully validated and parsed, and only when, you can check for revocation by pinging the `wabisabi` service.

See [exmaple middleware](#) to help you effectively accomplish this.

#### Data Store

`wabisabi` itself is backed by an in-memory SQLite database. Meant to be run as a single instance, this is more than capable of handling handling tens of thousands of RPS and should suffice for most use cases.

We will investigate other adapters in the future for higher volume usecases.

#### Data Access

No matter the datastore, an `allow-list` design is used. This ensures security by design by denying at default.

Datastores are also treated as *ephermeal*. If data is lost, then that token should just be denied rather than potentially introducing any race conditions.

Each token is a row, and each row contains a `token_id` and `user_id`. A token is made invalid by simply removing it from the table.

#### Code

We have a mission to keep code simple and avoid using depenencies; keep things as close to bare GoLang and trusted dependencies as possible.

#### Platform Level Security

Platform level security for `wabisabi` is up to you; you need to make sure not anyone can arbitrarily create new tokens and lock down endpoints.

(TBD: May change this design to be a middleware that also hosts a standlone validation server???)


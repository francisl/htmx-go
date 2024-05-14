# Introduction to Htmx with Go

## Branch

| Branch | Description | Command |
| --- | --- | --- |
| intro-1 | Barebone application to get started | `git clone git@github.com:francisl/htmx-go-intro.git -b intro-1` |
| intro-2 | hx-swap demo | `git clone git@github.com:francisl/htmx-go-intro.git -b intro-2` |
| intro-3 | Demo Application | `git clone git@github.com:francisl/htmx-go-intro.git -b intro-3`|


## Dependencies

1. Install Go
https://go.dev/dl/

2. Option 1 - Install a new project from scratch

```bash
git clone git@github.com:francisl/htmx-go-intro.git -b intro-3
cd htmx-go-intro 
```

2. Option 2 - Get Started with the intro project
```sh
git clone git@github.com:francisl/htmx-go-intro.git -b intro-3
cd htmx-go-intro 
```

3. Install Dependencies

```sh
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest
go get github.com/a-h/templ
go get github.com/labstack/echo/v4
```

4. Start the development server

The go compiler and runner
> air

Template compiler
> templ generate -watch 

# Repo Issues

This is a project that will display a list of issues that pertain the repository given. This was a project used for 
experimenting with the technologies below:

- Go with the [Gin Web Framework](https://github.com/gin-gonic/gin) for the backend
- [Go Templating](https://pkg.go.dev/text/template) and [HTMX](https://github.com/bigskysoftware/htmx) for the frontend
- [Tailwind](https://v2.tailwindcss.com/) for styling
- [Github's GraphQL API](https://docs.github.com/en/graphql) for the data

Some notable tools used for this project were:

- [Genqlient](https://github.com/Khan/genqlient/tree/main) for the GraphQL query to Golang code generation
- [Microsoft Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers) for a consistant development 
environment
- Docker for containerization

## Running the Devcontainer

To start the Devcontainer copy the `.devcontainer/.env.example`, name the copy `.env` and leave it in the
`.devcontainer` directory. Provide a value for `ENV`, `PORT` and `GH_PAT`.

- `ENV` tells the application what environment it is running in, `PROD` would run Gin in Release Mode
- `PORT` tells the application what port it should use to receive requests
- `GH_PAT` gives the application permission to perform actions on your behalf via the Github API, if you don't have a
personal access token from Github yet, you can create one by following
[this link](https://github.com/settings/tokens?type=beta) and then create a `Fine-grained token` no specific permissions
are needed.

After the environment variables, open the `Makefile` and set the `MOD_NAME` variables, this will be the name of the 
the application.

If you have the [devcontainer cli](https://github.com/devcontainers/cli) installed you can run the make commands below:

```bash
make dc_up
```

Once it has finished, you can open a VSCode instance in that container with the command:

```bash
make dc_open
```

If you don't have the Dev Container CLI, you can use the Dev Containers VSCode extension instead. In the VSCode command
palette, run the command `Dev Containers: Reopen in Container`, this would create a container with your code inside and
open VSCode with a remote connection already with Go installed.

Once inside the Dev Container, open a terminal and change your directory to the `workspaces/src`, from there you can
install the Project dependencies with the command:

```bash
make init
```

To run the application, run the command:

```bash
make run
```

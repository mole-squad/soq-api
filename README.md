# SoQ


## Setup

Install Taskfile

```
brew install go-task
```

Install Heroku CLI

```
brew tap heroku/brew && brew install heroku
```

### PG Admin

http://localhost:5444/ (test@test.com / pass)



## Deployment

### Set up heroku origin locally

Run the following command to add a heroku git remote. This should only be used for testing purposes

```
heroku git:remote -a <app_name>
```

### Get a database connection string

```
heroku pg:credentials:url DATABASE
```

### Set config vars

Run the following commands locally to configure necessary secrets

```
heroku config:set PUSHOVER_TOKEN="" -a <app_name>
heroku config:set JWT_SIGNING_SECRET="" -a <app_name>
```

### Run commands on remote dyno

```
heroku run -a <app_name> <cmd>
```


# Helpful Docs

### General
 - [Taskfile](https://taskfile.dev/)


### Server
 - [gORM](https://gorm.io/docs/)
 - [chi](https://go-chi.io/#/README)
 - [chi-render](https://github.com/go-chi/render)
 - [fx](https://uber-go.github.io/fx/)
 - [pushover](https://pushover.net/)
 - [pushover client](https://github.com/gregdel/pushover)
 - [cobra cmd](https://github.com/spf13/cobra)


### Heroku
 - [deploying with docker](https://devcenter.heroku.com/categories/deploying-with-docker)

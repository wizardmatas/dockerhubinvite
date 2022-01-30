# Docker Hub Invite
Tool to send Docker Hub Invites

The Docker Hub Invite tool is a tool automating member invite to 
[Docker Hub](https://hub.docker.com).
It extents Docker Hub UI with ability to invite tens or hundreds users to Docker Hub Orgasation.

## Get started

### Prerequisites

Login to the [Docker Hub](https://hub.docker.com) and note [organization]((https://hub.docker.com/orgs)) and team name, where you want to invite members.

Get token with:
```console
# curl -X POST 'https://hub.docker.com/v2/users/login' -H 'Content-Type: application/json' --data-raw '{"username": "myusername","password":"hunter2"}'`
```
Fill in env variables `.ENV`:
```
export DOCKERHUBTOKEN="Bearer adfdfglvjjdvj"
export DOCKERHUBORG=yourorg
export DOCKERHUBTEAM=teaminorg
export DOCKERHUBINVITEFILE=emails.txt
export DOCKERHUBTESTMODE=false
```
> NOTE: DOCKERHUBTESTMODE=true will set read only mode to get members from org and team.
Source it:
```
# source .ENV
```



### Install

- Download the latest release for your platform from
  [here](https://github.com/wizardmatas/dockerhubinvite/releases)
- Extract the package and place the `dockerhubinvite` binary somewhere in your `PATH`

OR

- Install from sources: `GO111MODULE=on go get github.com/wizardmatas/dockerhubinvite`

### Run

```
./dockerhubinvite
Token Check Started...
Token Check 200 OK
Start member invite...
Sending email to example1@noreply.moon 
------
Sending email to example2@noreply.moon 
------
```

### Feedback

Please leave your feedback in the
[issue tracker](https://github.com/wizardmatas/dockerhubinvite/issues)!
I'd love to know how you're using this tool and what features you'd like to see
us add.

### Code

At this stage of the project, we're mostly looking for feedback. We will accept
pull requests but these should be limited to minor improvements and fixes.
Anything larger should first be discussed as an issue.
If you spot a bug or see a typo, please feel free to fix it by putting up a
[pull request](https://github.com/wizardmatas/dockerhubinvite/pulls)!
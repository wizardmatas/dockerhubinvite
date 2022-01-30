![example workflow](https://github.com/wizardmatas/dockerhubinvite/actions/workflows/go.yml/badge.svg)

> Docker Desktop licencing changed and some [organisations](https://www.docker.com/pricing/faq) need to buy licence to run Docker Desktop, to do it you must invite members to Docker Hub where licencing is managed for Teams Plan. Tool was created to automate invite management for Teams plan.

# Docker Hub Invite
Tool to send Docker Hub Invites

The Docker Hub Invite tool is a tool automating member invite to 
[Docker Hub](https://hub.docker.com).
It extents Docker Hub UI (where you can invite only singe member) with ability to invite tens or hundreds users to Docker Hub Orgasation.

## Get started

### Prerequisites

Login to the [Docker Hub](https://hub.docker.com) and note [organization]((https://hub.docker.com/orgs)) and team name, where you want to invite members.

Add needed count of seats into your org account.

Get token with:
```console
# curl -X POST 'https://hub.docker.com/v2/users/login' -H 'Content-Type: application/json' --data-raw '{"username": "myusername","password":"hunter2"}'
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


### Run

```
git clone https://github.com/wizardmatas/dockerhubinvite.git
cd dockerhubinvite/
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
If you spot a bug or see a typo, please feel free to fix it by putting up a
[pull request](https://github.com/wizardmatas/dockerhubinvite/pulls)!

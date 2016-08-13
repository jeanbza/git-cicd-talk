# git-cicd-talk

[![Build Status](https://travis-ci.org/jadekler/git-cicd-talk.svg?branch=master)](https://travis-ci.org/jadekler/git-cicd-talk)

A talk on CI/CD.

## Running

1. Install [golang](https://golang.org/dl/)
1. `go run main.go`

## Testing

1. `go run main.go`
1. `go test`

## CI Servers

1. [TravisCI](https://travis-ci.org/jadekler/git-cicd-talk)
    - Simply include a `manifest.yml` and enable the project in [travisCI](https://travis-ci.org/)
    - Each commit will trigger the build
1. [ConcourseCI](http://concourse.ci/)
    - Example pipeline: https://ci.concourse.ci/teams/main/pipelines/main?groups=develop
    - See instructions in [concourse](concourse) folder
   
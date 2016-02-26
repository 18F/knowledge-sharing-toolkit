# Knowledge sharing toolkit

[![Build Status](https://travis-ci.org/18F/knowledge-sharing-toolkit.svg?branch=master)](https://travis-ci.org/18F/knowledge-sharing-toolkit)

**NOTE: THIS IS _ALMOST_, BUT NOT YET USABLE!!!** It will be completely ready
by 2016-03-04 (or sooner).

The knowledge sharing toolkit contains the [Hub](https://github.com/18F/hub/),
[18F Pages](https://github.com/18F/pages-server), and [Team
API](https://github.com/18F/team-api-server/). These are lightweight services
that enable a team to collect and radiate institutional knowledge and
information. This project contains [Docker](https://www.docker.com/)
components for these services, to enable rapid deployment of the entire suite.

## Installation

These steps are necessary on your development machine. They are included in
the [Deployment](#deployment) section below as well.

1. Install [Ruby](https://ruby-lang.org/) on your system. The [`./go`
   script](./go) command line interface requires version 2.3.0 or greater. You
   may wish to first install a version manager such as
   [rbenv](https://github.com/rbenv/rbenv) to manage and install different
   Ruby versions.

1. Install [Docker](https://www.docker.com/) on your system. The commands
   encapsulated in the [`./go` script](./go) are based on version 1.10.0.

1. Run `docker-machine` start to start the Docker host, followed by
   `eval $(docker-machine env)` to configure your shell environment.

1. After cloning this repository, install all the images by running `./go
   build_images` within your copy of the repository.

## Deployment

1. Install your public SSH key on the remote host machine.

1. Set the `REMOTE_HOST` and `REMOTE_ROOT` variables in the `./go` script as
   necessary.

1. Run `./go init_remote` if running on a brand new server. Otherwise run
   `./go sync_remote` to bring the server up-to-date with any changes.

1. Get the bundle of files containing secret data (`SECRETS_BUNDLE_FILE` in
   the `./go` script) and run `./go push_secrets` to install them on the
   remote host.

   These files are all masked out of the repository by the `.gitignore` file.

1. Run `./go ssh_remote` to log into the remote host. The working directory
   will be the root of the repository on the remote host.

1. Follow all of the steps from the [Installation](#installation) section
   above.

1. Run `./go start` to bring up all the system components, and `./go stop` to
   stop them all.

## Running locally

Add the following to the `/etc/hosts` file of your development machine,
commenting out any services you're not currently attempting to emulate
locally:

```
# Testing locally with 18F/knowledge-sharing-toolkit
# Run `docker-machine env` to get the current IP.
192.168.99.100 pages.18f.gov
192.168.99.100 pages-staging.18f.gov
192.168.99.100 pages-internal.18f.gov
192.168.99.100 pages-releases.18f.gov
192.168.99.100 team-api.18f.gov
192.168.99.100 hub.18f.gov
192.168.99.100 handbook.18f.gov
```

Then bring the entire system up using `./go start`. You should be able to
access any of the hosts from your `/etc/hosts` file that you've configured,
and have the content served by the Dockerized system running locally. You can
take the system down with `./go stop`.

### Rebuilding specific images

To rebuild one or more specific [images](#images):

```sh
$ ./go build_images <image_name_0> ... <image_name_n>
```

For example, this will attempt to rebuild [oauth2_proxy](#oauth2_proxy),
[hmacproxy](#hmacproxy), and [team-api](#team-api):

```sh
$ ./go build_images oauth2_proxy hmacproxy team-api
```

## Images

### dev-base

An image that contains all of the tools needed for the images.

### dev-standard

An image that pins the versions of [Go](https://golang.org/),
[Ruby](https://ruby-lang.org/), [Python](https://www.python.org/), and
[Node.js](https://nodejs.org/). Also, the basis for rest of the images in this
repository.

### oauth2_proxy

[oauth2_proxy](https://github.com/bitly/oauth2_proxy) enables nginx to
authenticate requests using an OAuth2 provider; in our case,
[MyUSA](https://staging.my.usa.gov/).

### hmacproxy

[hmacproxy](https://github.com/18F/hmacproxy) enables nginx to authenticate
requests using
[HMAC signatures](https://en.wikipedia.org/wiki/Hash-based_message_authentication_code).

### authdelegate

[authdelegate](https://github.com/18F/authdelegate) nginx to delegate
authentication of [Team API](https://team-api.18f.gov/) requests to _both_
`oauth2_proxy` and `hmacproxy`, allowing both browser-based (OAuth2) and
machine-based (HMAC) access to the same endpoints.

### pages

[18f-pages-server](https://github.com/18F/pages-server) is the server behind
[18F Pages](https://pages.18f.gov/), the [GitHub
Pages](https://pages.github.com/)-like service for publishing
[Jekyll](https://jekyllrb.com/)-based sites.

### lunr-server

[lunr-server](https://github.com/18F/lunr-server) is an early, experimental
[lunr.js](http://lunrjs.com/)-based search backend that performs a search
across statically-generated corpora from the Hub and 18F Pages. The corpora
are generated by the [`jekyll_pages_api_search` Jekyll
plugin](https://github.com/18F/jekyll_pages_api_search/) included in the Hub
and 18F Pages sites.

### team-api

The [team-api-server](https://github.com/18F/team-api-server) publishes
organizational metadata in the form of a complete graph between people,
projects, locations, skills, and interests.

### nginx

A custom [nginx web server](http://nginx.org/) build that builds with
[OpenSSL](https://www.openssl.org/) v1.0.2, enabling
[HTTP/2](http://nginx.org/en/docs/http/ngx_http_v2_module.html).

## Contributing

If you'd like to contribute to this repository, please follow our
[CONTRIBUTING guidelines](./CONTRIBUTING.md).

## Public domain

This project is in the worldwide [public domain](LICENSE.md). As stated in
[CONTRIBUTING](CONTRIBUTING.md):

> This project is in the public domain within the United States, and copyright
> and related rights in the work worldwide are waived through the
> [CC0 1.0 Universal public domain dedication](https://creativecommons.org/publicdomain/zero/1.0/).
>
> All contributions to this project will be released under the CC0 dedication.
> By submitting a pull request, you are agreeing to comply with this waiver of
> copyright interest.

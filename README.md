# Knowledge sharing toolkit

**NOTE: THIS IS NOT YET USABLE!!!** It should become completely deployable by
2016-02-26.

The knowledge sharing toolkit contains the [Hub](https://github.com/18F/hub/),
[18F Pages](https://github.com/18F/pages-server), and [Team
API](https://github.com/18F/team-api-server/). These are lightweight services
that enable a team to collect and radiate institutional knowledge and
information. This project contains Docker components for these services, to
enable rapid deployment of the entire suite.

## Dependencies
- Ruby >= 2.3.0
- Docker

## Installation

To install all the images at once: `./go build_images`

To install specific a specific image or specific [images](#images)

```

$ ./go build_images <image_name_0> ... <image_name_n>

# Example building oauth2_proxy, hmacproxy, and team-api

$ ./go build_images oauth2_proxy hmacproxy team-api

```

## Images

- dev-base: An image that contains all of the tools needed for the images.
- dev-standard: An image that pins the versions of tools to the latest. Also, the basis for rest of the images in this list.
- nginx-18f
- [oauth2_proxy](https://github.com/bitly/oauth2_proxy)
- [hmacproxy](https://github.com/18F/hmacproxy)
- [authdelegate](https://github.com/18F/authdelegate)
- [18f-pages](https://github.com/18F/pages)
- [lunr-server](https://github.com/18F/lunr-server)
- [team-api](https://github.com/18F/team-api-server)


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
